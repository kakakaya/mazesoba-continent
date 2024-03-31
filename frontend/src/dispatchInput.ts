
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
                return executeOrDryRun(dryRun, `Open help for ${helpArgs.join(" ")}`, helpCommand, ...helpArgs)
            case "/open":
                const openTarget = args.at(1) || "";
                switch (openTarget) {
                    case "search":
                        const searchArgs = args.slice(2);
                        return executeOrDryRun(dryRun, `Search ${searchArgs.join(" ")}`, searchCommand, ...searchArgs)
                    case "weather":
                        const addressArgs = args.slice(2);
                        return executeOrDryRun(dryRun, `${addressArgs.join(" ")}の天気を調べる`, weatherCommand, ...addressArgs)
                    default:
                        // If dryRun is false, 
                        // return a string that describes openTarget is invalid, with suitable emoji.
                        // otherwise, return a string that app is waiting for next input, with suitable emoji.
                        if (dryRun) {
                            return Promise.reject(`...`)
                        } else {
                            return Promise.reject(`😕「${openTarget}ってなに？」`)
                        }
                }
            default:
                if (dryRun) {
                    return Promise.reject(`...`);
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


export function countGrapheme(input: string): number {
    const segmenter = new Intl.Segmenter("ja", {
        granularity: "grapheme"
    });
    return [...segmenter.segment(input)].length;
}

export function countBytes(input: string): number {
    return new TextEncoder().encode(input).length
}
