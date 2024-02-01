<script lang="js">
    import { objectID } from "./store";
    import { onMount } from "svelte";
    import markdownit from "markdown-it";
    import "../app.css";
    import { basicSetup, EditorView } from "codemirror";
    import { markdown } from "@codemirror/lang-markdown";
    import { languages } from "@codemirror/language-data";

    let markdownText = "# Markdown preview";

    let html = "";
    let currentObjectID = "";

    let nameField = "";
    let nameFieldIsLoading = false;
    let safeTimeoutForName;

    let isEditing = false;

    let htmlEditorElement;

    onMount(() => {
        updatePreview();
    });

    function updatePreview() {
        const md = markdownit({
            breaks: true,
        });
        html = md.render(markdownText);
    }
    
    $:htmlEditorElement && initEditor();

 /*    $:if (view !== undefined) {
        markdownText = view.state.doc.toString();
        updatePreview();
        console.log(markdownText)
    } */

    let view;

    function initEditor() {
        console.log(htmlEditorElement);
        view = new EditorView({
            doc: markdownText,
            extensions: [
                basicSetup,
                markdown({ codeLanguages: languages }),
            ],
            parent: htmlEditorElement,
        });
        view.state.onChange(() => {
            markdownText = view.state.doc.toString();
            updatePreview();
        });
        console.log(view.state.doc.toString());
    }
    function toggleEditing() {
        isEditing = !isEditing;
    }

    objectID.subscribe(async (value) => {
        if (value !== "") {
            const url = "http://test.johafo.de:8080/object/" + value;
            try {
                const response = await fetch(url);
                const data = await response.json();
                console.log(data);
                nameField = data.Name;
                markdownText = data.Data;
                updatePreview();
            } catch (error) {
                console.error(error);
            }
            currentObjectID = value;
        }
    });

    const sendRequest = async () => {
        try {
            // Make PUT request here
            // ...

            // If request is successful
            nameFieldIsLoading = false;
        } catch (error) {
            console.error(error);
        }
    };

    function onFileNameInput() {
        nameFieldIsLoading = true;
        clearTimeout(safeTimeoutForName);
        safeTimeoutForName = setTimeout(sendRequest, 1000);
    }

    async function save() {
        const data = { Data: markdownText };
        try {
            const response = await fetch(
                "http://test.johafo.de:8080/object/" +
                    currentObjectID +
                    "/data",
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(data),
                },
            );
            if (response.ok) {
                console.log("Markdown text posted successfully");
            } else {
                console.error("Failed to post markdown text");
            }
        } catch (error) {
            console.error(error);
        }
    }
</script>

<div class="flex-col w-full h-full">
    <div class="flex items-end">
        <input
            id="input"
            class="w-96 h-12 text-xl border-b-2 border-b-gray-600 border-solid;"
            on:input={onFileNameInput}
            bind:value={nameField}
        />
        {#if nameFieldIsLoading}
            <div class="loading-icon">Loading...</div>
        {:else}
            <div class="success-icon">Success!</div>
        {/if}
    </div>
    {#if isEditing}
        <div class="flex flex-row h-full w-full">
            <div class="flex-1 pr-5">
                <div bind:this={htmlEditorElement}></div>
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
        on:click={save}
    >
        Save
    </button>
</div>
