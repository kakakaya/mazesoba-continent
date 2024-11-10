<script lang="ts">
    export let count = 0;
    export let maxCount = 300; // default value is for Bluesky
</script>

<div
    class="char-counter"
    class:char-counter--warning={count > maxCount}
    class:no-count={count < 0}
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
    .no-count {
        display: none;
    }
</style>
