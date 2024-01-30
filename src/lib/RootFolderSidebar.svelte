<script>
    import { onMount } from "svelte";
    import FileSidebar from "./FileSidebar.svelte";
    import FolderSidebar from "./FolderSidebar.svelte";

    let folders = [];

    onMount(async () => {
        const response = await fetch("http://test.johafo.de:8080/structure");
        var tmp = await response.json();
        folders = formatFolders(tmp);
        console.log(folders);
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

        // Return the root folder(s)
        return folders
            .filter((folder) => !folder.ParentID)
            .map((folder) => folderMap[folder.ID]);
    }
</script>

<div class="w-max h-max">
    {#if folders.length > 0}
        {#each folders[0].Children as folder}
            <div>
                <FolderSidebar json={folder} indent={4} />
            </div>
        {/each}
        {#each folders[0].Objects as file}
            <div>
                <FileSidebar data={file} indent={4} />
            </div>
        {/each}
    {/if}
</div>
