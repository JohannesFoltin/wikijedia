<script lang="ts">
    import { currentObject, serverURL } from "$lib/store";
    import FolderSidebar from "../lib/FolderSidebar.svelte";
    import FileSidebar from "../lib/FileSidebar.svelte";
    import { onMount } from "svelte";
    import MarkdownEditor from "../lib/EditView.svelte";
    import { PanelLeftClose, PanelLeftOpen } from "lucide-svelte";
    import { PinInput } from "bits-ui";
    import { Toggle } from "bits-ui";
    import { LockKeyhole, UnlockKeyhole } from "lucide-svelte";
    import AddButton from "../lib/AddButton.svelte";

    let value: string[] | undefined = ["5", "1", "3", "7"];

    let unlocked = true;
    let pinInputType: "text" | "password" = "password";
    $: pinInputType = unlocked ? "text" : "password";
    let sidebarVisible = true;
    let rootFolder = null;

    onMount(async () => {
        serverURL.set("http://localhost:8080/");
        rootFolder = await getFolderstruture();
        console.log(rootFolder)
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
</script>

{#if unlocked}
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
                    <AddButton on:update={async ()=>{rootFolder = await getFolderstruture()}} />
                </div>
                {/if}
            </div>
            {#if sidebarVisible}
                <div class="flex-none w-56">
                    <div class="w-full h-max">
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
                    Hallo in deinem Persönlichen Wiki. Wähle doch mal eine Datei
                    aus
                </h1>
            </div>
        {/if}
    </div>
{:else}
    <div class="flex w-screen h-screen justify-center items-center">
        <PinInput.Root
            bind:value
            class="min-h-input flex h-24 w-[176px] items-center gap-2 rounded-card-sm border border-border bg-background py-1 pl-[18px] pr-1.5 shadow-mini"
            type={pinInputType}
        >
            <PinInput.Input
                class="w-5 bg-inherit text-center font-alt text-[19px] text-foreground"
            />
            <PinInput.Input
                class="w-5 bg-inherit text-center font-alt text-[19px] text-foreground"
            />
            <PinInput.Input
                class="w-5 bg-inherit text-center font-alt text-[19px] text-foreground"
            />
            <PinInput.Input
                class="w-5 bg-inherit text-center font-alt text-[19px] text-foreground"
            />
            <PinInput.HiddenInput />
            <Toggle.Root
                aria-label="toggle code visibility"
                class="inline-flex size-10 items-center justify-center rounded-[9px] bg-background transition-all hover:bg-muted active:scale-98 active:bg-dark-10 data-[state=on]:bg-muted active:data-[state=on]:bg-dark-10"
                bind:pressed={unlocked}
            >
                {#if unlocked}
                    <LockKeyhole class="size-6" />
                {:else}
                    <UnlockKeyhole class="size-6" />
                {/if}
            </Toggle.Root>
        </PinInput.Root>
    </div>
{/if}
