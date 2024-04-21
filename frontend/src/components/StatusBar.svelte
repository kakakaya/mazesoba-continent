<script lang="ts">
    import CharCounter from "./CharCounter.svelte";

    export let helpMessage: string = "";
    export let postFooter: string = "";
    export let charCount: number = 0;
    export let maxCount: number;
</script>

<div class="status-bar">
    <div class="message status-element">{helpMessage}</div>
    <div class="footer">{postFooter}</div>    
    <CharCounter count={charCount} maxCount={maxCount} hidden />
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
