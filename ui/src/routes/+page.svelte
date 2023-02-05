<script lang="ts">
    import type { Test } from "$lib/models";
    import { Accordion } from 'flowbite-svelte';
    import { onMount } from "svelte";
    import TestAccordionItem from "$lib/components/TestAccordionItem.svelte";
    import { Outcome, Analysis, Resolution } from "$lib/consts.js";

    let tests: Test[] = []

    onMount(async function () {
        const response = await fetch("http://localhost:8080/tests")
        const data = await response.json()

        data.forEach(rawTest => {
            let test: Test = {
                id: rawTest["id"],
                summary: rawTest["summary"],
                outcome: Outcome[rawTest["outcome"]],
                analysis: Analysis[Object.keys(Analysis).find(key => key === rawTest["analysis"])],
                resolution: Resolution[rawTest["resolution"]],
                doc: rawTest["doc"],
            }
            tests = [...tests, test]
        })
    })
</script>

<div class="p-8">
    <Accordion>
        {#each tests as test}
            <TestAccordionItem test={test}/>
        {/each}
    </Accordion>
</div>
