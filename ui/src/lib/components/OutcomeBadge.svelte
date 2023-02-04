<script lang="ts">
    import {Badge} from 'flowbite-svelte'
    import {Outcome} from "$lib/consts";

    export let outcome: Outcome
    let badgeColor = (() => {
        switch(outcome){
            case Outcome.Passed: return "green"
            case Outcome.Failed: return "red"
            default: return "Invalid"
        }
    })()
</script>

<Badge large color={badgeColor}>{outcome}</Badge>
