
import {
    Post,
} from '../wailsjs/go/main/App.js'

import {
    BrowserOpenURL,
} from '../wailsjs/runtime/runtime.js'

export function dispatchInput(input: string, dryRun: boolean = false): Promise<string> {
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
                break;
            default:
                break;
        }
        return Promise.reject("Unknown command")
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
            break;
    }
    return Promise.reject("Error: Unknown help topic")
}
