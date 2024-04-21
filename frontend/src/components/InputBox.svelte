<script lang="ts">
    export let value: string;
    export let placeholder: string = '';
</script>

<textarea
    class="inputbox"
    autocomplete="off"
    bind:value
    {placeholder}
    on:input
    on:keydown
/>

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
        padding: 0px;
        outline: none;
    }
</style>
