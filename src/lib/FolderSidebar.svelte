<script>
    let folders = [
        {
            name: "Folder 1",
            files: ["File 1", "File 2", "File 3"]
        },
        {
            name: "Folder 2",
            files: ["File 4", "File 5"]
        },
        {
            name: "Folder 3",
            files: ["File 6"]
        }
    ];
</script>

<div class="w-fit">
    {#each folders as folder}
        <div class="ml-5" style="white-space: nowrap;">{folder.name}</div>
        {#each folder.files as file}
            <div class="ml-14" style="white-space: nowrap;">{file}</div>
        {/each}
    {/each}
</div>
