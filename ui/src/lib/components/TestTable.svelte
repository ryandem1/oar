<script lang="ts">
    import {Table, TableBody, TableHead, TableHeadCell} from 'flowbite-svelte'
    import type {Test} from "$lib/models";
    import TestTableBodyRow from "$lib/components/TestTableBodyRow.svelte";

    export let tests: Test[]
</script>

<Table>
    <TableHead>
        <TableHeadCell>Summary</TableHeadCell>
        <TableHeadCell>Outcome</TableHeadCell>
        <TableHeadCell>Analysis</TableHeadCell>
        <TableHeadCell>Resolution</TableHeadCell>
    </TableHead>
    <TableBody>
        <TableBody>
            {#each tests as test}
            <TestTableBodyRow test={test}/>
            {/each}
        </TableBody>
    </TableBody>
</Table>
