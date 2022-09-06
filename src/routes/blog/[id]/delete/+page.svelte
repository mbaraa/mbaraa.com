<script lang="ts">
    import { page } from "$app/stores";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import { onMount } from "svelte";

    onMount(async () => {
        await BlogRequests.deleteBlog($page.params.id);
    });
</script>
