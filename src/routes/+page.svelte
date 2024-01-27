<script lang="js">
    import { onMount } from 'svelte';
    import markdownit from 'markdown-it';

    let markdownText = '## markdown preview';
    let html = '';

    onMount(() => {
        updatePreview();
    });

    function updatePreview() {
       // const markdownIt = mdurl.require('markdown-it')();
        const md = markdownit({
            breaks: true,
        });
        html = md.render(markdownText);
    }

</script>

<style>
    textarea {
        width: 100%;
        height: 200px;
    }
</style>

<textarea bind:value={markdownText} on:input={updatePreview}></textarea>
<div>
    {@html html}
</div>
