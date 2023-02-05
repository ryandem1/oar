<script lang="ts">
    import {TableBodyCell, TableBodyRow, Button} from 'flowbite-svelte'
    import OutcomeBadge from "$lib/components/OutcomeBadge.svelte";
    import AnalysisDropdown from "$lib/components/AnalysisDropdown.svelte";
    import type {Test} from "$lib/models";
    import ResolutionDropdown from "$lib/components/ResolutionDropdown.svelte";

    export let test: Test
    let maxSummaryLength = 30  // Maximum displayed summary characters
</script>

<TableBodyRow>
    <TableBodyCell>
        {test.summary.length > maxSummaryLength ? test.summary.substring(0, maxSummaryLength) + "..." : test.summary}
    </TableBodyCell>
    <TableBodyCell><OutcomeBadge outcome={test.outcome}/></TableBodyCell>
    <TableBodyCell><AnalysisDropdown analysis={test.analysis}/></TableBodyCell>
    <TableBodyCell><ResolutionDropdown resolution={test.resolution}/></TableBodyCell>
    <TableBodyCell><Button>Enrich</Button></TableBodyCell>
</TableBodyRow>
