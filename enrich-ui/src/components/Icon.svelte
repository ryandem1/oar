<script>
  import feather from "feather-icons";
  export const directions = ["n", "ne", "e", "se", "s", "sw", "w", "nw"];

  export let name;
  export let direction = "n";
  export let width = "1em";
  export let height = "1em";

  if (name === undefined) {
    name = "default"
  }
  $: icon = feather.icons[name];
  $: rotation = directions.indexOf(direction) * 45;
</script>
<!-- eslint-disable -->
{#if icon}
  <svg
    {...icon.attrs}
    style="width: {width}; height: {height}; transform: rotate({rotation}deg);"
  >
    <g>
      {@html icon.contents}
    </g>
  </svg>
{/if}

<style>
    svg {
        width: 1em;
        height: 1em;
        overflow: visible;
        transform-origin: 50% 50%;
    }
</style>
