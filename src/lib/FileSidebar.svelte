<script lang="ts">
    import { currentObject } from "./store.js";
    import {BackendObject} from "$lib/structure";

    export let data : BackendObject;
    export let indent : number = 0;

    function setInhalt() {
        currentObject.set(data);
    }

    let isSelected = false;

    currentObject.subscribe((value) => {
        if(value === null){
            isSelected = false;
            return;
        }
        if ((value.ID == data.ID)) {
            isSelected = true;
        } else {
            isSelected = false;
        }
    });
</script>

{#if !isSelected}
    <button
        style="padding-left: {indent}px"
        class="flex hover:bg-gray-200 hover:rounded-lg w-full h-full truncate"
        on:dblclick={setInhalt}
    >
        {data.Name}
    </button>
{:else}
    <button
        style="padding-left: {indent}px"
        class="flex bg-gray-200 rounded-lg w-full h-full truncate"
        on:dblclick={setInhalt}
    >
        {data.Name}
    </button>
{/if}
