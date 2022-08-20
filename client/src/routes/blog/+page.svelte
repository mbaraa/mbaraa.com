<script lang="ts">
    import Blogs from "$lib/components/Blogs.svelte";
    import BlogRequests from "$lib/utils/requests/BlogRequests";

    import { onMount } from "svelte";
    let blogs: any = [];

    onMount(async () => {
        blogs = await BlogRequests.getBlogs();
    });
</script>

<svelte:head>
    <title>Made By Baraa - Blogs</title>
</svelte:head>

<div class="bg-[#2d333b] bg-[#2d333b] w-full h-max ">
    <Blogs {blogs} />
</div>
