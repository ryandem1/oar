<script lang="ts">
  import { Paginator } from '@skeletonlabs/skeleton';

  let page = {
    offset: 0,
    limit: 5,
    size: 50,
    amounts: [1, 2, 5, 10],
  };

  function onPageChange(e: CustomEvent): void {
    console.log('event:page', e.detail);
  }

  function onAmountChange(e: CustomEvent): void {
    console.log('event:amount', e.detail);
  }
</script>


<Paginator on:page={onPageChange} on:amount={onAmountChange}>

</Paginator>
