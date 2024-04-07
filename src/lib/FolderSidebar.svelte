<script lang="ts">
    import FileSidebar from "./FileSidebar.svelte";
    import { ChevronDown, ChevronRight } from "lucide-svelte";
    import type { BackendFolder } from "./types";
    import ActionIconsElements from "./ActionIconsElements.svelte";
    import { serverURL, showMoveDialog, updateStructure } from "./store";

    export let folder : BackendFolder;
    export let indent = 0;

    let open = false;

    function toggleOpen() {
        open = !open;
    }
    let newIdent = indent + 24;

    let hover = false;

    const deleteObjekt = async () => {
        try {
            const response = await fetch($serverURL + "folder/" + folder.ID, {
                method: "DELETE",
            });
            console.log(response);
            updateStructure.set();
        } catch (error) {
            console.error(error);
        }
    };

    const moveObjekt = async () =>{
        showMoveDialog.set()
    }
</script>

<button
class="justify-between h-full w-full relative hover:bg-gray-200 hover:rounded-lg "
    on:mouseenter={()=>(hover = true)}
    on:mouseleave={()=>(hover = false)}>
    <button
        style="padding-left: {indent}px"
        class="flex items-center w-full h-full"
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
        <!-- <Folder class="size-4 mx-1" /> -->
        <div class="">
            {folder.Name}
        </div>
    </button>
    {#if hover}
        <div class="z-10 absolute right-0 top-0 h-full">
            <ActionIconsElements
                on:delelte={deleteObjekt}
                on:move={moveObjekt}
            />
        </div>
    {/if}
</button>

{#if open}
    {#if folder.Children.length > 0}
        {#each folder.Children as child}
            <div>
                <svelte:self folder={child} indent={newIdent} />
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
