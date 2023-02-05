<script lang="ts">
    import { Button, Dropdown, DropdownItem, Chevron } from 'flowbite-svelte'
    import { Analysis } from "$lib/consts";

    export let analysis: Analysis
    let dropdownOpen = false
</script>

<Button color="light" pill={true}><Chevron>{analysis}</Chevron></Button>
<Dropdown bind:open={dropdownOpen}>
    {#each Object.keys(Analysis) as option}
        <DropdownItem on:click={() => {dropdownOpen = false; analysis = Analysis[option]}}>
            {Analysis[option]}
        </DropdownItem>
    {/each}
</Dropdown>
