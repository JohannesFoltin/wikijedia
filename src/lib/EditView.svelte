<script lang="ts">
    import MarkdownEditorField from './MarkdownEditorField.svelte';

    import {currentObject, serverURL} from "./store";
    import {createEventDispatcher, onMount} from "svelte";
    import "../app.css";
    import {BackendObject} from "$lib/structure";

    const dispatch = createEventDispatcher();

    let object : BackendObject;

    let nameFieldIsLoading = false;
    let safeTimeoutForName;

    onMount(async () => {
        await getObject(currentObject.get().ID);
        currentObject.subscribe(async (value) => {
            if (value !== null && value.ID !== object.ID) {
                await getObject(value.ID);
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

    const updateName = async () => {
        let tmp = {Name: object.Name};
        console.log(tmp);
        try {
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
            dispatch("updateName");
            nameFieldIsLoading = false;
        } catch (error) {
            console.error(error);
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
                    bind:value={object.Name}
            />
            {#if nameFieldIsLoading}
                <div class="loading-icon">Loading...</div>
            {:else}
                <div class="success-icon">Success!</div>
            {/if}
        </div>
        <div class="w-full h-full">
            {#if object.Type === "MD"}
                <MarkdownEditorField data={object.Data}></MarkdownEditorField>
            {:else if object.Type === "PNG"}
                <img src={object.Data} alt="Preview"/>
            {/if}
        </div>
    </div>
{/if}
