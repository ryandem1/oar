<script lang="ts">
  import { version } from "$app/environment";
  import { AppBar } from "@skeletonlabs/skeleton";
  import Icon from "./Icon.svelte";
  import { onDeleteButtonClick, onEnrichButtonClick, onFilterButtonClick, onViewButtonClick } from "$lib/actionBar";
</script>

<AppBar background="bg-secondary-500" class="rounded-md" shadow="shadow-md">
  <svelte:fragment slot="lead">
    <img src="oarLogo.png" alt="OAR Logo" class="h-10 w-10" />
    <div class="pl-2 bg-surface-500 text-transparent bg-clip-text">
      <h1>OAR Enrich v{version}</h1>
    </div>
  </svelte:fragment>
  <svelte:fragment slot="trail">
    <button
      type="button"
      on:click={onFilterButtonClick}
      class="btn bg-surface-400 border-r-2 border-surface-500 border-t-2 active:border-t-0 active:border-r-0"
    >
      <Icon name="filter"/>
    </button>
    <button
      type="button"
      on:click={onViewButtonClick}
      class="btn bg-primary-400 border-r-2 border-primary-500 border-t-2 active:border-t-0 active:border-r-0"
    >
      <Icon name="eye"/>
    </button>
    <button
      type="button"
      on:click={onEnrichButtonClick}
      class="btn bg-success-500 border-r-2 border-success-600 border-t-2 active:border-t-0 active:border-r-0"
    >
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
