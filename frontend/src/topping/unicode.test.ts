import { expect, test, describe } from 'vitest'
import { ConvertRichUnicode } from "./unicode";

describe.each([
    { mode: 'circled', expected: 'Ⓗⓔⓛⓛⓞ Ⓦⓞⓡⓛⓓ!' },
    { mode: 'doublestruck', expected: 'ℍ𝕖𝕝𝕝𝕠 𝕎𝕠𝕣𝕝𝕕!' },
    { mode: 'frakturBold', expected: '𝕳𝖊𝖑𝖑𝖔 𝖂𝖔𝖗𝖑𝖉!' },
    { mode: 'fraktur', expected: 'ℌ𝔢𝔩𝔩𝔬 𝔚𝔬𝔯𝔩𝔡!' },
    { mode: 'monospace', expected: '𝙷𝚎𝚕𝚕𝚘 𝚆𝚘𝚛𝚕𝚍!' },
    { mode: 'sansBold', expected: '𝗛𝗲𝗹𝗹𝗼 𝗪𝗼𝗿𝗹𝗱!' },
    { mode: 'sansBoldItalic', expected: '𝙃𝙚𝙡𝙡𝙤 𝙒𝙤𝙧𝙡𝙙!' },
    { mode: 'sans', expected: '𝖧𝖾𝗅𝗅𝗈 𝖶𝗈𝗋𝗅𝖽!' },
    { mode: 'sansItalic', expected: '𝘏𝘦𝘭𝘭𝘰 𝘞𝘰𝘳𝘭𝘥!' },
    { mode: 'scriptBold', expected: '𝓗𝓮𝓵𝓵𝓸 𝓦𝓸𝓻𝓵𝓭!' },
    { mode: 'script', expected: 'ℋℯ𝓁𝓁ℴ 𝒲ℴ𝓇𝓁𝒹!' },
    { mode: 'serifBold', expected: '𝐇𝐞𝐥𝐥𝐨 𝐖𝐨𝐫𝐥𝐝!' },
    { mode: 'serifBoldItalic', expected: '𝑯𝒆𝒍𝒍𝒐 𝑾𝒐𝒓𝒍𝒅!' },
    { mode: 'squaredBlack', expected: '🅷🅴🅻🅻🅾 🆆🅾🆁🅻🅳!' },
    { mode: 'squared', expected: '🄷🄴🄻🄻🄾 🅆🄾🅁🄻🄳!' },
])('ConvertRichUnicode("Hello World!", $mode) returns $expected', ({ mode, expected }) => {
    const base = 'Hello World!'
    test(`returns ${expected}`, () => {
        expect(ConvertRichUnicode(base, mode)).toBe(expected)

    })
})
