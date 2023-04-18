<script>
    import logo from './assets/images/logo-universal.png'
    import {
        Chikuwa,
        Post,
        OpenConfigDirectory,
        OpenLogDirectory,
        GetVersion
    } from '../wailsjs/go/main/App.js'
    import {
        WindowCenter,
        WindowHide,
        WindowShow,
        Quit
    } from '../wailsjs/runtime/runtime.js'

    import {
        GenerateDummyInviteCode,
    } from './topping/invite-code.js'

    import {
        ConvertRichUnicode
    } from './topping/unicode.js'

    let text
    let charCounter = 0
    let placeholder = Math.random() > 0.5 ? "最近どう？" : "どう最近？"


    function countGrapheme(string) {
        const segmenter = new Intl.Segmenter("ja", {
            granularity: "grapheme"
        });
        return [...segmenter.segment(string)].length;
    }

    function post() {
        Post(text).then(result => {
            placeholder = result
            text = ""
        })
    }


    function getDummyInviteCode() {
        if (text === undefined | !(typeof text === 'string' && text.startsWith('bsky-'))) {
            text = GenerateDummyInviteCode()
        } else {
            text += "\n" + GenerateDummyInviteCode()
        }

    }

    function earthquake() {
        text = "地震だ!"
    }

    function version() {
        GetVersion().then(version => text = "まぜそば大陸 バージョン" + version + "が浮上中")
    }

    function chikuwa() {
        if (text == "") {
            text = "ちくわ。"
        }
        Chikuwa(text).then(result => {
            placeholder = result
            text = ""
        })
    }

    function unicode() {
        text = ConvertRichUnicode(text, 'sansBold')
    }

    function handleKeyDown(event) {
        if ((event.ctrlKey || event.metaKey) && event.key == 'Enter') {
            event.preventDefault()
            post()
        } else if ((event.ctrlKey || event.metaKey) && event.key == ',') {
            event.preventDefault()
            OpenConfigDirectory()
        } else if (event.key == 'F9') {
            event.preventDefault();
            earthquake();
        } else if (event.key == 'Escape') {
            Quit();
        }
        charCounter = text.length;
    }
    window.addEventListener("keydown", handleKeyDown);
</script>

<main style="--wails-draggable:drag">
    <div class="input-box" id="input">
        <textarea autocomplete="off" bind:value={text} placeholder={placeholder} id="inputbox" />
        <p id="char-counter" class="char-count">{charCounter}</p>
    </div>
    <hr />
    <div class="buttons">
        <button class="btn" on:click={post}>Post(Ctrl+Enter)</button>
        <button class="btn" on:click={chikuwa}>ちくわ。</button>
        <button class="btn" on:click={WindowCenter}>画面中央</button>
        <button class="btn" on:click={earthquake}>地震だ！(F9)</button>
        <button class="btn" on:click={getDummyInviteCode}>嘘招待コード生成</button>
        <button class="btn" on:click={version}>バージョン</button>
        <button class="btn" on:click={unicode}>Unicode</button>
        <button class="btn" on:click={OpenConfigDirectory}>設定の場所を開く(Ctrl+,)</button>
        <button class="btn" on:click={OpenLogDirectory}>ログの場所を開く</button>
    </div>
    <p id="input-length"></p>
</main>
