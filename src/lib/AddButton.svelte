<script lang="ts">
  import { DropdownMenu } from "bits-ui";
  import { fly } from "svelte/transition";
  import { PlusCircle, FilePlus, FolderPlus } from "lucide-svelte";
  import { serverURL } from "./store";

  async function addFile() {
    console.log("addFile");
    const url = $serverURL + "object";
    let data;

    data = { Name: "New Object", Data: "# Hello World", FolderID: 1 };

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

<DropdownMenu.Root>
  <DropdownMenu.Trigger
    class="focus-visible inline-flex h-8 w-10  items-center justify-center rounded-full border border-border-input bg-background text-sm font-medium text-foreground shadow-btn hover:bg-muted focus-visible:ring-2 focus-visible:ring-foreground focus-visible:ring-offset-2 focus-visible:ring-offset-background active:scale-98"
  >
    <PlusCircle class="h-6 w-6 text-foreground" />
  </DropdownMenu.Trigger>
  <DropdownMenu.Content
    class="w-full max-w-[229px] z-10 rounded-xl border border-muted bg-background bg-white px-1 py-1.5 shadow-popover"
    transition={fly}
    sideOffset={8}
    overlap={true}
  >
    <DropdownMenu.Item
      class="flex h-10 select-none items-center rounded-button py-3 pl-3 pr-1.5 text-sm font-medium !ring-0 !ring-transparent data-[highlighted]:bg-muted"
    >
      <div class="flex items-center">
        <FilePlus class="mr-2 size-5 text-foreground-alt" />
        Add File
      </div>
    </DropdownMenu.Item>
    <DropdownMenu.Item
      class="flex h-10 select-none items-center rounded-button py-3 pl-3 pr-1.5 text-sm font-medium !ring-0 !ring-transparent data-[highlighted]:bg-muted"
    >
      <div class="flex items-center">
        <FolderPlus class="mr-2 size-5 text-foreground-alt" />
        Add Folder
      </div>
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
