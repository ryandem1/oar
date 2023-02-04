<script lang="ts">
    import { AccordionItem } from 'flowbite-svelte'
    import OutcomeBadge from "$lib/components/OutcomeBadge.svelte";
    import type { Test } from "$lib/models";

    export let test: Test
</script>

<AccordionItem>
    <span slot="header">{test.summary}
        <span id="oar">
            <OutcomeBadge outcome="{test.outcome}"/>
            Analysis: {test.analysis}
            Resolution: {test.resolution}
        </span>
    </span>
</AccordionItem>

<style>

</style>
