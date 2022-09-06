<script lang="ts">
    import BlogIntro from "$lib/components/BlogIntro.svelte";
    import Blogs from "$lib/components/Blogs.svelte";
    import type Blog from "$lib/models/Blog";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import { onMount } from "svelte";

    let blogs = new Array<Blog>();

    onMount(async () => {
        blogs = await BlogRequests.getBlogs();
    });
</script>

<svelte:head>
    <title>Blogs By Baraa</title>
</svelte:head>

<BlogIntro />
<Blogs {blogs} />
