<script>
  import { AppBar } from '@skeletonlabs/skeleton';
  import { version } from '$app/environment';
</script>

<AppBar background="bg-secondary-500" class="rounded-md" shadow="shadow-md">
  <svelte:fragment slot="lead">
    <img src="oarLogo.png" alt="OAR Logo" class="h-10 w-10" />
    <div class="pl-2 ">OAR Enrich v{version}</div>
  </svelte:fragment>
</AppBar>
