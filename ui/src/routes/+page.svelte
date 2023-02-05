<script lang="ts">
    import type {Test} from "$lib/models";
    import {onMount} from "svelte";
    import {Analysis, Outcome, Resolution} from "$lib/consts.js";
    import TestTable from "$lib/components/TestTable.svelte";
    import OarNavbar from "$lib/components/OarNavbar.svelte";

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
                resolution: Resolution[Object.keys(Resolution).find(key => key === rawTest["resolution"])],
                doc: rawTest["doc"],
            }
            tests = [...tests, test]
        })
    })
</script>

<div class="p-8 h-max">
    <OarNavbar/>
    <TestTable divClass="overflow-x-auto h-max" tests={tests}/>
</div>

<style>
    :global(body) {
        background-color: rgba(54, 127, 211, 0.6)
    }
</style>
