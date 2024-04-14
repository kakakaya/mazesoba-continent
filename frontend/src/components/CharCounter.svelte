<script lang="ts">
    export let count = 0;
    export const MAX_COUNT = 300; // this might be changed later
    export let hidden: boolean = false;

    function update(event: any) {
        count = event.target.value.length;
    }
</script>

<div
    class="char-counter"
    class:char-counter--warning={count > MAX_COUNT}
    {hidden}
    on:change={update}
>
    {count}
</div>

<style>
    .char-counter {
        color: #333;
    }
    .char-counter--warning {
        background-color: #ffcccc;
        color: #990000;
    }
</style>
