<script lang="ts">
    import {
        currentObject,
        serverURL,
        showMoveDialog,
        updateStructure,
    } from "../lib/store";
    import FolderSidebar from "../lib/FolderSidebar.svelte";
    import FileSidebar from "../lib/FileSidebar.svelte";
    import { onMount } from "svelte";
    import MarkdownEditor from "../lib/ObjectViews/ObjectView.svelte";
    import { PanelLeftClose, PanelLeftOpen } from "lucide-svelte";
    import AddButton from "../lib/AddButton.svelte";
    import type { BackendFolder } from "../lib/types";
    import MoveDialog from "../lib/MoveDialog/MoveDialog.svelte";

    let unlocked: boolean = true;

    let pinInputType: "text" | "password" = "password";
    $: pinInputType = unlocked ? "text" : "password";

    let sidebarVisible: boolean = true;

    let rootFolder: BackendFolder = null;

    // Die initiale URL des Servers wird gesetzt und die Ordnerstruktur wird geladen
    onMount(async () => {
        serverURL.set("http://localhost:8080/");
        rootFolder = await getFolderstruture();
    });

    // Wenn die Ordnerstruktur aktualisiert werden soll, wird die Funktion getFolderstruture aufgerufen
    updateStructure.subscribe(async (needResync) => {
        if (needResync) {
            rootFolder = await getFolderstruture();
            updateStructure.reset();
        }
    });

    function toggleSidebar() {
        sidebarVisible = !sidebarVisible;
    }

    // Die Ordnerstruktur wird vom Server geladen und in ein Baumstruktur umgewandelt
    async function getFolderstruture() {
        const response = await fetch($serverURL + "structure");
        const tmp: string = await response.json();
        const test: BackendFolder[] = JSON.parse(JSON.stringify(tmp));

        let folderMap = new Map<number, BackendFolder>();

        // First pass: create a map of all folders by ID
        for (const folder of test) {
            folderMap[folder.ID] = { ...folder, Children: [] };
        }

        // Second pass: add each folder to its parent's Children array
        for (const folder of test) {
            if (folder.ParentID in folderMap) {
                folderMap[folder.ParentID].Children.push(folderMap[folder.ID]);
            }
        }
        const rootFolder = folderMap[1];
        return rootFolder;
    }
</script>

<div class="flex h-screen w-screen">
    <div class="pr-4">
        <div class="h-10 w-full flex flex-row">
            <div class="flex w-min items-center pl-2">
                <button class="h-8" on:click={toggleSidebar}>
                    {#if sidebarVisible}
                        <PanelLeftClose class="h-full" />
                    {:else}
                        <PanelLeftOpen />
                    {/if}
                </button>
            </div>
            {#if sidebarVisible}
                <div class="flex flex-grow items-center justify-center">
                    <AddButton
                        on:update={async () => {
                            rootFolder = await getFolderstruture();
                        }}
                    />
                </div>
            {/if}
        </div>
        {#if sidebarVisible}
            <div class="flex-none w-56">
                <div class="w-full h-max">
                    {#if rootFolder !== undefined && rootFolder !== null}
                        {#each rootFolder.Children as folder}
                            <div>
                                <FolderSidebar {folder} indent={4} />
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
    {#if $currentObject !== null}
        <div class="flex-1">
            <MarkdownEditor
                on:updateName={async () => {
                    rootFolder = await getFolderstruture();
                }}
            />
        </div>
    {:else}
        <div class="h-full w-full flex items-center justify-center">
            <h1>
                Hallo in deinem Persönlichen Wiki. Wähle doch mal eine Datei aus
            </h1>
        </div>
    {/if}
</div>
<MoveDialog
    dialogOpen={$showMoveDialog != -1}
    {rootFolder}
    on:save={async ({ detail }) => {
        // Wenn die Datei verschoben werden soll, wird die Datei verschoben. Da sie unter dem aktuellen Objekt verschoben werden soll, wird -2 übergeben
        if ($showMoveDialog == -2) {
            let data = {parentFolderID: detail};
            try {
                const response = await fetch(
                    $serverURL + "object/" + $currentObject.ID + "/parent",
                    {
                        method: "PUT",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify(data),
                    },
                );
                console.log(response);
                updateStructure.set();
            } catch (error) {
                console.error(error);
            }
        } 
        // Sonst wird der Ordner verschoben udn die OrdnerID ist der showMoveDialog Wert
        else {
            let data = {parentFolderID: detail};
            try {
                const response = await fetch(
                    $serverURL + "folder/" + $showMoveDialog + "/parent",
                    {
                        method: "PUT",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify(data),
                    },
                );
                console.log(response);
                updateStructure.set();
            } catch (error) {
                console.error(error);
            }
        }
        showMoveDialog.reset();
    }}
    on:error={() => {
        showMoveDialog.reset();
    }}
></MoveDialog>
