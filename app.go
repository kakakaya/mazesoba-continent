package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	// cliutil "github.com/bluesky-social/indigo/cmd/gosky/util"
	"github.com/adrg/xdg"
	"github.com/bluesky-social/indigo/xrpc"
)

// App struct
type App struct {
	config Config
	ctx    context.Context
	xrpcc  *xrpc.Client
	logger *log.Logger
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
	var auth *xrpc.AuthInfo

	cred := a.config.Credential
	auth, err := createSession(cred.Host, cred.Identifier, cred.Password)

	if err != nil {
		a.logger.Fatalf("Failed creating session, bad identifier or password?: %v", err)
	}

	xrpcc := &xrpc.Client{
		Client: NewHttpClient(),
		Host:   cred.Host,
		Auth:   auth,
	}
	a.xrpcc = xrpcc
}

// createSession calls "/xrpc/com.atproto.server.createSession" and returns xrpc.AuthInfo.
func createSession(host string, identifier string, password string) (*xrpc.AuthInfo, error) {
	payload := fmt.Sprintf(`{"identifier": "%s", "password": "%s"}`, identifier, password)
	if host == "" {
		return nil, fmt.Errorf("Error: host not set. Check [Credential] section in config file.")
	}

	url := fmt.Sprintf("%s/xrpc/com.atproto.server.createSession", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP request:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: Unexpected status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}
	var auth xrpc.AuthInfo
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON response: %v", err)
	}

	return &auth, nil
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
	res, err := BskyFeedPost(a.xrpcc, text)
	if err != nil {
		errs := fmt.Sprintf("Posting error: %s", err.Error())
		a.logger.Println(errs)
		return errs

	}
	return res
}

func (a *App) Chikuwa(text string) string {
	return a.Post(fmt.Sprintf("%s %s", text, time.Now().Format(time.RFC3339)))
}

func (a *App) OpenConfigDirectory() error {
	f, _ := xdg.ConfigFile("mazesoba-continent")
	return openDirectory(f)
}

func (a *App) OpenLogDirectory() error {
	f, _ := xdg.CacheFile("mazesoba-continent")
	return openDirectory(f)
}

func (a *App) GetVersion() string {
	return Version
}
