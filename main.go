package main

import (
	"embed"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/adrg/xdg"

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

const Version = "v16"

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
	FileMenu.AddText("沈没", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
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

		// Windows allows A: 0 or 255 only, therefore we'll use 0 for simple.
		BackgroundColour: &options.RGBA{R: 48, G: 48, B: 48, A: 0},

		OnStartup:     app.startup,
		OnBeforeClose: app.beforeClose,
		Menu:          AppMenu,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			// BackdropType:         windows.None,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: true, // this makes true transparent
			Theme:                             windows.SystemDefault,
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
		app.logger.Warn("error on app close", "error", err.Error())
	}
}
