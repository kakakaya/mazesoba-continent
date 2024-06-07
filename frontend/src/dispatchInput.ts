
import {
    Post,
    Chikuwa,
    OpenConfigDirectory,
    OpenLogDirectory,
    GetContext,
} from '../wailsjs/go/main/App.js'

import {
    BrowserOpenURL, Quit,
} from '../wailsjs/runtime/runtime.js'

import { footers } from './stores.js'

const WAIT_FOR_INPUT_MESSAGE = "..."

export async function dispatchInput(input: string, dryRun: boolean = false): Promise<string> {
    // 1. /で始まる場合はコマンドとして処理
    // 2. 通常のメッセージの場合はそのままPost

    if (input.startsWith("/")) {
        input = input.trim()
        const args = input.split(" ");
        switch (args.at(0)) {
            case "/help":
                const helpTopic = args.at(1) || "";
                const helpArgs = args.slice(2);
                switch (helpTopic) {
                    case "":
                        return executeOrDryRun(dryRun, `READMEを開く`, helpCommand, ...helpArgs)
                    case "command":
                        return executeOrDryRun(dryRun, `スラッシュコマンドについて`, helpCommand, "command", ...helpArgs)
                    case "config":
                    case "setting":
                    case "settings":
                        return executeOrDryRun(dryRun, `設定について`, helpCommand, "config", ...helpArgs)
                    default:
                        return Promise.reject(dryRun ? WAIT_FOR_INPUT_MESSAGE : `😕「${helpTopic}ってなに？」`)
                }
            case "/open":
                const openTarget = args.at(1) || "";
                switch (openTarget) {
                    case "search":
                    case "s":
                        const searchArgs = args.slice(2);
                        return executeOrDryRun(dryRun, `検索：${searchArgs.join(" ")}`, searchCommand, ...searchArgs)
                    case "weather":
                        const addressArgs = args.slice(2);
                        return executeOrDryRun(dryRun, `天気を調べる：${addressArgs.join(" ")}`, weatherCommand, ...addressArgs)
                    case "profile":
                        const profileId = args.at(2) || "";
                        return executeOrDryRun(dryRun, `プロフィールを開く：${profileId}`, openProfileCommand, profileId)
                    case "config":
                    case "setting":
                    case "settings":
                        // Return "Opened config directory" message using executeOrDryRun
                        return executeOrDryRun(dryRun, "設定ファイルの場所を開く", () => OpenConfigDirectory().then(() => "Opened config directory"))
                    case "log":
                        // Return "Opened log directory" message using executeOrDryRun
                        return executeOrDryRun(dryRun, "ログの場所を開く", () => OpenLogDirectory().then(() => "Opened log directory"))
                    default:
                        // If dryRun is false, 
                        // return a string that describes openTarget is invalid, with suitable emoji.
                        // otherwise, return a string that app is waiting for next input, with suitable emoji.
                        return Promise.reject(dryRun ? WAIT_FOR_INPUT_MESSAGE : `😕「${openTarget}ってなに？」`)
                }
            case "/post":
                const postTarget = args.at(1) || "";
                switch (postTarget) {
                    case "chikuwa":
                    case "ckw":
                        const res = await executeOrDryRun(dryRun, `ちくわ。`, Chikuwa, "ちくわ。")
                        return res
                    case "earthquake":
                    case "eq":
                        return executeOrDryRun(dryRun, `地震だ！`, Chikuwa, "地震だ！")
                    case "version":
                    case "ver":
                        let AppContext: { version: string } = { version: "" };
                        return await GetContext()
                            .then((context) => {
                                AppContext = JSON.parse(context);
                                const Version = AppContext.version;
                                return executeOrDryRun(dryRun, `バージョン情報： ${Version}`, Chikuwa, `まぜそば大陸バージョン${Version}が浮上中`);
                            })
                            .catch((err) => {
                                return Promise.reject(err);
                            });
                    default:
                        return Promise.reject(dryRun ? WAIT_FOR_INPUT_MESSAGE : `😕「${postTarget}ってなに？」`)
                }
            case "/pizza":
                const pizzaAddress = args.slice(1).join("");
                // Check if pizzaAddress is 7 digits using regex
                if (!/^\d{7}$/.test(pizzaAddress)) {
                    return Promise.reject(dryRun ?  `ピザを注文する：〒${pizzaAddress}` : "😕「郵便番号を7桁の数字で入力してね」")
                }
                return executeOrDryRun(dryRun, `ピザを注文する：〒${pizzaAddress}`, pizzaCommand, pizzaAddress)
            case "/set":
                const setTarget = args.at(1) || "";
                switch (setTarget) {
                    case "f":
                    case "footer":
                        const newFooters = args.slice(2)
                        return executeOrDryRun(dryRun, `フッターを追加する：${newFooters.join(" ")}`, () => {
                            footers.set(newFooters);
                            return Promise.resolve("追加しました");                        

                        })  // FIXME
                    default:
                        return Promise.reject(dryRun ? "設定する対象を入力：(footer)" : `😕「${setTarget}ってなに？」`)
                }
                return ""
            case "/reset":
                const resetTarget = args.at(1) || "";
                switch (resetTarget) {
                    case "f":
                    case "footer":
                        return executeOrDryRun(dryRun, `フッターをリセットする`, () => {
                            footers.set([]);
                            return Promise.resolve("リセットしました");
                        });
                    default:
                        return Promise.reject(dryRun ? "リセットする対象を入力：(footer)" : `😕「${resetTarget}ってなに？」`)
                }
            case "/quit":
                return executeOrDryRun(dryRun, "沈没", () => {
                    Quit();
                    return Promise.resolve("沈没！");
                });
            case "/mzsb":
                const mzsbTarget = args.at(1) || "";
                switch (mzsbTarget) {
                    case "egosearch":
                    case "egs":
                        return executeOrDryRun(dryRun, "エゴサーチ", searchCommand, "まぜそば大陸")
                    default:
                        return Promise.reject(dryRun ? WAIT_FOR_INPUT_MESSAGE : `😕「${mzsbTarget}ってなに？」`)
                }
            default:
                if (dryRun) {
                    return Promise.reject(WAIT_FOR_INPUT_MESSAGE);
                } else {
                    return Promise.reject(`😕「${args.at(0)}の仕方がわからないよ」`)
                }
        }
    }
    if (dryRun) {
        return Promise.resolve(countGrapheme(input).toString());
    }
    return Post(input)
}

