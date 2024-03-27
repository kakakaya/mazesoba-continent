<script lang="ts">
    import {
        DispatchCommand,
    } from '../wailsjs/go/main/App.js'
    import {
        EventsOn,
        EventsEmit,
        WindowCenter,
        WindowHide,
        WindowShow,
        Quit,
        LogInfo,
    } from '../wailsjs/runtime/runtime.js'

    import {
        ConvertRichUnicode
    } from './topping/unicode.ts'

    let text: string
    let charCounter = 0
    let placeholder = Math.random() > 0.5 ? "最近どう？" : "どう最近？"

    function post() {
        DispatchCommand(text).then(result => {
            placeholder = result
            text = ""
        })
    }

    function clearText() {
        text = ""
    }

    function unicode() {
        text = ConvertRichUnicode(text, 'sansBold')
    }

    function countGrapheme(input: string) {
        const segmenter = new Intl.Segmenter("ja", {
            granularity: "grapheme"
        });
        return [...segmenter.segment(input)].length;
    }

    function onChange() {
        charCounter = countGrapheme(text)
    }


    // Setup Events
    EventsOn("call-clearText", () => {
        clearText()
    })
    EventsOn("call-post", () => {
        post()
    })
</script>

<main style="--wails-draggable:drag">
    <textarea autocomplete="off" bind:value={text} on:input={onChange} placeholder={placeholder} id="inputbox" />
    <p id="char-counter" class="char-count">{charCounter}</p>
</main>
