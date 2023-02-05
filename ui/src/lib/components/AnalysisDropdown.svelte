<script lang="ts">
    import { Button, Dropdown, DropdownItem, Chevron } from 'flowbite-svelte'
    import { Analysis } from "$lib/consts";

    export let analysis: Analysis
</script>

<Button><Chevron>{analysis}</Chevron></Button>
<Dropdown>
    {#each Object.keys(Analysis) as option}
        <DropdownItem on:click={() => {analysis = Analysis[option]}}>{Analysis[option]}</DropdownItem>
    {/each}
</Dropdown>
