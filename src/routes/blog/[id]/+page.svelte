<script lang="ts">
    import { page } from "$app/stores";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import { onMount } from "svelte";
    import { default as BlogV } from "$lib/components/Blog.svelte";
    import type Blog from "$lib/models/Blog";
    import { marked } from "marked";

    let blog: Blog;

    onMount(async () => {
        blog = await BlogRequests.getBlog($page.params.id);
    });
</script>

<svelte:head>
    <title>
        {blog?.name}
    </title>
    <meta property="og:title" content={blog?.name} />
</svelte:head>

{#if blog}
    <div
        class="font-[Vistol] absolute left-[50%] translate-x-[-50%] py-[50px] w-auto "
    >
        <h1
            class="text-white w-[100%] text-center text-[30px] md:text-[40px] font-[1000] "
        >
            {blog.name}
        </h1>
        <div
            class="p-[45px] text-[20px] m-[20px] bg-white rounded-[32px] w-[90vw] "
        >
            {@html marked.parse(blog.content)}
        </div>
    </div>
{/if}
