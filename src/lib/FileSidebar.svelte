<script lang="ts">
    import { currentObject } from "./store.js";
    import type { BackendObject } from "./types";
    import { File } from "lucide-svelte";
    import { FileText } from "lucide-svelte";
    import { Image } from "lucide-svelte";

    export let data: BackendObject;

    export let indent: number = 0;

    function setInhalt() {
        currentObject.set(data);
    }

    let isSelected = false;

    currentObject.subscribe((value) => {
        if (value === null) {
            isSelected = false;
            return;
        }
        if (value.ID == data.ID) {
            isSelected = true;
        } else {
            isSelected = false;
        }
    });
</script>

{#if !isSelected}
    <button
        style="padding-left: {indent}px"
        class="flex hover:bg-gray-200 hover:rounded-lg w-full h-full truncate items-center overflow-hidden"
        on:dblclick={setInhalt}
    >
        {#if data.Type == "MD"}
            <File class="size-4 mx-1 flex-shrink-0" />
        {:else if data.Type == "application/pdf"}
            <FileText class="size-4 mx-1 flex-shrink-0" />
        {:else}
            <Image class="size-4 mx-1 flex-shrink-0" />
        {/if}
        {data.Name}
    </button>
{:else}
    <button
        style="padding-left: {indent}px"
        class="flex bg-gray-200 rounded-lg w-full h-full truncate items-center overflow-hidden"
        on:dblclick={setInhalt}
    >
        {#if data.Type == "MD"}
            <File class="size-4 mx-1 flex-shrink-0" />
        {:else if data.Type == "application/pdf"}
            <FileText class="size-4 mx-1 flex-shrink-0" />
        {:else}
            <Image class="size-4 mx-1 flex-shrink-0 " />
        {/if}
        {data.Name}
    </button>
{/if}
