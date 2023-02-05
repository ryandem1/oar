<script lang="ts">
    import {Input, Navbar, NavBrand} from 'flowbite-svelte'
    import AnalysisDropdown from "$lib/components/AnalysisDropdown.svelte";
    import ResolutionDropdown from "$lib/components/ResolutionDropdown.svelte";
</script>

<Navbar rounded color="form" let:hidden let:toggle>
    <NavBrand href="/">
        <img src="oarLogo.png" class="sm:h-12" alt="OAR Logo"/>
    </NavBrand>
    <Input class="w-96" type="text" id="json_path_query" placeholder="JSON Path Query..." />
    <AnalysisDropdown/>
    <ResolutionDropdown/>
    <NavBrand href="/">
        <img src="oarLogo.png" class="sm:h-12" alt="OAR Logo"/>
    </NavBrand>
</Navbar>
