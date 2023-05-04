<script>

  // Your selected Skeleton theme:
  import '../theme.postcss';

  // This contains the bulk of Skeletons required styles:
  import '@skeletonlabs/skeleton/styles/all.css';

  // Finally, your application's global stylesheet (sometimes labeled 'app.css')
  import '../app.postcss';

  import { AppShell, AppBar } from '@skeletonlabs/skeleton';

</script>


<AppShell regionPage="relative" slotPageHeader="sticky top-0 z-10">
  <svelte:fragment slot="pageHeader">
    <AppBar background="bg-surface-300">Something else</AppBar>
  </svelte:fragment>

  <slot />
</AppShell>
