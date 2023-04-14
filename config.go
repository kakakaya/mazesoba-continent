package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
)

const configPath = "mazesoba-continent/config.toml"

type Config struct {
	General    generalConfig
	Window     windowConfig
	Credential credConfig
	Kakikomi   kakikomiConfig
}

type generalConfig struct {
	// TBW
}

type windowConfig struct {
	Width       int
	Height      int
	AlwaysOnTop bool
	Transparent bool
}

type credConfig struct {
	Identifier string
	Password   string
	Host       string
}

type kakikomiConfig struct {
	// TBW
	// a.k.a kakikomi.txt
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
		log.Println(
			"",
			"========Wrote default config toml========",
			buf.String(),
			"=========================================",
		)
		ioutil.WriteFile(configFile, buf.Bytes(), 0644)
		config = defaultConfig
	}
	return config, nil
}
