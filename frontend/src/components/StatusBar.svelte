<script lang="ts">
    import CharCounter from "./CharCounter.svelte";
    import { footers } from "../stores";

    export let helpMessage: string = "";
    export let charCount: number = 0;
    export let maxCount: number;

    let footer_text = "";
    footers.subscribe((value) => {
        footer_text = value.join(" ");
    });
</script>

<div class="status-bar">
    <div class="message status-element">{helpMessage}</div>
    <div class="footer">{footer_text}</div>    
    <CharCounter count={charCount} {maxCount} hidden />
</div>

<style>
    .status-bar {
        position: absolute;
        bottom: 0;
        right: 0;
        width: 100%;
        display: flex;
        color: white;
        margin: 0px;
        justify-content: flex-end;
        border-top: 1px solid rgba(255, 255, 255, 0.1);
        flex-wrap: wrap;
        justify-content: flex-end;
        gap: 0em 0.2em;
    }

    .status-bar > div {
        padding: 0.2em;
    }
</style>
