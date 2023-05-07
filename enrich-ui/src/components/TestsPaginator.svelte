<script lang="ts">
  import { Paginator, Table } from "@skeletonlabs/skeleton";
  import { OARServiceClient } from "$lib/client";
  import { onMount } from 'svelte';
  import { isOARServiceError, isEnrichUIError } from "$lib/models";
  import type { Test } from "$lib/models";

  const client = new OARServiceClient();
  $: offset = 0
  $: limit = 5
  let page = {
    offset: offset,
    limit: limit,
    size: 50,
    amounts: [5, 10, 25, 50, 100],
  };

  let headers = ["id", "summary", "outcome", "analysis", "resolution"]
  const toTestTable = (tests: Test[]): any[][] => {
    let out = [];
    tests.forEach((test) => {
      out.push([test.id, test.summary, test.outcome, test.analysis, test.resolution])
    })
    return out
  }

  let tests: Test[] = [];
  $: testTable = toTestTable(tests);

  onMount(async () => {
    const result = await client.getTests(null, offset, limit);
    if (isOARServiceError(result) || isEnrichUIError(result)) {
      console.error(result.error)
      return
    }

    tests = result.tests
  })

  // Paginator events
  function onPageChange(e: CustomEvent): void {
    console.log('event:page', e.detail);
  }

  function onAmountChange(e: CustomEvent): void {
    console.log('event:amount', e.detail);
  }
</script>


<Table source={{ head: headers, body: testTable }} class="w-full"/>
<Paginator
  settings={page}
  on:page={onPageChange}
  on:amount={onAmountChange}
  buttonClasses="btn-icon bg-surface-300"
/>
