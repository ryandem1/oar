<script lang="ts">
    import { AccordionItem } from 'flowbite-svelte'
    import type { Test } from "$lib/models";

    export let test: Test
</script>

<AccordionItem>
    <span slot="header">{test.summary}, {test.outcome}, {test.analysis}, {test.resolution}</span>
</AccordionItem>
