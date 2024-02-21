<script lang="ts">
    import FileSidebar from "./FileSidebar.svelte";
    import { ChevronDown, ChevronRight } from "lucide-svelte";
    import type { BackendFolder } from "./types";

    export let folder : BackendFolder;
    export let indent = 0;

    let open = false;

    function toggleOpen() {
        open = !open;
    }
    let newIdent = indent + 24;
</script>

<button
    style="padding-left: {indent}px"
    class="flex items-center hover:bg-gray-200 hover:rounded-lg w-full h-full"
    on:click={toggleOpen}
>
    {#if !open}
        <div class="flex items-center">
            <ChevronRight class="size-4 mx-auto" />
        </div>
    {:else}
        <div class="flex items-center">
            <ChevronDown class="size-4 mx-auto" />
        </div>
    {/if}
    <div class="">
        {folder.Name}
    </div>
</button>

{#if open}
    {#if folder.Children.length > 0}
        {#each folder.Children as child}
            <div>
                <svelte:self json={child} indent={newIdent} />
            </div>
        {/each}
    {/if}
    {#if folder.Objects.length > 0}
        {#each folder.Objects as file}
            <div>
                <FileSidebar data={file} indent={newIdent} />
            </div>
        {/each}
    {/if}
{/if}
