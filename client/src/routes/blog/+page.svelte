<script lang="ts">
    import Blogs from "$lib/components/Blogs.svelte";
    import BlogRequests from "$lib/utils/requests/BlogRequests";

    import { onMount } from "svelte";
    let blogs: any = [];

    onMount(async () => {
        blogs = await BlogRequests.getBlogs();
    });
</script>

<div class="bg-[#2d333b] bg-[#2d333b] w-full h-full ">
    <Blogs {blogs} />
</div>
