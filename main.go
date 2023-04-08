package main

import (
	"bytes"
	"embed"
	"io/ioutil"
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

	configPath := "mazesoba-continent/config.toml"
	config, err := loadOrCreateConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	app.config = config
	log.Println(app.config)

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
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// loadOrCreateConfig ...
func loadOrCreateConfig(configPath string) (Config, error) {
	var config Config
	configFile, err := xdg.ConfigFile(configPath)
	if err != nil {
		log.Fatalf("Can't open config file: %s", configPath)
	}
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		log.Printf("Config file not found, initializing: %s", configFile)
		defaultConfig := Config{
			Window: windowConfig{
				Width:       400,
				Height:      200,
				AlwaysOnTop: true,
				Transparent: true,
			},
		}
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(defaultConfig)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf(
			"========Wrote default config toml========\n%s\n================================",
			buf.String(),
		)
		ioutil.WriteFile(configFile, buf.Bytes(), 0644)
		config = defaultConfig
	}
	return config, nil
}
