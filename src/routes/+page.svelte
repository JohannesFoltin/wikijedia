<script>
    import { onMount } from "svelte";

    import { createEditor } from "svelte-editorjs";
    import Header from '@editorjs/header';

    const { editor, data, isReady } = createEditor({
        tools: {
            header: Header,
        },
    });

    $: {
        if (isReady) {
            console.log("Saved data:", $data);
        }
    }
    onMount(() => {});
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>
<div class="editor" use:editor />

<style>
    .editor {
        border: 2px solid #ccc;
    }
</style>
