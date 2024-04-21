<script lang="ts">
    import {
        OpenConfigDirectory,
        OpenLogDirectory,
        GetContext,
        Chikuwa,
    } from "../wailsjs/go/main/App";
    import InputBox from "./components/InputBox.svelte";
    import StatusBar from "./components/StatusBar.svelte";

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
    let charCount = 0;
    let placeholder = Math.random() > 0.5 ? "最近どう？" : "どう最近？";

    function clearText() {
        input = "";
        charCount = 0;
        const inputBox = document.getElementById("inputbox");
        if (inputBox) {
            inputBox.removeAttribute("readonly");
            inputBox.focus();
        }
    }

    function handleInput() {
        charCount = Math.random() * 100;
        dispatchInput(input, true)
            .then((result) => {
                if (/^\d+$/.test(result)) {
                    charCount = parseInt(result);
                    helpMessage = "";
                } else {
                    charCount = 0;
                    helpMessage = result;
                }
            })
            .catch((error) => {
                charCount = 0;
                helpMessage = error;
            });
    }

    function executeInput() {
        clearText();
        placeholder = "送信中...";
        dispatchInput(input, false)
            .then((result) => {
                placeholder = result;
            })
            .catch((error) => {
                helpMessage = `Error: ${error}`;
            });
    }
    function handleKeyDown(event: CustomEvent<KeyboardEvent>) {
        const keyEvent: KeyboardEvent = event.detail;
        // Press Ctrl+Enter to send message
        if (
            (keyEvent.ctrlKey || keyEvent.metaKey) &&
            keyEvent.key === "Enter"
        ) {
            keyEvent.preventDefault();
            executeInput();
        }

        // Press Ctrl+, to show settings
        if ((keyEvent.ctrlKey || keyEvent.metaKey) && keyEvent.key === ",") {
            keyEvent.preventDefault();
            OpenConfigDirectory();
        }

        // Press Ctrl+. to show settings
        if ((keyEvent.ctrlKey || keyEvent.metaKey) && keyEvent.key === ".") {
            keyEvent.preventDefault();
            OpenLogDirectory();
        }

        // Press Ctrl+F2 to post "earthquake"
        if ((keyEvent.ctrlKey || keyEvent.metaKey) && keyEvent.key === "F2") {
            keyEvent.preventDefault();
            Chikuwa("地震だ！")
                .then((result) => {
                    placeholder = result;
                })
                .catch((error) => {
                    helpMessage = `Error: ${error}`;
                });
        }
    }
</script>

<main style="--wails-draggable:drag">
    <InputBox
        bind:value={input}
        on:input={handleInput}
        on:keydown={handleKeyDown}
        {placeholder}
    />
    <StatusBar {postFooter} {helpMessage} {charCount} maxCount={300} />
</main>
