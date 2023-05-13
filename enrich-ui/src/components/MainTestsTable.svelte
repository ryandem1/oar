<script lang="ts">
  import { Paginator, tableMapperValues } from "@skeletonlabs/skeleton";
  import { OARServiceClient } from "$lib/client";
  import { onMount } from "svelte";
  import { isEnrichUIError, isOARServiceError } from "$lib/models";
  import { to_number } from "svelte/internal";

  const client = new OARServiceClient();

  /*
  TABLE LOAD AND PAGINATION FUNCTIONALITY
  */
  let fields = ["id", "summary", "outcome", "analysis", "resolution", "owner", "type", "app"]
  let testIDIndex: number = fields.findIndex((elem) => elem === "id");
  if (testIDIndex === -1) {
    console.error("Could not find test ID as a field in the table!");
  }

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

  /*
  SELECT FUNCTIONALITY
  */
  let selectedTestIDs: number[];
  $: selectedTestIDs = [];

  function toggleRow(row: string[]) {
    let testID = to_number(row[fields.indexOf("id")])

    if (selectedTestIDs.includes(testID)) {
      selectedTestIDs = selectedTestIDs.filter(i => i !== testID);
    } else {
      selectedTestIDs = [...selectedTestIDs, testID];
    }
  }
</script>

<style>
    .selected {
        border-color: #7d5a5a;
        background-color: #fff4f4;
        background-image: linear-gradient(90deg, silver 50%, transparent 50%), linear-gradient(90deg, silver 50%, transparent 50%), linear-gradient(0deg, silver 50%, transparent 50%), linear-gradient(0deg, silver 50%, transparent 50%);
        background-repeat: repeat-x, repeat-x, repeat-y, repeat-y;
        background-size: 15px 2px, 15px 2px, 2px 15px, 2px 15px;
        background-position: left top, right bottom, left bottom, right   top;
        animation: border-dance 1s infinite linear;
    }
    @keyframes border-dance {
        0% {
            background-position: left top, right bottom, left bottom, right   top;
        }
        100% {
            background-position: left 15px top, right 15px bottom , left bottom 15px , right   top 15px;
        }
    }

</style>

<div class="card bg-surface-50 shadow-xl p-2 outline-double outline-4 outline-surface-400">
  <div class="table-container w-full">
    <table class="table-auto table-compact table-interactive w-full">
      <thead>
      <tr>
        {#each fields as header}
          <th>{header}</th>
        {/each}
      </tr>
      </thead>
      <tbody>
      {#each paginatedSource as row, i}
        <tr
          class:selected={selectedTestIDs.includes(row[fields.indexOf("id")])}
          on:click={() => toggleRow(row)}
        >
          {#each fields as field, j}
            {#if field === "summary"}
              <td class="pr-4 pb-4">{row[j]}</td>
            {:else}
              {#if row[j] === undefined}
                <td class="pr-4 pb-4 text-center">-</td>
              {:else}
                <td class="pr-4 pb-4 text-center">{row[j]}</td>
              {/if}
            {/if}
          {/each}
        </tr>
      {/each}
      </tbody>
    </table>
  </div>
  <Paginator
    bind:settings={page}
    buttonClasses="btn-icon bg-surface-300"
  />
</div>
