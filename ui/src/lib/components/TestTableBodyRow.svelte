<script lang="ts">
    import {TableBodyCell, TableBodyRow} from 'flowbite-svelte'
    import OutcomeBadge from "$lib/components/OutcomeBadge.svelte";
    import AnalysisDropdown from "$lib/components/AnalysisDropdown.svelte";
    import type {Test} from "$lib/models";
    import ResolutionDropdown from "$lib/components/ResolutionDropdown.svelte";

    export let test: Test
</script>

<TableBodyRow>
    <TableBodyCell>{test.summary}</TableBodyCell>
    <TableBodyCell><OutcomeBadge outcome={test.outcome}/></TableBodyCell>
    <TableBodyCell><AnalysisDropdown analysis={test.analysis}/></TableBodyCell>
    <TableBodyCell><ResolutionDropdown resolution={test.resolution}/></TableBodyCell>
</TableBodyRow>
