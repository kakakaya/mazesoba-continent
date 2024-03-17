<script>
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

    let text
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

    function countGrapheme(string) {
        const segmenter = new Intl.Segmenter("ja", {
            granularity: "grapheme"
        });
        return [...segmenter.segment(string)].length;
    }

    function onChange() {
        charCounter = countGrapheme(text)
    }

    function setBackgroundColor(r, g, b, a) {
        const newBackgroundStyle = `background-color: rgba(${r}, ${g}, ${b}, ${a});`
        LogInfo(newBackgroundStyle)
    }

    // Setup Events
    EventsOn("call-clearText", () => {
        clearText()
    })
    EventsOn("call-post", () => {
        post()
    })
    EventsOn("call-ondomready", (config) => {
        LogInfo(JSON.stringify(config))
        setBackgroundColor(config.Window.BackgroundR, config.Window.BackgroundG, config.Window.BackgroundB, config.Window.BackgroundA)
    })
</script>

<main style="--wails-draggable:drag">
    <textarea autocomplete="off" bind:value={text} on:input={onChange} placeholder={placeholder} id="inputbox" />
    <p id="char-counter" class="char-count">{charCounter}</p>
</main>
