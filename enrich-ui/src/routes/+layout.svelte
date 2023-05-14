<script>

  // Your selected Skeleton theme:
  import "../theme.postcss";

  // This contains the bulk of Skeletons required styles:
  import "@skeletonlabs/skeleton/styles/all.css";

  // Finally, your application's global stylesheet (sometimes labeled 'app.css')
  import "../app.postcss";

  import { AppShell, Modal, storeHighlightJs, Toast } from "@skeletonlabs/skeleton";
  import ActionBar from "../components/ActionBar.svelte";
  import hljs from "highlight.js";
  import "highlight.js/styles/github-dark.css";

  storeHighlightJs.set(hljs);
</script>

<Modal />
<Toast />
<AppShell regionPage="relative p-8 pr-10 bg-tertiary-50" slotPageHeader="sticky top-0 z-10">

  <svelte:fragment slot="pageHeader">
    <ActionBar />
  </svelte:fragment>

  <slot />

</AppShell>
