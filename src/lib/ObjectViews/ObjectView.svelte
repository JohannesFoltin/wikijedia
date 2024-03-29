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

    onMount(async () => {
        await getObject(currentObject.get().Name);
        nameBuffer = currentObject.get().Name;
        currentObject.subscribe(async (value) => {
            if (value !== null && value.Name !== object.Name) {
                nameBuffer = value.Name;
                await getObject(value.Name);
            }
        });
    });

    $: console.log(object);

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
            const response = await fetch($serverURL + "object/" + object.Name, {
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

    const updateName = async () => {
        let tmp = { Name: nameBuffer };
        console.log(tmp);
        const response = await fetch(
            $serverURL + "object/" + object.Name + "/name",
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
                <PictureView data={$serverURL + "/object/data/" + object.Name}
                ></PictureView>
            {:else if object.Type === "application/pdf"}
                <iframe
                    src={$serverURL + "/object/data/" + object.Name}
                    class="w-full h-full p-4"
                    title={object.Name}
                ></iframe>
            {:else}
                <p>Not supported</p>
            {/if}
        </div>
    </div>
{/if}
