<script lang="ts">
  import { DropdownMenu } from "bits-ui";
  import { fly } from "svelte/transition";
  import { PlusCircle, FilePlus, FolderPlus } from "lucide-svelte";
  import { serverURL } from "./store";
  import { createEventDispatcher } from "svelte";
  import Dialog from "./Dialog.svelte";
  import { Upload } from "lucide-svelte";

  const depatch = createEventDispatcher();

  let dialog = false;

  async function addFile() {
    dialog = true;
    console.log("addFile");
    const url = $serverURL + "object";
    let data;

    data = {
      Name: "New Object",
      Type: "MD",
      Data: "# Hello World",
      FolderID: 1,
    };

    try {
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        depatch("update");
        console.log("Post request successful");
      } else {
        console.error("Post request failed");
      }
    } catch (error) {
      console.error("An error occurred:", error);
    }
  }

  async function addFolder() {
    console.log("addFolder");
    const url = $serverURL + "folder";

    const data = { Name: "Hallo Welt", ParentID: 1 };

    console.log(JSON.stringify(data));

    try {
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (response.ok) {
        depatch("update");
        console.log("Post request successful");
      } else {
        console.error("Post request failed");
      }
    } catch (error) {
      console.error("An error occurred:", error);
    }
  }

  let fileinput : HTMLInputElement;

  const onFileSelected = (e : Event) => {
    console.log("file selected");
    let image = e.target;
    let reader = new FileReader();
    reader.readAsDataURL(image);
    console.log("reader");
    reader.onload = (e) => {
      console.log("loaded");
      console.log(e.type);

    };
  };
</script>

<DropdownMenu.Root>
  <DropdownMenu.Trigger
    class="focus-visible inline-flex h-8 w-10  items-center justify-center rounded-full border border-border-input bg-background text-sm font-medium text-foreground shadow-btn hover:bg-muted focus-visible:ring-2 focus-visible:ring-foreground focus-visible:ring-offset-2 focus-visible:ring-offset-background active:scale-98"
  >
    <PlusCircle class="h-6 w-6 text-foreground" />
  </DropdownMenu.Trigger>
  <DropdownMenu.Content
    class="w-full max-w-[229px] z-10 rounded-xl border border-muted bg-background bg-white px-1 py-1.5 shadow-popover"
    transition={fly}
    sideOffset={4}
    overlap={true}
  >
    <DropdownMenu.Item
      class="flex h-10 select-none items-center rounded-button rounded-xl py-3 pl-3 pr-1.5 text-sm font-medium !ring-0 !ring-transparent data-[highlighted]:bg-muted hover:bg-gray-200"
      on:click={addFile}
    >
      <div class="flex items-center">
        <FilePlus class="mr-2 size-5 text-foreground-alt" />
        Neue Datei
      </div>
    </DropdownMenu.Item>
    <DropdownMenu.Item
      class="flex h-10 select-none items-center rounded-button rounded-xl hover:bg-gray-200 py-3 pl-3 pr-1.5 text-sm font-medium !ring-0 !ring-transparent data-[highlighted]:bg-muted"
      on:click={addFolder}
    >
      <div class="flex items-center">
        <FolderPlus class="mr-2 size-5 text-foreground-alt" />
        Neuer Ordner
      </div>
    </DropdownMenu.Item>
    <DropdownMenu.Item
      class="flex h-10 select-none items-center rounded-button rounded-xl hover:bg-gray-200 py-3 pl-3 pr-1.5 text-sm font-medium !ring-0 !ring-transparent data-[highlighted]:bg-muted"
      on:click={() => {fileinput.click();}}
    >
      <div class="flex items-center">
        <Upload class="mr-2 size-5 text-foreground-alt" />
        Neues File
      </div>
      <input
        style="display:none"
        type="file"
        accept=".jpeg, .png"
        on:change={(e) => {
          console.log("tr5ewst");
          onFileSelected(e);
        }}
        bind:this={fileinput}
      />
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
{#if dialog}
  <Dialog></Dialog>
{/if}
