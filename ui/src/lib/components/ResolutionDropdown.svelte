<script lang="ts">
    import { Button, Dropdown, DropdownItem, Chevron } from 'flowbite-svelte'
    import { Resolution } from "$lib/consts";

    export let resolution: Resolution
    let dropdownOpen = false
</script>

<Button color="light" pill={true}><Chevron>{resolution}</Chevron></Button>
<Dropdown bind:open={dropdownOpen}>
    {#each Object.keys(Resolution) as option}
        <DropdownItem on:click={() => {dropdownOpen = false; resolution = Resolution[option]}}>
            {Resolution[option]}
        </DropdownItem>
    {/each}
</Dropdown>
