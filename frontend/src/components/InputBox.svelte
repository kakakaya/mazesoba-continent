<script lang="ts">
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

    function handleKeyDown(event: number) {
        if (event.key === "Enter") {
            executeInput();
        }
    }
</script>

<textarea
    class="inputbox"
    autocomplete="off"
    bind:value={input}
    {placeholder}
/>

<!-- on:input={onInputChange}
    on:keydown={handleKeyDown} -->
<style>
    .inputbox {
        background: transparent;
        color: white;
        /* height: 95vh; */
        display: flex;
        height: 100%;
        width: 100%;
        resize: none;
        border: none;
        outline: none;
    }
</style>
