<script lang="ts">
    import { page } from "$app/stores";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import { onMount } from "svelte";
    import { default as BlogV } from "$lib/components/Blog.svelte";
    import type Blog from "$lib/models/Blog";

    let blog: Blog;

    onMount(async () => {
        blog = await BlogRequests.getBlog($page.params.id);
    });
</script>

<svelte:head>
    <title>
        {blog?.name}
    </title>
</svelte:head>

{#if blog}
    <div class="bg-[#2d333b] w-full h-full">
        <BlogV {blog} />
    </div>
{/if}
