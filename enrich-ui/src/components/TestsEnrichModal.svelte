<script lang="ts">
  import { refreshTestTable } from "../stores";
  import { modalStore } from "@skeletonlabs/skeleton";
  import { OARServiceClient } from "$lib/client";
  import { getSelectedTestIDs } from "$lib/table";
  import type { Test } from "$lib/models";
  import { throwFailureToast, throwSuccessToast, throwWarningToast } from "$lib/toasts";
  import { JSONEditor } from "svelte-jsoneditor";

  interface ParentModal {
    onClose(): null
    regionFooter: any
    buttonNeutral: any
    buttonTextCancel: any
    buttonPositive: any
    buttonTextSubmit: any
  }
  export let parent: ParentModal;

  const client = new OARServiceClient();

  let content = {
    text: "{}"
  }

  const options = {
    mode: 'text'
  };
  const cBase = 'card p-4 w-modal shadow-xl space-y-4';
  const cHeader = 'text-2xl font-bold';

  const onSubmit = async () => {
    const details: Test = JSON.parse(content.text)
    const testIDs = getSelectedTestIDs();
    const statusCode = await client.enrichTests(details, {ids: testIDs})
    if (statusCode == 304) {
      throwWarningToast("No tests were changed!")
    } else if (statusCode == 200) {
      throwSuccessToast('Tests enriched successfully!');
    } else {
      throwFailureToast("Error occurred when enriching tests!")
    }
    parent.onClose()
    refreshTestTable.set(true);
  }
</script>

<!-- @component This example creates a simple form modal. -->

{#if $modalStore[0]}
  <div class="modal-example-form max-h-screen overflow-y-scroll {cBase}">
    <header class={cHeader}>{$modalStore[0].title ?? '(title missing)'}</header>
    <article class="my-json-editor">
      <JSONEditor bind:content {options} mode="text" mainMenuBar={false} navigationBar={false}/>
    </article>
    <footer class="modal-footer {parent.regionFooter}">
      <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>{parent.buttonTextCancel}</button>
      <button class="btn {parent.buttonPositive}" on:click={onSubmit}>{parent.buttonTextSubmit}</button>
    </footer>
  </div>
{/if}

<style>
    .my-json-editor {
        /* define a custom theme color */
        --jse-theme-color: #988a6c;
        --jse-theme-color-highlight: #5db7ee;
    }
</style>
