<script>

  // Your selected Skeleton theme:
  import '../theme.postcss';

  // This contains the bulk of Skeletons required styles:
  import '@skeletonlabs/skeleton/styles/all.css';

  // Finally, your application's global stylesheet (sometimes labeled 'app.css')
  import '../app.postcss';

  import { AppShell, AppBar } from '@skeletonlabs/skeleton';

  import { version } from '$app/environment';

</script>


<AppShell regionPage="relative" slotPageHeader="sticky top-0 z-10" class="p-8 pr-10 bg-secondary-50">

  <svelte:fragment slot="pageHeader">
    <AppBar background="bg-secondary-500">
      <svelte:fragment slot="lead">
        <img src="oarLogo.png" alt="OAR Logo" class="h-10 w-10" />
      </svelte:fragment>
      OAR Enrich v{version}
    </AppBar>
  </svelte:fragment>

  <slot />

</AppShell>
