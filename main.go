package main

import (
	"embed"
	"fmt"
	"os"
	"time"

	"github.com/adrg/xdg"
	"golang.org/x/exp/slog"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

const Version = "v8"

func main() {
	// Create an instance of the app structure
	app := NewApp()

	now := time.Now()

	logPath, err := xdg.CacheFile(
		fmt.Sprintf("mazesoba-continent/%d%02d%02d-%02d%02d%02d.log",
			now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(),
		),
	)
	if err != nil {
		slog.Error("Can't open log file path", "logPath", logPath) // will it happen?
	}
	logFile, err := os.Create(logPath)
	if err != nil {
		slog.Error("Can't create log file", "logPath", logPath) // will it happen? #2
	}

	defer logFile.Close()
	app.logger = slog.New(slog.NewJSONHandler(logFile, nil))

	config, err := loadOrCreateConfig(configPath)
	if err != nil {
		app.logger.Error("Error on loadOrCreateConfig", "error", err)
		panic(err)
	}

	app.config = config

	// menu ================
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("ファイル")
	FileMenu.AddText("Quit", keys.Key("Escape"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})
	FileMenu.AddText("設定の場所を開く", keys.CmdOrCtrl(","), func(_ *menu.CallbackData) {
		app.OpenConfigDirectory()
	})
	FileMenu.AddText("ログの場所を開く", keys.CmdOrCtrl("."), func(_ *menu.CallbackData) {
		app.OpenLogDirectory()
	})

	CommandMenu := AppMenu.AddSubmenu("技")
	CommandMenu.AddText("Post", keys.CmdOrCtrl("Enter"), func(_ *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "call-post")
	})
	CommandMenu.AddText("ちくわ。", keys.OptionOrAlt("c"), func(_ *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "call-chikuwa")
	})
	CommandMenu.AddText("地震だ！", keys.OptionOrAlt("F9"), func(_ *menu.CallbackData) {
		app.Chikuwa("地震だ！")
	})
	CommandMenu.AddSubmenu("バージョン")
	CommandMenu.AddText("バージョン", keys.OptionOrAlt("v"), func(_ *menu.CallbackData) {
		app.Chikuwa("")
	})
	BlueskyCommandMenu := CommandMenu.AddSubmenu("Bluesky")
	BlueskyCommandMenu.AddText("偽招待コード生成", nil, func(_ *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "call-appendDummyInviteCode")
	})
	WindowMenu := AppMenu.AddSubmenu("Window")
	WindowMenu.AddText("中央に移動する", nil, func(_ *menu.CallbackData) {
		runtime.WindowCenter(app.ctx)
	})

	if app.environment.Platform == "darwin" {
		AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title: "まぜそば大陸",

		Width:  config.Window.Width,
		Height: config.Window.Height,

		AlwaysOnTop: config.Window.AlwaysOnTop,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless: true,
		// CSSDragProperty:  "windows",
		// CSSDragValue:     "1",

		BackgroundColour: &options.RGBA{R: 48, G: 128, B: 48, A: 0},

		OnStartup: app.startup,
		Menu:      AppMenu,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			// BackdropType:                      windows.Auto,
			BackdropType:                      windows.Acrylic,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			// Theme:                             windows.SystemDefault,
			// User messages that can be customised
			// Messages *windows.Messages
			// OnSuspend is called when Windows enters low power mode
			// OnSuspend func(),
			// OnResume is called when Windows resumes from low power mode
			// OnResume func(),
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				// Tekitou
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               true,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			// Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