// A function which executes and returns given command if dryRun = false
// otherwise returns a string which describes what the command will do
function executeOrDryRun(dryRun: boolean, description: string, command: (...args: string[]) => Promise<string>, ...args: string[]): Promise<string> {
    if (dryRun) {
        return Promise.resolve(description);
    } else {
        return command(...args);
    }
}

export function helpCommand(...topics: string[]): Promise<string> {
    const README = 'https://github.com/kakakaya/mazesoba-continent/blob/main/README.md'
    const COMMAND = 'https://github.com/kakakaya/mazesoba-continent/blob/main/docs/SLASH_COMMAND.md'
    const CONFIG = 'https://github.com/kakakaya/mazesoba-continent/blob/main/docs/CONFIG.md'
    if (topics.length == 0) {
        BrowserOpenURL(README)
        return Promise.resolve(`Open: ${README}`)
    }
    switch (topics.at(0)) {
        case 'command':
            BrowserOpenURL(COMMAND)
            return Promise.resolve(`Open: ${COMMAND}`)
        case 'config':
            BrowserOpenURL(CONFIG)
            return Promise.resolve(`Open: ${CONFIG}`)
        default:
            return Promise.reject(`😕「${topics}ってなに？」`)
    }
}

export function searchCommand(...searchArgs: string[]): Promise<string> {
    if (searchArgs.length < 1) {
        return Promise.reject("😕検索ワードを指定してね")
    }
    const params = encodeURIComponent(searchArgs.join(' '))
    const url = `https://bsky.app/search?q=${params}`
    BrowserOpenURL(url)
    return Promise.resolve(`Open: ${url}`)
}

export function weatherCommand(...addressArgs: string[]): Promise<string> {
    if (addressArgs.length < 1) {
        return Promise.reject("😕地名を指定してね")
    }
    const params = encodeURIComponent(addressArgs.join(' '))
    const url = `https://tenki.jp/search/?keyword=${params}`

    BrowserOpenURL(url)
    return Promise.resolve(`Open: ${url}`)
}

export function pizzaCommand(pizzaAddress: string): Promise<string> {
    // Order Pizzahut pizza
    const url = `https://www.pizzahut.jp/localization/${pizzaAddress}`
    BrowserOpenURL(url)
    return Promise.resolve(`Open: ${url}`)
}

export function openProfileCommand(profileId: string): Promise<string> {
    // TODO: Merge this function with searchCommand
    const url = `https://bsky.app/profile/${profileId}`
    BrowserOpenURL(url)
    return Promise.resolve(`Open: ${url}`)
}

export function countGrapheme(input: string): number {
    const segmenter = new Intl.Segmenter("ja", {
        granularity: "grapheme"
    });
    return [...segmenter.segment(input)].length;
}

export function countBytes(input: string): number {
    return new TextEncoder().encode(input).length
}
