<script lang="js">
    import MarkdownEditor from "../lib/MarkdownEditor.svelte";
    import RootFolderSidebar from "../lib/RootFolderSidebar.svelte";
    import { PanelLeftClose, PanelLeftOpen } from "lucide-svelte";
    
    let sidebarVisible = true;

    function toggleSidebar() {
        sidebarVisible = !sidebarVisible;
    }
</script>

<div class="flex h-screen w-screen">
    <div class="pr-4">
        <button class="" on:click={toggleSidebar}>
            {#if sidebarVisible}
                    <PanelLeftClose />
                {:else}
                <PanelLeftOpen />
            {/if}
        </button>
        {#if sidebarVisible}
        <div class="flex-none pr-10 w-56">
            <RootFolderSidebar/>
        </div>
        {/if}
    </div>
    <MarkdownEditor class="flex-1" markdownText=" # Asd" />
</div>
