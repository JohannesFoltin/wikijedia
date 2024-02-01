<script lang="js">
    import { objectID } from "./store";
    import { onMount } from "svelte";
    import markdownit from "markdown-it";
    import "../app.css";
    import { basicSetup, EditorView } from "codemirror";
    import { markdown } from "@codemirror/lang-markdown";
    import { languages } from "@codemirror/language-data";

    /**
     *
     * @param {Object} Object Das Objekt, welches auch so im Backend gespeichert wird
     * @param {number} object.ID Die ID des Objekts
     * @param {string} object.Name Name des Objekts
     * @param {string} object.Type Type des Objekts
     * @param {string} object.Data Daten des Objekts
     * @param {number} object.FolderID ID des Ordners, in dem das Objekt liegt
     *
     */
    let object = getObject({ $objectID });

    let html = "";

    let nameFieldIsLoading = false;
    let safeTimeoutForName;

    let isEditing = false;

    let htmlEditorElement;

    const md = markdownit({
        breaks: true,
    });

   $: object.Data !== undefined && updatePreview();

    function updatePreview() {
        html = md.render(object.Data);
    }

    //  htmlEditorElement && initEditor();

    // /*    $:if (view !== undefined) {
    //     markdownText = view.state.doc.toString();
    //     updatePreview();
    //     console.log(markdownText)
    // } */

    // let view;

    // function initEditor() {
    //     console.log(htmlEditorElement);
    //     view = new EditorView({
    //         doc: markdownText,
    //         extensions: [basicSetup, markdown({ codeLanguages: languages })],
    //         parent: htmlEditorElement,
    //     });
    //     view.state.onChange(() => {
    //         markdownText = view.state.doc.toString();
    //         updatePreview();
    //     });
    //     console.log(view.state.doc.toString());
    // }

    function toggleEditing() {
        isEditing = !isEditing;
    }

    objectID.subscribe(async (value) => {
        if (value !== "" && value !== object.ID) {
            await getObject(value);
            updatePreview();
        }
    });

    async function getObject(value) {
        try {
            const url = "http://test.johafo.de:8080/object/" + value;
            const response = await fetch(url);
            const data = await response.json();
            return data;
        } catch (error) {
            console.error(error);
        }
    }

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
                "http://test.johafo.de:8080/object/" + object.ID + "/data",
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

{#await object}
<p>Loading...</p>
{:then}
    <div class="flex-col w-full h-full">
        <div class="flex items-end">
            <input
                id="input"
                class="w-96 h-12 text-xl border-b-2 border-b-gray-600 border-solid;"
                on:input={onFileNameInput}
                bind:value={object.Name}
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
                    <textarea
                    class="w-full h-full resize-none border-black border-2 rounded"
                    bind:value={object.Data}
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
            on:click={save}
        >
            Save
        </button>
    </div>
{/await}
