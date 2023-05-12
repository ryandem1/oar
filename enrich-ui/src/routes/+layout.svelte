<script>

  // Your selected Skeleton theme:
  import "../theme.postcss";

  // This contains the bulk of Skeletons required styles:
  import "@skeletonlabs/skeleton/styles/all.css";

  // Finally, your application's global stylesheet (sometimes labeled 'app.css')
  import "../app.postcss";

  import { AppBar, AppShell } from "@skeletonlabs/skeleton";
  import { version } from "$app/environment";
</script>


<AppShell regionPage="relative p-8 pr-10 bg-tertiary-50" slotPageHeader="sticky top-0 z-10">

  <svelte:fragment slot="pageHeader">
    <AppBar background="bg-secondary-500" class="rounded-md" shadow="shadow-md">
      <svelte:fragment slot="lead">
        <img src="oarLogo.png" alt="OAR Logo" class="h-10 w-10" />
        <div class="pl-2 ">OAR Enrich v{version}</div>
      </svelte:fragment>
      <svelte:fragment slot="trail">
        <button type="button" class="btn bg-primary-300">Enrich</button>
        <button type="button" class="btn bg-error-300">Delete</button>
      </svelte:fragment>
    </AppBar>
  </svelte:fragment>

  <slot />

</AppShell>
