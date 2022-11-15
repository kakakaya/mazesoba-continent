package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	// "github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "まぜそば大陸",

		Width:  300,
		Height: 200,

		AlwaysOnTop: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		// CSSDragProperty:  "windows",
		// CSSDragValue:     "1",

		BackgroundColour: &options.RGBA{R: 48, G: 128, B: 48, A: 0},

		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
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
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
