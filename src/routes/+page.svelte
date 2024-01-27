<script lang="js">
    import { onMount } from 'svelte';
    import markdownit from 'markdown-it';
    import "../app.css";

    let markdownText = '# Markdown preview';
    let html = '';
    let isEditing = false;

    onMount(() => {
        updatePreview();
    });

    function updatePreview() {
        const md = markdownit({
            breaks: true,
        });
        html = md.render(markdownText);
    }

    function toggleEditing() {
        isEditing = !isEditing;
    }

</script>

<div class="flex h-screen">
    {#if isEditing}
        <div class="flex-1 p-10">
            <textarea class="w-full h-full resize-none border-black border-2 rounded" bind:value={markdownText} on:input={updatePreview}></textarea>
        </div>
        <div class="flex-1 p-10 overflow-y-auto">
            {@html html}
        </div>
    {:else}
        <div class="flex-1 p-10">
                <div class="flex-1 overflow-y-auto">
                    {@html html}
                </div>
        </div>
    {/if}
    <button class="fixed bottom-4 right-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" on:click={toggleEditing}>
        {#if isEditing}
            Preview
        {:else}
            Edit
        {/if}
    </button>
</div>
