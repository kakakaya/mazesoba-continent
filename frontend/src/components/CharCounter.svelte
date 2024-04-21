<script lang="ts">
    export let count = 0;
    export let maxCount = 300; // default value is for Bluesky
    export let hidden: boolean = false;
</script>

<div
    class="char-counter"
    class:char-counter--warning={count > maxCount}
    hidden={count < 0}
>
    {count}
</div>

<style>
    .char-counter {
        width: min-content;
        padding: 0.2em;
    }
    .char-counter--warning {
        background-color: #ffcccc;
        color: #990000;
    }
    .char-counter[hidden] {
        display: none;
    }
</style>
