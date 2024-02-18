package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) DispatchCommand(text string) string {
	a.logger.Debug("dispatchcommand", "text", text)
	words := strings.Fields(text)
	if len(words) < 1 {
		return a.Post(text)
	}
	action := words[0]
	switch action {
	case "/help":
		{
			if len(words) < 2 {
				OpenURL("https://github.com/kakakaya/mazesoba-continent/blob/main/README.md")
			}
			switch words[1] {
			case "config":
				{
					OpenURL("https://github.com/kakakaya/mazesoba-continent/tree/main/docs/CONFIG.md")
				}
			case "command":
				{
					OpenURL("https://github.com/kakakaya/mazesoba-continent/tree/main/docs/SLASH_COMMAND.md")
				}
			default:
				{
					OpenURL("https://github.com/kakakaya/mazesoba-continent/blob/main/README.md")
				}
			}
			return ""
		}
	case "/open":
		{
			if len(words) < 2 {
				return "🍜「何を開く？アジ？」"
			}
			switch words[1] {
			case "config":
				{
					a.OpenConfigDirectory()
				}
			case "log":
				{
					a.OpenLogDirectory()
				}
			case "profile":
				{
					handle_or_did := words[2]
					OpenURL(fmt.Sprintf("https://bsky.app/profile/%s", handle_or_did))
				}
			case "search", "s":
				{
					if len(words) < 3 {
						return "🍜「検索ワードを指定してね」"
					}
					req, err := http.NewRequest("GET", "https://bsky.app/search", nil)
					if err != nil {
						a.logger.Warn("error making HTTP request", "error", err)
						return ""
					}

					q := req.URL.Query()
					q.Add("q", strings.Join(words[2:], " "))
					req.URL.RawQuery = q.Encode()

					OpenURL(req.URL.String())
				}
			default:
				{
					OpenURL(words[1]) // NOTE: TBC
				}
			}
			return ""
		}
	case "/post":
		{
			if len(words) < 2 {
				return "🍜「サブコマンドを指定してね。 /post version とか」"
			}
			switch words[1] {
			case "chikuwa", "ckw":
				{
					return a.Chikuwa("ちくわ。")
				}
			case "earthquake", "eq":
				{
					return a.Chikuwa("地震だ！")
				}
			case "version", "ver":
				{
					return a.Chikuwa(fmt.Sprintf("まぜそば大陸バージョン%sが浮上中", Version))
				}
			}
		}
	case "/quit":
		{
			runtime.Quit(a.ctx)
		}
	case "/mzsb": // hidden functions
		{
			if len(words) < 2 {
				return ""
			}
			switch words[1] {
			case "egosearch", "egs":
				{
					OpenURL("https://bsky.app/search?q=%E3%81%BE%E3%81%9C%E3%81%9D%E3%81%B0%E5%A4%A7%E9%99%B8")
				}
			}
			return ""
		}
	default:
		{
			return a.Post(text)
		}
	}
	return ""
}
