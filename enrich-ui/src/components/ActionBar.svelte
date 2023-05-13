<script>
import { version } from "$app/environment";
import { AppBar } from "@skeletonlabs/skeleton";
import Icon from "./Icon.svelte";
</script>

<AppBar background="bg-secondary-500" class="rounded-md" shadow="shadow-md">
  <svelte:fragment slot="lead">
    <img src="oarLogo.png" alt="OAR Logo" class="h-10 w-10" />
    <div class="pl-2 bg-surface-500 text-transparent bg-clip-text">
      <h1>OAR Enrich v{version}</h1>
    </div>
  </svelte:fragment>
  <svelte:fragment slot="trail">
    <button type="button" class="btn bg-primary-300 border-r-2 border-primary-700 border-t-2 active:border-t-0 active:border-r-0"><Icon name="edit"/></button>
    <button type="button" class="btn bg-error-300 border-r-2 border-error-700 border-t-2 active:border-t-0 active:border-r-0"><Icon name="trash"/></button>
  </svelte:fragment>
</AppBar>

<style>
  h1 {
      font-family: "pixelated",monospace;
      font-size: 18px;
  }
</style>
