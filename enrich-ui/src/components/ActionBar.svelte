<script lang="ts">
import { version } from "$app/environment";
import { AppBar, modalStore, toastStore } from "@skeletonlabs/skeleton";
import type { ModalSettings, ToastSettings } from '@skeletonlabs/skeleton';
import Icon from "./Icon.svelte";
import { selectedTestIDs, refreshTestTable } from "../stores.ts";
import { OARServiceClient } from "$lib/client";

const client = new OARServiceClient();

const onDeleteButtonClick = (): void => {
  let localSelectedTestIDs: number[];
  const unsubscribe = selectedTestIDs.subscribe(ids => {
    localSelectedTestIDs = ids;
  });
  unsubscribe();

  let numSelectedTests = localSelectedTestIDs.length;

  if (numSelectedTests === 0) {
    const t: ToastSettings = {
      message: "No selected tests to delete!",
      timeout: 3000,
      background: "bg-surface-400"
    };
    toastStore.trigger(t);
    return
  }

  const deleteConfirmationModal: ModalSettings = {
    type: "confirm",
    title: `Are you sure you would like to delete ${numSelectedTests} tests?`,
    body: "WARNING, this is not reversible",
    response: async (r: boolean) => {
      if (r) {
        await client.deleteTests({ids: localSelectedTestIDs});
        refreshTestTable.set(true);

        const t: ToastSettings = {
          message: "Tests deleted successfully",
          timeout: 3000,
          background: "bg-success-400"
        };
        toastStore.trigger(t);
      }
    },
  };

  modalStore.trigger(deleteConfirmationModal);
}

</script>

<AppBar background="bg-secondary-500" class="rounded-md" shadow="shadow-md">
  <svelte:fragment slot="lead">
    <img src="oarLogo.png" alt="OAR Logo" class="h-10 w-10" />
    <div class="pl-2 bg-surface-500 text-transparent bg-clip-text">
      <h1>OAR Enrich v{version}</h1>
    </div>
  </svelte:fragment>
  <svelte:fragment slot="trail">
    <button type="button" class="btn bg-primary-400 border-r-2 border-primary-500 border-t-2 active:border-t-0 active:border-r-0">
      <Icon name="eye"/>
    </button>
    <button type="button" class="btn bg-success-500 border-r-2 border-success-600 border-t-2 active:border-t-0 active:border-r-0">
      <Icon name="edit"/>
    </button>
    <button
      type="button"
      on:click={onDeleteButtonClick}
      class="btn bg-tertiary-500 border-r-2 border-tertiary-600 border-t-2 active:border-t-0 active:border-r-0"
    >
      <Icon name="trash"/>
    </button>
  </svelte:fragment>
</AppBar>

<style>
  h1 {
      font-family: "pixelated",monospace;
      font-size: 18px;
  }
</style>
