<script>
    import logo from './assets/images/logo-universal.png'
    import {
        Chikuwa,
        Post,
    } from '../wailsjs/go/main/App.js'
    import {
        EventsOn,
        EventsEmit,
        WindowCenter,
        WindowHide,
        WindowShow,
        Quit
    } from '../wailsjs/runtime/runtime.js'

    import {
        ConvertRichUnicode
    } from './topping/unicode.js'

    let text
    let charCounter = 0
    let placeholder = Math.random() > 0.5 ? "最近どう？" : "どう最近？"

    function post() {
        Post(text).then(result => {
            placeholder = result
            text = ""
        })
    }

    function chikuwa() {
        const c = text ? text : "ちくわ。"
        Chikuwa(c).then(result => {
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


    // Setup Events
    EventsOn("call-clearText", () => {
        clearText()
    })
    EventsOn("call-post", () => {
        post()
    })
    EventsOn("call-chikuwa", () => {
        chikuwa()
    })
</script>

<main style="--wails-draggable:drag">
    <div class="input-box" id="input">
        <textarea autocomplete="off" bind:value={text} on:input={onChange} placeholder={placeholder} id="inputbox" />
        <p id="char-counter" class="char-count">{charCounter}</p>
    </div>
    <p id="input-length"></p>
</main>
