<script lang="js">
    import { objectID, serverURL } from "$lib/store.js";
    import { onMount } from "svelte";
    import MarkdownEditor from "../lib/MarkdownEditor.svelte";
    import RootFolderSidebar from "../lib/RootFolderSidebar.svelte";
    import { PanelLeftClose, PanelLeftOpen } from "lucide-svelte";

    let sidebarVisible = true;

    onMount(()=>{
        serverURL.set("http://localhost:8080/")
    });

    function toggleSidebar() {
        sidebarVisible = !sidebarVisible;
    }
    async function addFile() {
        console.log("addFile");
        const url = $serverURL + "jsonobj";
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
        const url = "http://test.johafo.de:8080/folder";
        const data = { Name: "Hallo Welt", ParentID: 2 };
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
                <RootFolderSidebar />
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
