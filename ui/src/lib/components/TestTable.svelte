<script lang="ts">
    import {Table, TableBody, TableHead, TableHeadCell} from 'flowbite-svelte'
    import type {Test} from "$lib/models";
    import TestTableBodyRow from "$lib/components/TestTableBodyRow.svelte";

    export let tests: Test[]
</script>

<Table striped={true} hoverable={true} shadow={true} color="default" style="border:1px solid rgba(54, 127, 242, 0.4);">
    <TableHead>
        <TableHeadCell>Summary</TableHeadCell>
        <TableHeadCell>Outcome</TableHeadCell>
        <TableHeadCell>Analysis</TableHeadCell>
        <TableHeadCell>Resolution</TableHeadCell>
        <TableHeadCell><span class="sr-only">Enrich</span></TableHeadCell>
    </TableHead>
    <TableBody>
        {#each tests as test}
        <TestTableBodyRow test={test}/>
        {/each}
    </TableBody>
</Table>
