<script lang="ts">
    import { onMount} from "svelte";
    import markdownit from "markdown-it";
    import "../app.css";


    export let data : string;

    let html = "";


    let isEditing = false;
    onMount(() => {
        console.log("onMount");
        console.log(data);
        updatePreview()
    });

    const md = markdownit({
        breaks: true,
    });

    $: console.log(data);


    function updatePreview() {
        html = md.render(data);
    }

    function toggleEditing() {
        isEditing = !isEditing;
    }


    function onPaste(event: ClipboardEvent) {
        console.log("paste Type: " + event.clipboardData.files[0].type);
    }

    function sendObjectToServer() {
        console.log("sendObjectToServer");
        console.log(data);
    }
</script>

{#if isEditing}
    <div class="flex flex-row h-full w-full">
        <div class="flex-1 pr-5">
                        <textarea
                                class="w-full h-full resize-none b
                            order-black border-2 rounded"
                                bind:value={data}
                                on:paste={onPaste}
                                on:input={updatePreview}
                        ></textarea>
        </div>
        <div class="flex-1 pl-5 overflow-y-auto">
            {@html html}
        </div>
    </div>
{:else}
    <div class="flex-1">
        <div class="flex-1 overflow-y-auto">
            {@html html}
        </div>
    </div>
{/if}
<button
        class="fixed bottom-4 right-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        on:click={toggleEditing}
>
    {#if isEditing}
        Preview
    {:else}
        Edit
    {/if}
</button>
<button
        class="fixed bottom-4 right-28 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        on:click={sendObjectToServer}
>
    Save
</button>

