<script lang="ts">
    import Header from "$lib/components/Header.svelte";

    import NewEditBlog from "$lib/components/NewEditBlog.svelte";
    import type Blog from "$lib/models/Blog";

    let blog: Blog = {
        name: "",
        description: "",
        content: "",
        createdAt: new Date(),
        updatedAt: new Date(),
    };

    let scrollY = 420;
</script>

<svelte:head>
    <title>Made By Baraa</title>
</svelte:head>

<Header {scrollY} />
<div class="pb-[85px] block" />
<div class="font-[Vistol]">
    <NewEditBlog bind:blog />
</div>
