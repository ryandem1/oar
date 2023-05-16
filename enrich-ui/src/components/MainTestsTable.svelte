<script lang="ts">
  import { Paginator } from "@skeletonlabs/skeleton";
  import { OARServiceClient } from "$lib/client";
  import { onMount } from "svelte";
  import { to_number } from "svelte/internal";
  import { refreshTestTable, selectedTestIDs } from "../stores";
  import { getTestQuery, getTestTable, getTestTableFields } from "$lib/table";

  const client = new OARServiceClient();

  /*
  TABLE LOAD AND PAGINATION FUNCTIONALITY
  */
  let fields = getTestTableFields();
  let testIDIndex: number = fields.findIndex((elem) => elem === "id");
  if (testIDIndex === -1) {
    console.error("Could not find test ID as a field in the table!");
  }

  let testTable: string[][] = [];
  $: testTable = [];

  onMount(async () => {
    refreshTestTable.subscribe(async refresh => {
      if (refresh === true) {
        refreshTestTable.set(false);
        localSelectedTestIDs = [];
        selectedTestIdxes = [];
        let tableQuery = getTestQuery();
        fields = getTestTableFields();
        testTable = await getTestTable(tableQuery, fields);
      }
    })
  })

  let page = {
    offset: 0,
    limit: 7,
    size: testTable.length,
    amounts: [7, 15, 25, 100],
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
  let localSelectedTestIDs: number[];
  let selectedTestIdxes: number[];
  $: localSelectedTestIDs = [];
  $: selectedTestIdxes = [];
  $: {
    selectedTestIDs.set(localSelectedTestIDs)
  }

  function toggleRow(row: string[], index: number) {
    let testID = to_number(row[fields.indexOf("id")])

    if (localSelectedTestIDs.includes(testID)) {
      localSelectedTestIDs = localSelectedTestIDs.filter(i => i !== testID);
      selectedTestIdxes = selectedTestIdxes.filter(i => i !== index)
    } else {
      localSelectedTestIDs = [...localSelectedTestIDs, testID];
      selectedTestIdxes = [...selectedTestIdxes, index];
    }
  }
</script>

<style>
    .selected {
        border-color: #7d5a5a;
        background-color: #fff4f4;
        background-image: linear-gradient(90deg, #5db7ee 50%, transparent 50%), linear-gradient(90deg, #5db7ee 50%, transparent 50%), linear-gradient(0deg, #5db7ee 50%, transparent 50%), linear-gradient(0deg, #5db7ee 50%, transparent 50%);
        background-repeat: repeat-x, repeat-x, repeat-y, repeat-y;
        background-size: 15px 2px, 15px 2px, 2px 15px, 2px 15px;
        background-position: left top, right bottom, left bottom, right top;
        animation: border-dance 1s infinite linear;
    }
    @keyframes border-dance {
        0% {
            background-position: left top, right bottom, left bottom, right   top;
        }
        100% {
            background-position: left 15px top, right 15px bottom, left bottom 15px, right top 15px;
        }
    }

    .top-selected {
        background-image:
                linear-gradient(90deg, #fff4f4 50%, transparent 50%),
                linear-gradient(90deg, #5db7ee 50%, transparent 50%),
                linear-gradient(0deg, #5db7ee 50%, transparent 50%),
                linear-gradient(0deg, #5db7ee 50%, transparent 50%)
        !important;
    }

    .bottom-selected {
        background-image:
                linear-gradient(90deg, #5db7ee 50%, transparent 50%),
                linear-gradient(90deg, #fff4f4 50%, transparent 50%),
                linear-gradient(0deg, #5db7ee 50%, transparent 50%),
                linear-gradient(0deg, #5db7ee 50%, transparent 50%)
        !important;
    }

    .top-and-bottom-selected {
        background-image:
                linear-gradient(90deg, #fff4f4 50%, transparent 50%),
                linear-gradient(90deg, #fff4f4 50%, transparent 50%),
                linear-gradient(0deg, #5db7ee 50%, transparent 50%),
                linear-gradient(0deg, #5db7ee 50%, transparent 50%)
        !important;
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
          class:selected={localSelectedTestIDs.includes(row[fields.indexOf("id")])}
          class:top-selected={
            localSelectedTestIDs.includes(row[fields.indexOf("id")]) &&
            selectedTestIdxes.includes(i - 1)
          }
          class:bottom-selected={
            localSelectedTestIDs.includes(row[fields.indexOf("id")]) &&
            selectedTestIdxes.includes(i + 1)
          }
          class:top-and-bottom-selected={
            localSelectedTestIDs.includes(row[fields.indexOf("id")]) &&
            selectedTestIdxes.includes(i - 1) &&
            selectedTestIdxes.includes(i + 1)
          }
          on:click={() => toggleRow(row, i)}
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
