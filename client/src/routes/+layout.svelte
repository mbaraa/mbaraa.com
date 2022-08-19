<script lang="ts">
    import "../app.css";
    import Header from "$lib/components/Header.svelte";
    import Name from "$lib/components/Name.svelte";
    import Separator from "$lib/components/Separator.svelte";
</script>

<svelte:head>
    <title>Made By Baraa</title>
</svelte:head>

<div class="font-[SFUI] bg-[#2d333b] h-[100%] ">
    <Header />
    <Separator />
    <Name />
    <Separator />
    <slot />
</div>
