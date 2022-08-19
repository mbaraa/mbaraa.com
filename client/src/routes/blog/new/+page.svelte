<script lang="ts">
    import { default as BlogV } from "$lib/components/Blog.svelte";
    import type Blog from "$lib/models/Blog";

    let blog: Blog;
</script>

<BlogV {blog} isNew={true} />
