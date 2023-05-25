<script lang="ts">
  export let parent: unknown;

  import { CodeBlock, modalStore } from "@skeletonlabs/skeleton";

  const cBase = 'card p-4 w-modal shadow-xl space-y-4';
  const cHeader = 'text-2xl font-bold';
</script>

<!-- @component This example creates a simple form modal. -->

{#if $modalStore[0]}
  <div class="modal-example-form max-h-screen overflow-y-scroll {cBase}">
    <header class={cHeader}>{$modalStore[0].title ?? '(title missing)'}</header>
    <article>

    </article>
    <footer class="modal-footer {parent.regionFooter}">
      <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>{parent.buttonTextCancel}</button>
    </footer>
  </div>
{/if}