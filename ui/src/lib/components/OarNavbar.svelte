<script lang="ts">
    import {Navbar, NavBrand} from 'flowbite-svelte'
</script>

<Navbar rounded color="form" let:hidden let:toggle>
    <NavBrand href="/">
        <img src="oarLogo.png" class="sm:h-12" alt="OAR Logo"/>
    </NavBrand>
</Navbar>