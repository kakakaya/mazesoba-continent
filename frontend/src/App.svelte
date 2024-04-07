<script lang="ts">
    import {
        OpenConfigDirectory,
        OpenLogDirectory,
        GetContext,
        Chikuwa,
    } from "../wailsjs/go/main/App";

    import {
        EventsOn,
        EventsEmit,
        WindowCenter,
        WindowHide,
        WindowShow,
        Quit,
        LogInfo,
        LogWarning,
    } from "../wailsjs/runtime/runtime.js";
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
    // Setup Events
    EventsOn("call-clearText", () => {
        clearText();
    });
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
