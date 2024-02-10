<script lang="ts">
    import { onMount } from 'svelte';

    let inputValue = '';
    let isLoading:boolean = false;
    let isSuccess = false;

    const sendRequest = async () => {
        isLoading = true;

        // Simulating delay for demonstration purposes
        await new Promise(resolve => setTimeout(resolve, 1000));

        // Make PUT request to API with inputValue
        try {
            // Make PUT request here
            // ...

            // If request is successful
            isSuccess = true;
        } catch (error) {
            console.error(error);
        } finally {
            isLoading = false;
        }
    };

    onMount(() => {
        const input = document.getElementById('input');

        input.addEventListener('input', () => {
            // Reset isSuccess when input changes
            isSuccess = false;
        });

        let timeout;
        input.addEventListener('input', () => {
            clearTimeout(timeout);
            timeout = setTimeout(sendRequest, 1000);
        });
    });
</script>

<div class="flex items-start">
    <input id="input" class="w-52 h-12 text-xl border-b-2 border-b-gray-600 border-solid;" bind:value={inputValue} />
    {#if isLoading}
        <div class="loading-icon">Loading...</div>
    {:else if isSuccess}
        <div class="success-icon">Success!</div>
    {/if}
</div>

