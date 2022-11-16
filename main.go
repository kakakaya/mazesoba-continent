package main

import (
	"bytes"
	"embed"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"

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

	configFile, _ := xdg.ConfigFile("mazesoba-continent/config.toml2")
	var config Config
	a, err := toml.DecodeFile(configFile, &config)
	log.Println(a)
	if err != nil {
		log.Printf("Can't open or decode config file: %s", configFile)
		window := windowConfig{
			Width:       400,
			Height:      200,
			AlwaysOnTop: true,
			Transparent: true,
		}
		defaultConfig := Config{
			Window: window,
		}
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(defaultConfig)
		if err != nil {
			log.Println(err)
		}
		log.Printf(
			"========Default config toml========\n%s================================",
			buf.String(),
		)
		log.Printf("%+v\n", defaultConfig)
		config = defaultConfig
	}
	log.Fatal("bye bye")

	// Create application with options
	err = wails.Run(&options.App{
		Title: "まぜそば大陸",

		Width:  300,
		Height: 200,

		AlwaysOnTop: true,
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
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
