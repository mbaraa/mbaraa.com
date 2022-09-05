<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import Center from "$lib/components/Center.svelte";

    let scrollY = 0;
</script>

<svelte:head>
    <title>Made By Baraa</title>
</svelte:head>

<Header {scrollY} />
<div class="pb-[85px] block" />
<div class="font-[Vistol]">
    <Center />
</div>

<svelte:window bind:scrollY />
