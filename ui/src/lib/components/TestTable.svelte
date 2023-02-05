<script lang="ts">
    import {Table, TableBody, TableHead, TableHeadCell} from 'flowbite-svelte'
    import type {Test} from "$lib/models";
    import TestTableBodyRow from "$lib/components/TestTableBodyRow.svelte";

    export let tests: Test[]
</script>

<Table striped={true} hoverable={true} shadow={true} style="border:1px solid rgba(54, 127, 242, 0.4);">
    <TableHead>
        <TableHeadCell><p style="color: white">Summary</p></TableHeadCell>
        <TableHeadCell><p style="color: white">Outcome</p></TableHeadCell>
        <TableHeadCell><p style="color: white">Analysis</p></TableHeadCell>
        <TableHeadCell><p style="color: white">Resolution</p></TableHeadCell>
        <TableHeadCell><span class="sr-only">Enrich</span></TableHeadCell>
    </TableHead>
    <TableBody>
        {#each tests as test}
        <TestTableBodyRow test={test}/>
        {/each}
    </TableBody>
</Table>
