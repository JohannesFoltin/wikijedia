<script lang="ts">
    import {
        currentObject,
        serverURL,
        showMoveDialog,
        updateStructure,
    } from "./store.js";
    import type { BackendObject } from "./types";
    import { File } from "lucide-svelte";
    import { FileText } from "lucide-svelte";
    import { Image } from "lucide-svelte";
    import ActionIconsElements from "./ActionIconsElements.svelte";

    export let data: BackendObject;

    export let indent: number = 0;

    let hover = false;

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

    const deleteObjekt = async () => {
        try {
            const response = await fetch($serverURL + "object/" + data.ID, {
                method: "DELETE",
            });
            console.log(response);
            updateStructure.set();
        } catch (error) {
            console.error(error);
        }
    };

    function moveObjekt() {
        showMoveDialog.set(-2);
    }

    let style: string;

    $: if (isSelected) {
        style = "bg-gray-200 rounded-lg ";
    } else {
        style = "hover:bg-gray-200 hover:rounded-lg";
    }
</script>

<button
    class="justify-between h-full w-full relative {style}"
    on:mouseenter={() => (hover = true)}
    on:mouseleave={() => (hover = false)}
>
    <button
        style="padding-left: {indent}px"
        class="flex w-full h-full truncate items-center overflow-hidden"
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
    {#if hover && isSelected}
        <div class="z-10 absolute right-0 top-0 h-full">
            <ActionIconsElements
                on:delelte={deleteObjekt}
                on:move={moveObjekt}
            />
        </div>
    {/if}
</button>
