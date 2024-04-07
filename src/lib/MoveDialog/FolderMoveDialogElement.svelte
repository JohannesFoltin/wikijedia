<script lang="ts">
    import { FolderOpen } from 'lucide-svelte';
    import type { BackendFolder } from "../types";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let folder: BackendFolder;
    export let indent = 0;

    let newIdent = indent + 24;
</script>

<button
    class="justify-between h-full w-full relative hover:bg-gray-200 hover:rounded-lg"
    on:dblclick={() => dispatch("selected", folder.ID)}
>
    <button
        style="padding-left: {indent}px"
        class="flex items-center w-full h-full"
    >
        <div class="flex items-center">
            <FolderOpen class="size-4 mx-auto" />
        </div>
        <!-- <Folder class="size-4 mx-1" /> -->
        <div class="">
            {folder.Name}
        </div>
    </button>
</button>

{#if folder.Children.length > 0}
    {#each folder.Children as child}
        <div>
            <svelte:self
                folder={child}
                indent={newIdent}
                on:selected={({ detail }) => dispatch("selected", detail)}
            />
        </div>
    {/each}
{/if}
