<script lang="js">
    import { objectID, serverURL } from "./store";
    import { onMount } from "svelte";
    import markdownit from "markdown-it";
    import { createEventDispatcher } from "svelte";
    import "../app.css";

    const dispatch = createEventDispatcher();

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
    let object;

    let html = "";

    let nameFieldIsLoading = false;
    let safeTimeoutForName;

    let isEditing = false;

    onMount(async () => {
        await getObject(objectID.get());
    objectID.subscribe(async (value) => {
        if (value !== "" && value !== object.ID) {
            await getObject(value);
            updatePreview();
        }
    }); 
    });
    const md = markdownit({
        breaks: true,
    });

    $: console.log(object);

    $:if(object !== undefined) updatePreview();

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



    async function getObject(value) {
        try {
            const url = $serverURL + "object/" + value;
            const response = await fetch(url);
            const data = await response.json();
            console.log("hallo");
            object = data;
        } catch (error) {
            console.error(error);
        }
    }

    const sendObjectToServer = async () => {
        try {
            const response = await fetch($serverURL + "object/" + object.ID, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(object),
            });
            console.log(response);
        } catch (error) {
            console.error(error);
        }
    };  

    const updateName = async () =>{
        let tmp = {Name: object.Name};
        console.log(tmp);
        try {
            const response = await fetch($serverURL + "object/" + object.ID + "/name", {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(tmp),
            });
            console.log(response);
            dispatch("updateName");
            nameFieldIsLoading = false;
        } catch (error) {
            console.error(error);
        }
    }

    function onFileNameInput() {
        nameFieldIsLoading = true;
        clearTimeout(safeTimeoutForName);
        safeTimeoutForName = setTimeout(updateName, 500);
    }
</script>

{#if object == undefined}
    <p class="flex justify-center items-center">Loading...</p>
{:else}
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
            on:click={sendObjectToServer}
        >
            Save
        </button>
    </div>
{/if}
