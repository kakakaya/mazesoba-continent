
import {
    Post,
} from '../wailsjs/go/main/App.js'

import {
    BrowserOpenURL,
} from '../wailsjs/runtime/runtime.js'

export function dispatchInput(input: string, dryRun: boolean = false): Promise<string> {
    // 1. /で始まる場合はコマンドとして処理
    // 2. 通常のメッセージの場合はそのままPost

    if (input.startsWith("/")) {
        input = input.trim()
        const args = input.split(" ");
        switch (args.at(0)) {
            case "/help":
                const helpArgs = args.slice(1);
                if (dryRun) {
                    return Promise.resolve(`Open help for ${helpArgs.join(" ")}`);
                } else {
                    return helpCommand(...helpArgs)
                }
            case "/open":
                const openTarget = args.at(1)
                switch (openTarget) {
                    case "search":
                        const searchArgs = args.slice(2);
                        if (dryRun) {
                            return Promise.resolve(`Search ${searchArgs.join(" ")}`);
                        } else {
                            return searchCommand(...searchArgs)
                        }
                    case "weather":
                        const addressArgs = args.slice(2);
                        if (dryRun) {
                            return Promise.resolve(`${addressArgs.join(" ")}の天気を開く`);
                        } else {
                            return weatherCommand(...addressArgs)
                        }
                    default:
                        return Promise.reject(`😕「${openTarget}の開き方がわからないよ」`)
                }
            default:
                return Promise.reject(`😕「${args.at(0)}の仕方がわからないよ」`)
        }
    }
    if (dryRun) {
        return Promise.resolve(countGrapheme(input).toString());
    }
    return Post(input)
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
            break;
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


function countGrapheme(input: string) {
    const segmenter = new Intl.Segmenter("ja", {
        granularity: "grapheme"
    });
    return [...segmenter.segment(input)].length;
}
