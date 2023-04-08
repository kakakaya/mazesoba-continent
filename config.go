package main

import (
	"github.com/bluesky-social/indigo/xrpc"
)

type Config struct {
	General  generalConfig
	Window   windowConfig
	AuthInfo xrpc.AuthInfo
	Kakikomi kakikomiConfig
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

type kakikomiConfig struct {
	// TBW
	// a.k.a kakikomi.txt
}
