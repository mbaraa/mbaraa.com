<script lang="ts">
    import BlogIntro from "$lib/components/BlogIntro.svelte";
    import Blogs from "$lib/components/Blogs.svelte";
    import type Blog from "$lib/models/Blog";
    import type {Load} from "@sveltejs/kit";

    export let data: Load;
    const blogs: Blog[] = data.blogs;
</script>

<svelte:head>
    <title>Blogs By Baraa</title>
</svelte:head>

<BlogIntro />
<Blogs {blogs} />
