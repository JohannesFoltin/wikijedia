<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import { Trash2 } from "lucide-svelte";
    import { ArrowUpDown } from "lucide-svelte";

    const dispatch = createEventDispatcher();
</script>

<div class="flex items-center justify-end gap-1 z-10 h-full">
    <button  on:click={() => dispatch("delelte")}>
        <Trash2 class="size-5" />
    </button>
    <button on:click={() => dispatch("move")}>
        <ArrowUpDown class="size-5"/>
    </button>
    <div class="w-2"/>
</div>
