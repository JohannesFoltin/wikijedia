<script>
    import lex from "lexical";
    const { createEditor, LineBreakNode, ParagraphNode, TextNode } = lex;
    import { onMount } from "svelte";
    import lexicalList from "@lexical/list";

    const { ListItemNode, ListNode } = lexicalList;

    import lexicalRichtText from "@lexical/rich-text";
    import lexicalPlainText from "@lexical/plain-text";
    const { registerPlainText } = lexicalPlainText;

    let editorElement;

    onMount(() => {
        const config = {
            nodes: [TextNode, LineBreakNode, ParagraphNode],
            onError: console.error,
        };
        const editor = createEditor(config);
        editor.update(() => {
            editor.setRootElement(editorElement);
        });
        registerPlainText(editor);
        console.log(editor.getEditorState());
        console.log(editor.isEditable());
        editor.update(() => {
            // Get the RootNode from the EditorState
            const root = lex.$getRoot();

            // Get the selection from the EditorState
            const selection = lex.$getSelection();

            // Create a new ParagraphNode
            const paragraphNode = lex.$createParagraphNode();

            // Create a new TextNode
            const textNode = lex.$createTextNode("Hello world");

            // Append the text node to the paragraph
            paragraphNode.append(textNode);

            // Finally, append the paragraph to the root
            root.append(paragraphNode);
            editor.setEditable(true);
        });
        editor.registerUpdateListener(({ editorState }) => {
            editorState.read(() => {
                const root = lex.$getRoot();
                console.log(root.getTextContent());
            });
        });
    });
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>
<div contenteditable="true" bind:this={editorElement}></div>
