<script lang="ts">
    import { Dialog, Separator, Label } from "bits-ui";
    //import { flyAndScale } from "bits-ui/utils";
    import { fly } from "svelte/transition";
    import { fade } from "svelte/transition";
    import type { BackendFolder } from "../types";
    import { createEventDispatcher } from "svelte";
    import FolderMoveDialogElement from "./FolderMoveDialogElement.svelte";

    const dispatch = createEventDispatcher();

    export let dialogOpen = true;

    export let rootFolder: BackendFolder;


    function closeDialog() {
        dispatch("error");
    }

    function save(id:number){
      console.log(id);
        dispatch("save", id);
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
          >
            <div class="text-2xl font-bold">
            In anderen Ordner verschieben
          </div>
          </Dialog.Title
        >
        <div class="flex flex-col items-start gap-1 pb-11 pt-7">
          <Label.Root for="apiKey" class="text-lg font-medium">
            <div >
              Doppelklick auf den passenden Überordner
            </div>
            </Label.Root
            >
        <div class="my-2"/>
          <div class="relative w-full">
            <div class="">
              <FolderMoveDialogElement folder={rootFolder} indent={0} on:selected={({detail}) => save(detail)} />
            </div>
          </div>
        </div>          
      </Dialog.Content>
    </Dialog.Portal>
  </Dialog.Root>