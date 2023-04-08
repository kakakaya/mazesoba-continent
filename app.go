package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	// cliutil "github.com/bluesky-social/indigo/cmd/gosky/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// App struct
type App struct {
	config Config
	ctx    context.Context
	xrpcc  *xrpc.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Build xrpc.Client
	xrpcc := &xrpc.Client{
		Client: NewHttpClient(),
		Host:   "https://bsky.social", // only bsky.social is supported for now
		Auth:   &a.config.AuthInfo,
	}
	a.xrpcc = xrpcc
}

func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time and current time is %s!", name, time.Now().Format(time.RFC3339))
}

func (a *App) Post(text string) string {
	res, err := PostFeed(a.xrpcc, text)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}
	return res
}

func (a *App) Chikuwa() string {
	text := fmt.Sprintf("ちくわ。 %s", time.Now().Format(time.RFC3339))
	res, err := PostFeed(a.xrpcc, text)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}
	return res
}

