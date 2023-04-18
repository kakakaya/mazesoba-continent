package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/adrg/xdg"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/exp/slog"
	"github.com/gen2brain/beeep"
)

// App struct
type App struct {
	config      Config
	ctx         context.Context
	xrpcc       *xrpc.Client
	logger      *slog.Logger
	environment runtime.EnvironmentInfo
}

// dialogResults is used for casting Dialog's response to boolean. Note that keys are lowercase.
var dialogResults = map[string]bool{
	"ok":        true,
	"cancel":    false,
	"abort":     false,
	"retry":     false,
	"ignore":    false,
	"yes":       true,
	"no":        false,
	"try again": false,
	"continue":  false,
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.environment = runtime.Environment(ctx)
	if a.environment.BuildType == "dev" {
		// Use stdout in dev mode
		a.logger = slog.New(slog.NewTextHandler(os.Stdout))
	}
	a.logger.Info("Startup", "environment", a.environment)

	// Build xrpc.Client
	var auth *xrpc.AuthInfo
	cred := a.config.Credential
	if cred.Host == "" || cred.Identifier == "" || cred.Password == "" {
		a.logger.Warn("Credentials not set!")
		result, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.QuestionDialog,
			Title:         "認証情報がないよ",
			Message:       "config.tomlにIDとパスワードを入力してね。\nこのまま沈没するけど、設定ファイルの場所を開く？",
			DefaultButton: "Yes",
		})
		if err != nil {
			a.logger.Warn("Error creds not set dialog", "error",  err)
		}
		a.logger.Info(result)
		if dialogResults[strings.ToLower(result)] {
			a.OpenConfigDirectory()
		}
		a.logger.Warn("Missing credentials")
	}

	auth, err := createSession(cred.Host, cred.Identifier, cred.Password)

	if err != nil {
		// If auth failed,
		result, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.QuestionDialog,
			Title:         "認証に失敗しました",
			Message:       fmt.Sprintf("%sへの認証に失敗しました。\nIDとパスワードを確認してください。\nこのまま沈没するけど、設定ファイルの場所を開く？", cred.Host),
			DefaultButton: "Yes",
		})

		if err != nil {
			a.logger.Warn("Error creds not set dialog", "error", err)
		}
		a.logger.Info(result)
		if dialogResults[strings.ToLower(result)] {
			a.OpenConfigDirectory()
		}

		a.logger.Warn("Failed creating session, bad identifier or password?", "error", err)
		panic(err)
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

func (a *App) Post(text string) string {
	a.logger.Info("Post", "text", text)
	err := beeep.Notify("まぜそば大陸", fmt.Sprintf("BskyFeedPost: %s", text), "")

	if a.environment.BuildType == "dev" {
		return "<MOCK URI>"
	}
	res, err := BskyFeedPost(a.xrpcc, text)
	if err != nil {
		errs := fmt.Sprintf("Posting error: %s", err.Error())
		a.logger.Error("Post: Error on BskyFeedPost", "error", errs)
		return errs
	}
	a.logger.Info("Posted", "result", res)
	return res
}

func (a *App) Chikuwa(text string) string {
	return a.Post(fmt.Sprintf("%s %s", text, time.Now().Format(time.RFC3339)))
}

func (a *App) OpenConfigDirectory() error {
	f, _ := xdg.ConfigFile("mazesoba-continent")
	a.logger.Info("OpenConfigDirectory", "path", f)
	return openDirectory(f)
}

func (a *App) OpenLogDirectory() error {
	f, _ := xdg.CacheFile("mazesoba-continent")
	a.logger.Info("OpenLogDirectory", "path", f)
	return openDirectory(f)
}

func (a *App) GetVersion() string {
	return Version
}
