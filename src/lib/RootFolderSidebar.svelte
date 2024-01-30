<script>
    import { onMount } from "svelte";
    import FileSidebar from "./FileSidebar.svelte";
    import FolderSidebar from "./FolderSidebar.svelte";

    let folders = [];

    onMount(async () => {
        const response = await fetch("http://test.johafo.de:8080/folders");
        folders = await response.json();
        console.log(folders);
    });
</script>

<div class="w-max h-max">
    {#if folders.length > 0}
        {#each folders[0].Children as folder}
            <div>
                <FolderSidebar json={folder} indent={4} />
            </div>
        {/each}
        {#each folders[0].JSONObjChildren as file}
            <div>
                <FileSidebar data={file} indent={4} />
            </div>
        {/each}
    {/if}
</div>
