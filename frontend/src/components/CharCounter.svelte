<script lang="ts">
    export let count = 0;
    export let max_count = 300;  // default value is for Bluesky
    export let hidden: boolean = false;
</script>

<div
    class="char-counter"
    class:char-counter--warning={count > max_count}
    {hidden}
>
    {count}
</div>

<style>
    .char-counter {
        color: #333;
        width: min-content;
    }
    .char-counter--warning {
        background-color: #ffcccc;
        color: #990000;
    }
    .char-counter[hidden] {
        display: none;
    }
</style>
