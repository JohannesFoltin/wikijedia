<script>
    import { objectID } from "./store";

    export let data = {};
    export let indent = 0;

    function setInhalt() {
        objectID.set(data.ID);
    }

    let isSelected = false;

    objectID.subscribe((value) => {
        if ((value == data.ID)) {
            isSelected = true;
        } else {
            isSelected = false;
        }
    });
</script>

{#if !isSelected}
    <button
        style="padding-left: {indent}px"
        class="flex hover:bg-gray-200 hover:rounded-lg w-full h-full"
        on:click={setInhalt}
    >
        {data.Name}
    </button>
{:else}
    <button
        style="padding-left: {indent}px"
        class="flex bg-gray-200 rounded-lg w-full h-full"
        on:click={setInhalt}
    >
        {data.Name}
    </button>
{/if}
