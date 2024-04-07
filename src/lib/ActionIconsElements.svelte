<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import { Trash2 } from "lucide-svelte";
    import { ArrowUpDown } from "lucide-svelte";

    const dispatch = createEventDispatcher();
</script>

<div class="flex items-center justify-end gap-1 z-10">
    <button on:click={() => dispatch("delelte")}>
        <Trash2 />
    </button>
    <button on:click={() => dispatch("move")}>
        <ArrowUpDown />
    </button>
</div>
