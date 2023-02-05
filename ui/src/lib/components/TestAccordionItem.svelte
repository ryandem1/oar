<script lang="ts">
    import { AccordionItem } from 'flowbite-svelte'
    import OutcomeBadge from "$lib/components/OutcomeBadge.svelte";
    import AnalysisDropdown from "$lib/components/AnalysisDropdown.svelte";
    import type { Test } from "$lib/models";

    export let test: Test
</script>

<div id="testItem">
    <AccordionItem>
    <span slot="header">{test.summary}
        <span id="oar">
            <div id="outcome"><OutcomeBadge outcome="{test.outcome}"/></div>
            <AnalysisDropdown analysis={test.analysis}/>
            Resolution: {test.resolution}
        </span>
    </span>
    </AccordionItem>
</div>

<style>
    #outcome {
        position: relative;
        left: 100%;
    }
</style>
