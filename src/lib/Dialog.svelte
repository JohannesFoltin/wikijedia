<script lang="ts">
    import { Dialog, Separator, Label } from "bits-ui";
    //import { flyAndScale } from "bits-ui/utils";
    import { fly } from "svelte/transition";
    import { X} from "lucide-svelte";
    import { fade } from "svelte/transition";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let dialogOpen = true;

    export let title : string = "";
    export let message :string = "";
    
    let content ="";

    $:console.log(content);

    function closeDialog() {
        content = "";
        dispatch("error");
    }

    function save(){
        dispatch("save", content);
        content = "";
    }

  </script>
   
  <Dialog.Root bind:open={dialogOpen} onOutsideClick={closeDialog}>
    <Dialog.Portal>
      <Dialog.Overlay
        transition={fade}
        transitionConfig={{ duration: 150 }}
        class="fixed inset-0 z-50 bg-black/80"
      />
      <Dialog.Content
        transition={fly}
        class="fixed left-[50%] top-[50%] z-50 w-full max-w-[94%] translate-x-[-50%] translate-y-[-50%] rounded-card-lg border bg-background p-5 shadow-popover outline-none sm:max-w-[490px] md:w-full bg-white"
      >
        <Dialog.Title
          class="flex w-full items-center justify-center text-lg font-semibold tracking-tight"
          >{title}</Dialog.Title
        >
        <Separator.Root class="-mx-5 mb-6 mt-5 block h-px bg-muted" />
        <div class="flex flex-col items-start gap-1 pb-11 pt-7">
          <Label.Root for="apiKey" class="text-sm font-medium">{message}</Label.Root
          >
          <div class="relative w-full">
            <input
            bind:value={content}
              id="apiKey"
              class="inline-flex h-input w-full items-center rounded-card-sm border border-border-input bg-background px-4 text-sm placeholder:text-foreground-alt/50 hover:border-dark-40 focus:outline-none focus:ring-2 focus:ring-foreground focus:ring-offset-2 focus:ring-offset-background"
              placeholder="Name"
              type="text"
              autocomplete="off"
            />
          </div>
        </div>
        <div class="flex w-full justify-end">
          <Dialog.Close
            class="inline-flex h-input items-center justify-center rounded-input bg-dark px-[50px] text-[15px] font-semibold text-background shadow-mini hover:bg-dark/95 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-dark focus-visible:ring-offset-2 focus-visible:ring-offset-background active:scale-98"
          on:click={save}>
            Save
          </Dialog.Close>
        </div>
        <Dialog.Close
          class="absolute right-5 top-5 rounded-md focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-foreground focus-visible:ring-offset-2 focus-visible:ring-offset-background active:scale-98"
          on:click={closeDialog}
          >
          <div>
            <X class="size-5 text-foreground" />
            <span class="sr-only">Close</span>
          </div>
        </Dialog.Close>
      </Dialog.Content>
    </Dialog.Portal>
  </Dialog.Root>