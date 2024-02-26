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

  let fileinput: HTMLInputElement;

  async function uploadFile(file :File) {
    const url = 'http://localhost:8080/upload';
    const formData = new FormData();

    formData.append('file', file);

    const response = await fetch(url, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    depatch("update");
  }

  const onFileSelected = (e: Event) => {
    console.log(e);
    let file = (e.target as HTMLInputElement).files[0];
            let reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = e => {
                uploadFile(file);
                 //avatar = e.target.result
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
      on:click={() => {
        fileinput.click();
      }}
    >
      <div class="flex items-center">
        <Upload class="mr-2 size-5 text-foreground-alt" />
        Neues File
      </div>
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
{#if dialog}
  <Dialog></Dialog>
{/if}
<input
  class="hidden"
  type="file"
  accept="image/png, image/jpeg, application/pdf"
  name="file"
  id="file"
  on:change={(e)=>onFileSelected(e)}
    bind:this={fileinput}
  required
/>
