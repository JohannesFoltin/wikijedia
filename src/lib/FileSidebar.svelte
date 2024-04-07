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
    let style: string;

    // Das Objekt wird ausgewählt
    function setInhalt() {
        currentObject.set(data);
    }

    let isSelected = false;
    // Es wird geschaut ob das derzeitige Objekt ausgewählt ist
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
    // Das Objekt wird gelöscht
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

    // Der Dialog zum verschieben des Objekts wird geöffnet
    function moveObjekt() {
        showMoveDialog.set(-2);
    }

    // Es wird geschaut ob das Objekt ausgewählt ist und dementsprechend wird das Styling angepasst
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
