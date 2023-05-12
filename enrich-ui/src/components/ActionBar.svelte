<script>
import { version } from "$app/environment";
import { AppBar } from "@skeletonlabs/skeleton";
import Icon from "./Icon.svelte";
</script>

<AppBar background="bg-secondary-500" class="rounded-md" shadow="shadow-md">
  <svelte:fragment slot="lead">
    <img src="oarLogo.png" alt="OAR Logo" class="h-10 w-10" />
    <div class="pl-2 ">OAR Enrich v{version}</div>
  </svelte:fragment>
  <svelte:fragment slot="trail">
    <button type="button" class="btn bg-primary-300"><Icon name="edit"/></button>
    <button type="button" class="btn bg-error-300"><Icon name="trash"/></button>
  </svelte:fragment>
</AppBar>
