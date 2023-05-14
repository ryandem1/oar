<script lang="ts">
  export let parent: any;

  import { CodeBlock, modalStore } from "@skeletonlabs/skeleton";
  import { OARServiceClient } from "$lib/client";
  import { onMount } from "svelte";
  import { getSelectedTestIDs } from "$lib/table";
  import type { Test } from "$lib/models";
  import { isEnrichUIError, isOARServiceError } from "$lib/models";
  import { throwFailureToast } from "$lib/toasts";

  const client = new OARServiceClient();

  let tests: Test[];
  $: tests = [];
  let stringifiedTests: string;
  $: stringifiedTests = tests.map((test) => JSON.stringify(test, null, 2)).join(",\n");

  onMount(async () => {
    const testIDs = getSelectedTestIDs();
    const result = await client.getTests({ids: testIDs});

    if (isEnrichUIError(result) || isOARServiceError(result)) {
      throwFailureToast("Could not get test details");
      return
    }

    tests = result.tests;
  })

  const cBase = 'card p-4 w-modal shadow-xl space-y-4';
  const cHeader = 'text-2xl font-bold';
</script>

<!-- @component This example creates a simple form modal. -->

{#if $modalStore[0]}
  <div class="modal-example-form h-screen overflow-y-scroll {cBase}">
    <header class={cHeader}>{$modalStore[0].title ?? '(title missing)'}</header>
    <article>
      <CodeBlock language="JSON" code={`${stringifiedTests}`}/>
    </article>
    <footer class="modal-footer {parent.regionFooter}">
      <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>{parent.buttonTextCancel}</button>
    </footer>
  </div>
{/if}