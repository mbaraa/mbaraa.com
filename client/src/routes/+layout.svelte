<script lang="ts">
    import Header from "$lib/components/Header.svelte";

    import "../app.css";

    let scrollY = 0;
</script>

<Header {scrollY} />
<div class="pb-[85px] block" />
<slot />

<svelte:window bind:scrollY />
