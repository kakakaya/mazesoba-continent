import { expect, test, describe, vi } from 'vitest'
import { dispatchInput } from './dispatchInput'

describe.each([
    { input: '', expected: '' },
    { input: 'hello, world', expected: '' },
    { input: '', expected: '' },
])('dispatchInput($input) returns $expected', ({ input, expected }) => {
    test(`returns ${expected}`, () => {
        expect(dispatchInput(input)).toBe(expected)
    })
})

