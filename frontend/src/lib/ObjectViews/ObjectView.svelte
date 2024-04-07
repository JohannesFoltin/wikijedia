<script lang="ts">
    import MarkdownEditorField from "./MarkdownEditorView.svelte";

    import { currentObject, serverURL } from "../store";
    import { createEventDispatcher, onMount } from "svelte";
    import "../../app.css";
    import type { BackendObject } from "../types";
    import PictureView from "./PictureView.svelte";

    const dispatch = createEventDispatcher();

    let object: BackendObject;

    let nameFieldIsLoading = false;
    let safeTimeoutForName;
    let nameBuffer = "";

    // Lädt das Objekt vom Server und sollte sich was ändern, wird das Objekt neu geladen
    onMount(async () => {
        await getObject(currentObject.get().ID);
        nameBuffer = currentObject.get().Name;
        currentObject.subscribe(async (value) => {
            if (value !== null && value.ID !== object.ID) {
                nameBuffer = value.Name;
                await getObject(value.ID);
            }
        });
    });

    $: console.log(object);

    // Holt das Objekt vom Server
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

    // Sendet das Objekt an den Server
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

    // Aktualisiert den Namen des Objekts
    const updateName = async () => {
        let tmp = { Name: nameBuffer };
        console.log(tmp);
        const response = await fetch(
            $serverURL + "object/" + object.ID + "/name",
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(tmp),
            },
        );
        console.log(response);
        if (response.ok) {
            dispatch("updateName");
            object.Name = nameBuffer;
            nameFieldIsLoading = false;
        }else{
            console.error("Post request failed");
            nameBuffer = object.Name;
        }
    };

    // Funktion die aufgerufen wird, wenn der Name geändert wird und einen Offset zum speichern des Namens setzt
    function onFileNameInput() {
        nameFieldIsLoading = true;
        clearTimeout(safeTimeoutForName);
        safeTimeoutForName = setTimeout(updateName, 500);
    }
</script>

{#if object === undefined}
    <p class="flex justify-center items-center">Loading...</p>
{:else}
    <div class="flex-col w-full h-full">
        <div class="flex items-end">
            <input
                id="input"
                class="w-96 h-12 text-xl border-b-2 border-b-gray-600 border-solid;"
                on:input={onFileNameInput}
                bind:value={nameBuffer}
            />
            {#if nameFieldIsLoading}
                <div class="loading-icon">Loading...</div>
            {:else}
                <div class="success-icon">Success!</div>
            {/if}
        </div>
        <div class="w-full h-full">
            {#if object.Type === "MD"}
                <MarkdownEditorField
                    bind:data={object.Data}
                    on:update={sendObjectToServer}
                ></MarkdownEditorField>
            {:else if object.Type === "image/png" || object.Type === "image/jpeg"}
                <PictureView data={$serverURL + "/object/data/" + object.ID}
                ></PictureView>
            {:else if object.Type === "application/pdf"}
                <iframe
                    src={$serverURL + "/object/data/" + object.ID}
                    class="w-full h-full p-4"
                    title={object.Name}
                ></iframe>
            {:else}
                <p>Not supported</p>
            {/if}
        </div>
    </div>
{/if}
