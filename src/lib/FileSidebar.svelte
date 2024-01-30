<script>
    export let data = {};
    export let indent = 0;

    function setInhalt() {
        console.log(data);
    }
</script>

<button style="padding-left: {indent}px" on:click={setInhalt}>
    {data.Name}
</button>
