<script lang="ts">
  import { Paginator, Table, tableMapperValues } from "@skeletonlabs/skeleton";
  import { OARServiceClient } from "$lib/client";
  import { onMount } from "svelte";
  import { isEnrichUIError, isOARServiceError } from "$lib/models";
  import { toTitleCase } from "$lib/util"

  const client = new OARServiceClient();

  let fields = ["id", "summary", "outcome", "analysis", "resolution", "owner", "type", "app"]

  let testTable: string[][] = [];
  $: testTable = [];

  onMount(async () => {
    const response = await client.getTests(null, 0, 250);
    if (isEnrichUIError(response) || isOARServiceError(response)) {
      console.error(response.error)
      return
    }

    testTable = tableMapperValues(response.tests, fields.map((f) => f.toLowerCase()));
  })

  let page = {
    offset: 0,
    limit: 25,
    size: testTable.length,
    amounts: [5, 10, 25, 100],
  };

  $: {
    page.size = testTable.length;
  }

  $: paginatedSource = testTable.slice(
    page.offset * page.limit,             // start
    page.offset * page.limit + page.limit // end
  );
</script>


<Table
  interactive={true}
  source={{ head: fields.map((f) => toTitleCase(f)), body: paginatedSource }}
  element="table-auto w-full"
  regionBody="bg-surface-50"
  regionCell="pr-4 pb-4"
/>
<Paginator
  bind:settings={page}
  buttonClasses="btn-icon bg-surface-300"
/>
