package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"golang.org/x/exp/slog"

	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

const Version = "v7"

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
		log.Fatalf("Can't open config file: %s", configPath)
	}
	logFile, err := os.Create(logPath)
	if err != nil {
		log.Fatalf("Can't open config file: %s", configPath)
	}

	defer logFile.Close()
	app.logger = slog.NewLogLogger(slog.NewJSONHandler(logFile), slog.LevelInfo)

	config, err := loadOrCreateConfig(configPath)
	if err != nil {
		app.logger.Fatal(err)
	}

	app.config = config

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
