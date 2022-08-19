<script lang="ts">
    import { page } from "$app/stores";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import { onMount } from "svelte";
    import { default as BlogV } from "$lib/components/Blog.svelte";
    import type Blog from "$lib/models/Blog";

    let blog: Blog;
    let isEdit = true;

    onMount(async () => {
        blog = await BlogRequests.getBlog($page.params.id);
    });
</script>

{#if blog}
    <BlogV {blog} {isEdit} />
{/if}
