<script>
    import { AccordionItem } from 'flowbite-svelte'
    import { Test } from "$lib/models.js";

    export let test: Test
</script>

<AccordionItem>
    <span slot="header">{test.summary}, {test.outcome}, {test.analysis}, {test.resolution}</span>
</AccordionItem>
