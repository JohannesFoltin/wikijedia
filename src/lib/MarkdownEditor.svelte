<script lang="js">
    import { objectID } from './store';
    import { onMount } from 'svelte';
    import markdownit from 'markdown-it';
    import "../app.css";

    export let markdownText = '# Markdown preview';

    objectID.subscribe((value) => {
        console.log(value);
        if(value !== ''){
            console.log(value);
            const url = "http://test.johafo.de:8080/object/" + value;
            console.log(url);
            async function fetchData() {
                try {
                    const response = await fetch(url);
                    const data = await response.json();
                    console.log(data);
                    markdownText = data.Data;
                    updatePreview();
                } catch (error) {
                    console.error(error);
                }
            }

            fetchData();
        }
    });

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

<div class="flex w-full h-full">
    {#if isEditing}
        <div class="flex-1 pr-5">
            <textarea class="w-full h-full resize-none border-black border-2 rounded" bind:value={markdownText} on:input={updatePreview}></textarea>
        </div>
        <div class="flex-1 pl-5 overflow-y-auto">
            {@html html}
        </div>
    {:else}
        <div class="flex-1">
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
