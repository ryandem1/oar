<script lang="ts">
    import { Button, Dropdown, DropdownItem, Chevron } from 'flowbite-svelte'
    import { Resolution } from "$lib/consts";

    export let resolution: Resolution
</script>

<Button><Chevron>{resolution}</Chevron></Button>
<Dropdown>
    {#each Object.keys(Resolution) as option}
        <DropdownItem on:click={() => {resolution = Resolution[option]}}>{Resolution[option]}</DropdownItem>
    {/each}
</Dropdown>
