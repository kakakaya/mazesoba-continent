<script>
    import logo from './assets/images/logo-universal.png'
    import {
        Chikuwa,
        Post
    } from '../wailsjs/go/main/App.js'
    import {
        WindowCenter,
        WindowHide,
        WindowShow
    } from '../wailsjs/runtime/runtime.js'

    let text
    let placeholder

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

 function earthquake() {
     text = "地震だ!"
     // post();
    }

    function chikuwa() {
        Chikuwa().then(result => {
            placeholder = result
            text = ""
        })
    }

    function windowHide() {
        WindowHide()
    }

    function windowShow() {
        WindowShow()
    }

    function handleKeyDown(event) {
        if ((event.ctrlKey || event.metaKey) && event.key == 'a') {
            event.preventDefault();
            text = "全部選択"
        } else if ((event.ctrlKey || event.metaKey) && event.key == 'Enter') {
            event.preventDefault();
            post();
        } else if (event.key == 'F2') {
            event.preventDefault();
            // open config
        } else if (event.key == 'F9') {
            event.preventDefault();
            earthquake();
        }
    }
    window.addEventListener("keydown", handleKeyDown);
</script>

<main style="--wails-draggable:drag">
    <div class="input-box" id="input">
        <textarea autocomplete="off" bind:value={text} placeholder={placeholder} id="name" />
    </div>
    <hr/>
    <div class="buttons">
        <p>
            ボタン群(仮)
        </p>
        <button class="btn" on:click={post}>Post(Ctrl+Enter)</button>
        <button class="btn" on:click={chikuwa}>ちくわ。</button>
        <button class="btn" on:click={WindowCenter}>画面中央</button>
        <button class="btn" on:click={windowHide}>画面hide</button>
        <button class="btn" on:click={windowShow}>画面show</button>
        <button class="btn" on:click={earthquake}>地震だ！(F9)</button>
    </div>
    <p id="input-length"></p>
</main>

<style>
 body {
     font-family: "Avenir-Roman", "Arial", "游ゴシック体", YuGothic, "游ゴシック Medium", "Yu Gothic Medium", "游ゴシック", "Yu Gothic", sans-serif, ui-sans-serif, system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, Noto Sans, sans-serif, Helvetica Neue, Arial, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol, Noto Color Emoji;
 }

    .result {
        height: 1em;
        line-height: 2em;
        margin: 1.5rem auto;
    }

    .input-box .input {
        border: none;
        border-radius: 3px;
        outline: none;
        height: 30px;
        line-height: 30px;
        padding: 0 10px;
        background-color: rgba(240, 240, 240, 0);
        -webkit-font-smoothing: antialiased;

    }

    .input-box .input:hover {
        border: none;
        background-color: rgba(255, 255, 255, 0.2);
    }

    .input-box .input:focus {
        border: none;
        background-color: rgba(255, 255, 255, 0.2);
    }

 textarea {
     color: white;
     font-family: "Avenir-Roman", "Arial", "游ゴシック体", YuGothic, "游ゴシック Medium", "Yu Gothic Medium", "游ゴシック", "Yu Gothic", sans-serif, ui-sans-serif, system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, Noto Sans, sans-serif, Helvetica Neue, Arial, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol, Noto Color Emoji;
     background: transparent;
     resize: none;
     border: 0 none;
     width: 100%;
 }
</style>
