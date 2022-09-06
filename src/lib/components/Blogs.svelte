<script lang="ts">
    import type Blog from "$lib/models/Blog";
    import { default as BlogView } from "./Blog.svelte";

    export let blogs = new Array<Blog>();
</script>

<div class="w-auto absolute top-[420px] left-[50%] translate-x-[-50%]">
    {#each blogs as blog}
        <BlogView {blog} />
    {/each}
</div>
