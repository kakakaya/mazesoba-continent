package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"time"

	"github.com/adrg/xdg"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	config Config
	ctx    context.Context
	xrpcc  *xrpc.Client
	logger *log.Logger
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
	// Build xrpc.Client
	var auth *xrpc.AuthInfo

	cred := a.config.Credential
	if cred.Host == "" || cred.Identifier == "" || cred.Password == "" {
		a.logger.Println("Credentials not set!")
		result, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.QuestionDialog,
			Title:         "認証情報がないよ",
			Message:       "config.tomlにIDとパスワードを入力してね。\nこのまま沈没するけど、設定ファイルの場所を開く？",
			DefaultButton: "Yes",
		})
		if err != nil {
			a.logger.Fatalln("Error creds not set dialog: ", err)
		}
		a.logger.Println(result)
		if dialogResults[strings.ToLower(result)] {
			a.OpenConfigDirectory()
		}
		a.logger.Fatalln("Missing credentials")
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
			a.logger.Fatalln("Error creds not set dialog: ", err)
		}
		a.logger.Println(result)
		if dialogResults[strings.ToLower(result)] {
			a.OpenConfigDirectory()
		}

		a.logger.Fatalf("Failed creating session, bad identifier or password?: %v", err)
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
	a.logger.Println("Post: ", text)
	res, err := BskyFeedPost(a.xrpcc, text)
	if err != nil {
		errs := fmt.Sprintf("Posting error: %s", err.Error())
		a.logger.Println(errs)
		return errs
	}
	a.logger.Println("Post result: ", res)
	return res
}

func (a *App) Chikuwa(text string) string {
	return a.Post(fmt.Sprintf("%s %s", text, time.Now().Format(time.RFC3339)))
}

func (a *App) OpenConfigDirectory() error {
	f, _ := xdg.ConfigFile("mazesoba-continent")
	a.logger.Println("OpenConfigDirectory: ", f)
	return openDirectory(f)
}

func (a *App) OpenLogDirectory() error {
	f, _ := xdg.CacheFile("mazesoba-continent")
	a.logger.Println("OpenLogDirectory: ", f)
	return openDirectory(f)
}

func (a *App) GetVersion() string {
	return Version
}
