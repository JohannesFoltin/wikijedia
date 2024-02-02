<script lang="js">
    import { objectID, serverURL } from "$lib/store.js";
    import FolderSidebar from "../lib/FolderSidebar.svelte";
    import FileSidebar from "../lib/FileSidebar.svelte";
    import { onMount } from "svelte";
    import MarkdownEditor from "../lib/MarkdownEditor.svelte";
    import { PanelLeftClose, PanelLeftOpen } from "lucide-svelte";

    let sidebarVisible = true;
    let rootFolder = null;

    onMount(async () => {
        serverURL.set("http://localhost:8080/");
        rootFolder = await getFolderstruture();
    });

    function toggleSidebar() {
        sidebarVisible = !sidebarVisible;
    }

    async function getFolderstruture() {
        const response = await fetch($serverURL + "structure");
        var tmp = await response.json();

        const folderMap = {};

        // First pass: create a map of all folders by ID
        for (const folder of tmp) {
            folderMap[folder.ID] = { ...folder, Children: [] };
        }

        // Second pass: add each folder to its parent's Children array
        for (const folder of tmp) {
            if (folder.ParentID in folderMap) {
                folderMap[folder.ParentID].Children.push(folderMap[folder.ID]);
            }
        }
        const rootFolder = folderMap[1];

        return rootFolder;
    }

    async function addFile() {
        console.log("addFile");
        const url = $serverURL + "object";
        const data = { Name: "Hallo Welt", Data: "# Hello World", FolderID: 3 };
        try {
            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });

            if (response.ok) {
                console.log("Post request successful");
            } else {
                console.error("Post request failed");
            }
        } catch (error) {
            console.error("An error occurred:", error);
        }
    }
    async function addFolder() {
        console.log("addFile");
        const url = $serverURL + "folder";
        const data = { Name: "Hallo Welt", ParentID: 1 };
        try {
            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });

            if (response.ok) {
                console.log("Post request successful");
            } else {
                console.error("Post request failed");
            }
        } catch (error) {
            console.error("An error occurred:", error);
        }
    }
</script>

<div class="flex h-screen w-screen">
    <div class="pr-4">
        {#if sidebarVisible}
            <div class="flex items-center justify-around">
                <button class="" on:click={toggleSidebar}>
                    <PanelLeftClose />
                </button>
                <button on:click={addFile}> AddF </button>
                <button on:click={addFolder}> AddFo </button>
            </div>
        {:else}
            <div class="flex items-center">
                <button class="" on:click={toggleSidebar}>
                    <PanelLeftOpen />
                </button>
            </div>
        {/if}
        {#if sidebarVisible}
            <div class="flex-none pr-10 w-56">
                <div class="w-max h-max">
                    <button on:click={getFolderstruture}> Reload </button>
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
            </div>
        {/if}
    </div>
    {#if $objectID !== ""}
        <MarkdownEditor class="flex-1" />
    {:else}
        <div class="h-full w-full flex items-center justify-center">
            <h1>
                Hallo in deinem Persönlichen Wiki. Wähle doch mal eine Datei aus
            </h1>
        </div>
    {/if}
</div>
