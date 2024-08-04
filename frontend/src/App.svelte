<script lang="ts">
    import {
        OpenConfigDirectory,
        OpenLogDirectory,
        GetContext,
        Chikuwa,
        UploadImage,
        Post,
    } from "../wailsjs/go/main/App";
    import { Environment } from "../wailsjs/runtime";
    import InputBox from "./components/InputBox.svelte";
    import StatusBar from "./components/StatusBar.svelte";

    import {
        EventsOn,
        EventsEmit,
        WindowCenter,
        WindowHide,
        WindowShow,
        Quit,
        LogInfo,
        LogWarning,
        OnFileDrop,                 
    } from "../wailsjs/runtime/runtime";
    import { dispatchInput } from "./dispatchInput.js";

    import { ConvertRichUnicode } from "./topping/unicode";

    struct : {}
    let input: string = "";
    let images
    let helpMessage = "Ready";
    let charCount = 0;  // -1 to hide CharCounter           
    let placeholder = Math.random() > 0.5 ? "最近どう？" : "どう最近？";


    GetContext()
        .then((context) => {
            const ctx = JSON.parse(context);
            const Version = ctx.version;
            helpMessage = `Ready: ${ctx.id}@${ctx.host}`;
        })
        .catch((err) => {
            return Promise.reject(err);
        });

    OnFileDrop((_x, _y, paths) => {
        LogInfo(`Dropped: ${paths}, ${_x}, ${_y}`);
        placeholder = '[' + paths.join(", ") + ']';
        input = paths.join(", ");
        let foo = '';
        UploadImage(paths[0])
            .then((result) => {
                // result is base64 encoded, so decode it.
                LogInfo(result)
                foo = result;
                result = atob(result);
                LogInfo(result)
                input = result;
            })
            .catch((error) => {
                helpMessage = `Error: ${error}`;
            });
        PostWithImage(input, foo)
            .then((result) => {
                placeholder = result;
            })
            .catch((error) => {
                helpMessage = `Error: ${error}`;
            });
        LogInfo('ende')
    }, true);

    function clearText() {
        input = "";
        charCount = -1;
        const inputBox = document.getElementById("inputbox");
        if (inputBox) {
            inputBox.removeAttribute("readonly");
            inputBox.focus();
        }
    }

    function handleInput() {
        dispatchInput(input, true)
            .then((result) => {
                if (/^\d+$/.test(result)) {
                    // If input is usual message,
                    charCount = parseInt(result);
                    helpMessage = "";
                } else {
                    // Otherwise hide CharCounter and show help
                    charCount = -1;
                    helpMessage = result;
                }
            })
            .catch((error) => {
                charCount = -1;
                helpMessage = error;
            });
    }

    function executeInput() {
        // Store current input value and clear input
        const currentInput = input;
        clearText();
        placeholder = "送信中...";
        dispatchInput(currentInput, false)
            .then((result) => {
                placeholder = result;
            })
            .catch((error) => {
                helpMessage = `Error: ${error}`;
            });
    }
    function handleKeyDown(keyEvent: KeyboardEvent) {
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

<body>
    <InputBox
        bind:value={input}
        on:keydown={handleKeyDown}
        on:input={handleInput}
        {placeholder}
    />
    <StatusBar {helpMessage} {charCount} maxCount={300} />
</body>

<style>
    body {
        --wails-draggable: drag;
        --wails-drop-target: drop;
        background-color: rgba(27, 38, 54, 0.5);
        height: 100vh;
        width: 100vw;
    }
</style>
