<script>
    import FileSidebar from "./FileSidebar.svelte";
    import { ChevronDown, ChevronRight, ChevronUp } from "lucide-svelte";

    export let json = {};
    export let indent = 0;

    let open = true;

    function toggleOpen() {
        open = !open;
    }
    let newIdent = indent + 24;
</script>

<button
    style="padding-left: {indent}px"
    class="flex items-center"
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
        {json.Name}
    </div>
</button>

{#if open}
    {#if json.Children.length > 0}
        {#each json.Children as child}
            <div>
                <svelte:self json={child} indent={newIdent} />
            </div>
        {/each}
    {/if}
    {#if json.JSONObjChildren.length > 0}
        {#each json.JSONObjChildren as file}
            <div>
                <FileSidebar data={file} indent={newIdent} />
            </div>
        {/each}
    {/if}
{/if}
