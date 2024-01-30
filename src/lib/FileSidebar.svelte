<script>
    import{ objectID } from './store';

    export let data = {};
    export let indent = 0;

    function setInhalt() {
        console.log(data.ID);
        objectID.set(data.ID);
    }
</script>

<button style="padding-left: {indent}px" on:click={setInhalt}>
    {data.Name}
</button>
