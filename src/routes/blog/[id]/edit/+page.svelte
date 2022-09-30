<script lang="ts">
    import { page } from "$app/stores";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import { onMount } from "svelte";
    import type Blog from "$lib/models/Blog";
    import NewEditBlog from "$lib/components/NewEditBlog.svelte";

    let blog: Blog;

    onMount(async () => {
        blog = await BlogRequests.getBlog($page.params.id);
    });
</script>

{#if blog}
    <NewEditBlog bind:blog />
{/if}
