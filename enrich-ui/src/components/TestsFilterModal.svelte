<script lang="ts">
  import { refreshTestTable, selectedTestIDs, testTableFields, testTableQuery } from "../stores";
  import { InputChip, modalStore } from "@skeletonlabs/skeleton";
  import { OARServiceClient } from "$lib/client";
  import type { TestQuery } from "$lib/models";
  import { JSONEditor } from "svelte-jsoneditor";
  import { getTestTableFields } from "$lib/table";
  import { throwFailureToast } from "$lib/toasts";

  interface ParentModal {
    onClose(): null
    regionFooter: unknown
    buttonNeutral: unknown
    buttonTextCancel: unknown
    buttonPositive: unknown
    buttonTextSubmit: unknown
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
  const cHeader = 'text-2xl font-bold p-4';

  const onSubmit = async () => {
    const query: TestQuery = JSON.parse(content.text)

    if (!localFields.includes("id")) {
      throwFailureToast("Must include the id field!");
      return
    }
    testTableQuery.set(query);
    testTableFields.set(localFields);
    selectedTestIDs.set([]);
    parent.onClose()
    refreshTestTable.set(true);
  }

  let localFields: string[];
  $: localFields = getTestTableFields();
</script>

<!-- @component This example creates a simple form modal. -->

{#if $modalStore[0]}
  <div class="modal-example-form max-h-screen overflow-y-scroll {cBase}">
    <header class={cHeader}>{$modalStore[0].title ?? '(title missing)'}</header>
    <article class="my-json-editor">
      <JSONEditor bind:content {options} mode="text" mainMenuBar={false} navigationBar={false}/>
      <header class={cHeader}>Table Fields</header>
      <InputChip bind:value={localFields} allowUpperCase name="chips" placeholder="Enter any value..." />
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
