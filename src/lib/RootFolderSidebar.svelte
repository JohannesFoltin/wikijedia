<script>
    import { onMount } from "svelte";
    import {serverURL} from "$lib/store";
    import FileSidebar from "./FileSidebar.svelte";
    import FolderSidebar from "./FolderSidebar.svelte";

    let rootFolder = null;

    onMount(async () => {
        const response = await fetch($serverURL +"structure");
        var tmp = await response.json();
        rootFolder = formatFolders(tmp);
        console.log(rootFolder);
    });

    function formatFolders(folders) {
        const folderMap = {};

        // First pass: create a map of all folders by ID
        for (const folder of folders) {
            folderMap[folder.ID] = { ...folder, Children: [] };
        }

        // Second pass: add each folder to its parent's Children array
        for (const folder of folders) {
            if (folder.ParentID in folderMap) {
                folderMap[folder.ParentID].Children.push(folderMap[folder.ID]);
            }
        }
        const rootFolder = folderMap[1];

        return rootFolder;
    }
</script>

<div class="w-max h-max">
    {#if rootFolder !== null}
        {#each rootFolder.Children as folder}
            <div>
                <FolderSidebar json={folder} indent={4} />
            </div>
        {/each}
        {#each rootFolder.Objects as file}
            <div>
                <FileSidebar data={file} indent={4} />
            </div>
        {/each}
    {/if}
</div>
