import { expect, it, describe, vi, afterEach, assert, beforeAll } from 'vitest'
import { dispatchInput, countGrapheme, countBytes } from './dispatchInput'
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
        assert.notEqual(resFail, '')
        expect(BrowserOpenURL).not.toHaveBeenCalled()
        expect(Post).not.toHaveBeenCalled()

    })

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
    ])('return length of message if input is usual messages', async ({input, expected}) => {
        let resOk = '', resFail = ''
        await dispatchInput(input, true).then((res) => { resOk = res }).catch((res) => { resFail = res })
        assert.equal(resOk, expected)
        assert.equal(resFail, '')
        expect(Post).not.toBeCalled() // not called if dryrun
    })
})

describe('countGrapheme', () => {
    it.each([
        { input: 'Hello, World', expected: 12 },
        { input: 'こんにちは', expected: 5 },
        { input: '👋', expected: 1 },
        { input: '👩‍👩‍👧‍👧', expected: 1 },
    ])('return displayed length of input', ({input, expected}) => {
        expect(countGrapheme(input)).toBe(expected)
    })
})

describe('countBytes', () => {
    it.each([
        { input: 'Hello, World', expected: 12 },
        { input: 'こんにちは', expected: 15 },
        { input: '👋', expected: 4 },
        { input: '👩‍👩‍👧‍👧', expected: 25 },
    ])('return length in UTF-8 bytes of input: $input', ({input, expected}) => {
        expect(countBytes(input)).toBe(expected)
    })
})
