<script lang="ts">
	import MoveDialog from './../lib/MoveDialog.svelte';
    import { currentObject, serverURL, showMoveDialog, updateStructure } from "../lib/store";
    import FolderSidebar from "../lib/FolderSidebar.svelte";
    import FileSidebar from "../lib/FileSidebar.svelte";
    import { onMount } from "svelte";
    import MarkdownEditor from "../lib/ObjectViews/ObjectView.svelte";
    import { PanelLeftClose, PanelLeftOpen } from "lucide-svelte";
    import { PinInput } from "bits-ui";
    import { Toggle } from "bits-ui";
    import { LockKeyhole, UnlockKeyhole } from "lucide-svelte";
    import AddButton from "../lib/AddButton.svelte";
    import type { BackendFolder } from "../lib/types";

    let value: string[] | undefined = ["5", "1", "3", "7"];

    let unlocked: boolean = true;

    let pinInputType: "text" | "password" = "password";
    $: pinInputType = unlocked ? "text" : "password";

    let sidebarVisible: boolean = true;

    let rootFolder: BackendFolder = null;

    onMount(async () => {
        serverURL.set("http://localhost:8080/");
        rootFolder = await getFolderstruture();
        console.log(rootFolder);
    });

    updateStructure.subscribe(async (needResync) => {
        if (needResync) {
            rootFolder = await getFolderstruture();
            updateStructure.reset();
        }
    });

    function toggleSidebar() {
        sidebarVisible = !sidebarVisible;
    }

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
        console.log(rootFolder);
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
    dialogOpen={$showMoveDialog}
    rootFolder={rootFolder}
    on:save={(test) => {
        showMoveDialog.reset();
    }}
    on:error={() => {
        showMoveDialog.reset();
    }}
></MoveDialog>
