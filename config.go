package main

type Config struct {
	Window     windowConfig
	Credential credConfig
}

type windowConfig struct {
	Width       int
	Height      int
	AlwaysOnTop bool
	Transparent bool
}

type credConfig struct {
	// TBW
}
