<script lang="ts">
    import { page } from "$app/stores";
    import Achievements from "$lib/components/Achievements.svelte";
    import Header from "$lib/components/Header.svelte";
    import Name from "$lib/components/Name.svelte";
    import Separator from "$lib/components/Separator.svelte";
</script>

<Header />
<Separator />
<Name />
<Separator />
<Achievements />
