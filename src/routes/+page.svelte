<script lang="js">
    import MarkdownEditor from "../lib/MarkdownEditor.svelte";
    import FolderSidebar from "../lib/FolderSidebar.svelte";
    import { PanelLeftClose, PanelLeftOpen, } from 'lucide-svelte';

    let sidebarVisible = true;

    function toggleSidebar() {
        sidebarVisible = !sidebarVisible;
    }
</script>

<div class="flex h-screen w-screen">
    {#if sidebarVisible}
        <div class="pr-10">
            <button class="" on:click={toggleSidebar}>    
                <PanelLeftClose/>
            </button>
            <FolderSidebar class="flex-1" />
        </div>
        {:else}
        <button class="" on:click={toggleSidebar}>    
            <PanelLeftOpen/>
        </button>    
        {/if}
    <MarkdownEditor class="flex-1" markdownText=" # Asd" />
</div>
