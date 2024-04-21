<script lang="ts">
    import {
        OpenConfigDirectory,
        OpenLogDirectory,
        GetContext,
        Chikuwa,
    } from "../wailsjs/go/main/App";

    // import {
    //     EventsOn,
    //     EventsEmit,
    //     WindowCenter,
    //     WindowHide,
    //     WindowShow,
    //     Quit,
    //     LogInfo,
    //     LogWarning,
    // } from "../wailsjs/runtime/runtime";
    import { dispatchInput } from "./dispatchInput.js";

    import { ConvertRichUnicode } from "./topping/unicode";

    let input: string = "";
    let helpMessage = "Ready";
    let postFooter = "";
    let charCounter = 0;
    let placeholder = Math.random() > 0.5 ? "最近どう？" : "どう最近？";

    function clearText() {
        input = "";
        charCounter = 0;
        const inputBox = document.getElementById("inputbox");
        if (inputBox) {
            inputBox.removeAttribute("readonly");
            inputBox.focus();
        }
    }

    function onInputChange() {
        dispatchInput(input, true)
            .then((result) => {
                if (/^\d+$/.test(result)) {
                    charCounter = parseInt(result);
                    helpMessage = "";
                } else {
                    charCounter = 0;
                    helpMessage = result;
                }
            })
            .catch((error) => {
                charCounter = 0;
                helpMessage = error;
            });
    }

    function executeInput() {
        dispatchInput(input, false)
            .then((result) => {
                clearText();
                placeholder = result;
            })
            .catch((error) => {
                clearText();
                helpMessage = `Error: ${error}`;
            });
    }
    function handleKeyDown(event: KeyboardEvent) {
        // Press Ctrl+Enter to send message
        if ((event.ctrlKey || event.metaKey) && event.key === "Enter") {
            event.preventDefault();
            executeInput();
        }

        // Press Ctrl+, to show settings
        if ((event.ctrlKey || event.metaKey) && event.key === ",") {
            event.preventDefault();
            OpenConfigDirectory();
        }

        // Press Ctrl+. to show settings
        if ((event.ctrlKey || event.metaKey) && event.key === ".") {
            event.preventDefault();
            OpenLogDirectory();
        }

        // Press Ctrl+F2 to post "earthquake"
        if ((event.ctrlKey || event.metaKey) && event.key === "F2") {
            event.preventDefault();
            Chikuwa("地震だ！")
                .then((result) => {
                    placeholder = result;
                })
                .catch((error) => {
                    helpMessage = `Error: ${error}`;
                });
        }
    }
    // // Setup Events
    // EventsOn("call-clearText", () => {
    //     clearText();
    // });
</script>

<main style="--wails-draggable:drag">
    <textarea
        autocomplete="off"
        bind:value={input}
        on:input={onInputChange}
        on:keydown={handleKeyDown}
        {placeholder}
        id="inputbox"
    />
    <div id="status-bar">
        <div id="footer">{postFooter}</div>
        <div id="message">{helpMessage}</div>
        <div id="char-counter">{charCounter}</div>
    </div>
</main>

<style>

body {
    /*
    margin: 0%;
    padding: 0%;
    */
    font-family: "Nunito", -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
        "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
        sans-serif;
    display: flex;
    display: grid;
    background: transparent;
    background-color: rgba(27, 38, 54, 0.5);
    text-align: center;
}

@font-face {
    font-family: "Nunito";
    font-style: normal;
    font-weight: 400;
    src: local(""),
        url("assets/fonts/nunito-v16-latin-regular.woff2") format("woff2");
}

#inputbox {
    background: transparent;
    color: white;
    height: 95vh;
    width: 100vw;
    resize: none;
    border: none;
    outline: none;
}

#status-bar {
    /* position: absolute; */
    position: fixed;
    bottom: 0;
    width: 100%;
    /* background-color: yellow; */
    display: flex;
    color: white;
    padding: 0px;
    justify-content: flex-end;
    border-top: 1px solid rgba(255, 255, 255, 0.1);

    text-align: left;
}

#status-bar>div {
    margin-left: 10px;
    padding: 5px;
}

#char-counter {
    color: white;
    /* background-color: yellowgreen; */
    padding: 5px;
    border-radius: 5px;
}

#message {
    /* background-color: royalblue; */
}

#footer {
    /* background-color: hotpink; */
}

.char-count {
    position: absolute;
    bottom: 0;
    right: 0;
    margin: 0;
    color: black;
    background-color: white;
    padding: 2px 5px;
}
</style>