import { expect, it, describe, vi, afterEach, assert, beforeAll } from 'vitest'
import { dispatchInput, countGrapheme, countBytes, weatherCommand, searchCommand } from './dispatchInput'
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

    // Section for /help commands
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
        assert.notEqual(resFail, '')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toHaveBeenCalled()

    })

    // Section for /open commands
    it.each([
        { input: "/open search foo", expected: "https://bsky.app/search?q=foo" },
        { input: "/open search まぜそば大陸", expected: "https://bsky.app/search?q=%E3%81%BE%E3%81%9C%E3%81%9D%E3%81%B0%E5%A4%A7%E9%99%B8" },
        { input: "/open search two words", expected: "https://bsky.app/search?q=two%20words" },
    ])(`Open search page if input=$input`, async ({ input, expected }) => {
        let resOk = '', resFail = ''
        await dispatchInput(input).then((res) => { resOk = res }).catch((res) => { resFail = res })
        expect(resOk).toContain(expected)
        assert.equal(resFail, '')
        expect(BrowserOpenURL).toHaveBeenCalledWith(expected)
    })

    it.each([
        { input: "/open weather 東京", expected: "https://tenki.jp/search/?keyword=%E6%9D%B1%E4%BA%AC" },
        { input: "/open weather 100-0001", expected: "https://tenki.jp/search/?keyword=100-0001" },
    ])(`Open weather page if input=$input`, async ({ input, expected }) => {
        let resOk = '', resFail = ''
        await dispatchInput(input).then((res) => { resOk = res }).catch((res) => { resFail = res })
        expect(resOk).toContain(expected)
        assert.equal(resFail, '')
        expect(BrowserOpenURL).toHaveBeenCalledWith(expected)
    })

    it('Rejects if open target is not recognized', async () => {
        const input = '/open unknown-target'
        let resOk = '', resFail = ''
        await dispatchInput(input).then((res) => { resOk = res }).catch((res) => { resFail = res })
        assert.equal(resOk, '')
        assert.notEqual(resFail, '')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toHaveBeenCalled()
    })

    it('Rejects if command is not recognized', async () => {
        const input = '/unknown-command'
        let resOk = '', resFail = ''
        await dispatchInput(input).then((res) => { resOk = res }).catch((res) => { resFail = res })
        assert.equal(resOk, '')
        assert.notEqual(resFail, '')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toHaveBeenCalled()
    })
})

describe('dispatchInput with dryrun', () => {
    afterEach(() => {
        vi.resetAllMocks();
    });

    it.each([
        { input: 'Hello, World', expected: '12' },
        { input: 'こんにちは', expected: '5' },
        { input: '👋', expected: '1' },
        { input: '👩‍👩‍👧‍👧', expected: '1' },
    ])('Returns length of message if input is usual messages', async ({ input, expected }) => {
        let resOk = '', resFail = ''
        await dispatchInput(input, true).then((res) => { resOk = res }).catch((res) => { resFail = res })
        assert.equal(resOk, expected)
        assert.equal(resFail, '')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toBeCalled() // not called if dryrun
    })

    it.each([
        // { input: "/help", expected: 'READMEを開く' },
        // { input: "/help command", expected: 'スラッシュコマンドのヘルプを開く' },
        // { input: "/help config", expected: '設定のヘルプを開く' },
        { input: "/open weather 東京", expected: '東京の天気を調べる' },
    ])('Returns help message if dryRun is true', async ({ input, expected }) => {
        let resOk = '', resFail = ''
        await dispatchInput(input, true).then((res) => { resOk = res }).catch((res) => { resFail = res })
        expect(resOk).toContain(expected)
        assert.equal(resFail, '')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toHaveBeenCalled()
    })

    it.each([
        // { input: "/help unknown" },
        { input: "/open unknown" },
        { input: '/unknown-command' },
    ])('Rejects if unknown command or argumnet, input=$input', async ({ input }) => {
        let resOk = '', resFail = ''
        await dispatchInput(input, true).then((res) => { resOk = res }).catch((res) => { resFail = res })
        assert.equal(resOk, '')
        assert.notEqual(resFail, '')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toHaveBeenCalled()
    })
})

describe('countGrapheme', () => {
    it.each([
        { input: 'Hello, World', expected: 12 },
        { input: 'こんにちは', expected: 5 },
        { input: '👋', expected: 1 },
        { input: '👩‍👩‍👧‍👧', expected: 1 },
    ])('return displayed length of input', ({ input, expected }) => {
        expect(countGrapheme(input)).toBe(expected)
    })
})

describe('countBytes', () => {
    it.each([
        { input: 'Hello, World', expected: 12 },
        { input: 'こんにちは', expected: 15 },
        { input: '👋', expected: 4 },
        { input: '👩‍👩‍👧‍👧', expected: 25 },
    ])('return length in UTF-8 bytes of input: $input', ({ input, expected }) => {
        expect(countBytes(input)).toBe(expected)
    })
})

describe('searchCommand', () => {
    it('searchCommand rejects when no search arguments are provided', async () => {
        try {
            await searchCommand();
        } catch (error) {
            expect(error).toBe('😕検索ワードを指定してね');
        }
        expect(BrowserOpenURL).not.toHaveBeenCalled();
    });
})

describe('weatherCommand', () => {
    it('weatherCommand rejects when no address arguments are provided', async () => {
        try {
            await weatherCommand();
        } catch (error) {
            expect(error).toBe('😕地名を指定してね');
        }
        expect(BrowserOpenURL).not.toHaveBeenCalled();
    });
})
