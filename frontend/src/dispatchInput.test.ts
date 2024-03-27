import { expect, it, describe, vi, afterEach, assert, beforeAll } from 'vitest'
import { dispatchInput } from './dispatchInput'
import {
    Post,
} from '../wailsjs/go/main/App.js'
import {
    BrowserOpenURL,
} from '../wailsjs/runtime/runtime.js'


vi.mock('../wailsjs/go/main/App.js', () => ({
    Post: vi.fn(),
}))
vi.mock('../wailsjs/runtime/runtime.js', () => ({
    BrowserOpenURL: vi.fn(),
}))


describe('dispatchInput', () => {
    afterEach(() => {
        vi.resetAllMocks();
    });

    it('Post if input is usual messages', () => {
        const msg = 'Hello, World'
        dispatchInput(msg)
        expect(Post).toHaveBeenCalledWith(msg)
    })

    it.each([
        { input: "/help", helpUrl: 'https://github.com/kakakaya/mazesoba-continent/blob/main/README.md' },
        { input: "/help command", helpUrl: 'https://github.com/kakakaya/mazesoba-continent/blob/main/docs/SLASH_COMMAND.md' },
        { input: "/help config", helpUrl: 'https://github.com/kakakaya/mazesoba-continent/blob/main/docs/CONFIG.md' },
    ])(`Open help page if input=$input`, async ({ input, helpUrl }) => {
        let resOk = '', resFail = ''
        await dispatchInput(input).then((res) => { resOk = res }).catch((res) => { resFail = res })
        assert.equal(resOk, `Open: ${helpUrl}`)
        assert.equal(resFail, '')
        expect(BrowserOpenURL).toHaveBeenCalledWith(helpUrl)
        expect(Post).not.toHaveBeenCalled()
    })

    it('Show error message if unknown help topic', async () => {
        const input = '/help unknown-topic'
        let resOk = '', resFail = ''
        await dispatchInput(input).then((res) => { resOk = res }).catch((res) => { resFail = res })
        assert.equal(resOk, '')
        assert.equal(resFail, 'Error: Unknown help topic')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toHaveBeenCalled()

    })

    it.each([
        // { input: "/open search foo", expected: "https://bsky.app/search?q=foo" },
        // { input: " /open   search　まぜそば大陸　", expected: "https://bsky.app/search?q=まぜそば大陸" },
    ])(`Open search page if input=$input`, ({ input, expected }) => {
        expect(BrowserOpenURL).toHaveBeenCalledWith(expected)
    })
})

