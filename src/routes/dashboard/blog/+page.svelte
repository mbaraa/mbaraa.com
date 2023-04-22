<script lang="ts">
    import type Blog from "$lib/models/Blog";
	import { onMount } from "svelte";
	import Button from "$lib/ui/Button.svelte";
	import BlogEntry from "../../../lib/components/dashboard/BlogEntry.svelte";

	let blogs: Blog[] = [];

	function addBlog() {
		blogs.push({})
        blogs = blogs
    }

	onMount(async () => {
		blogs = await fetch("/api/blog", {
			method: "GET",
        })
            .then((resp) => resp.json())
            .then((blogs) => blogs)
            .catch(() => []);

		for (const blog: Blog of blogs) {
			blog.content = await fetch(`/api/blog?id=${blog.publicId}`, {
				method: "GET",
			})
				.then((resp) => resp.json())
				.then((blog) => blog.content)
				.catch(() => "")
        }
	});
</script>

<svelte:head>
    <title>Dashboard - mbaraa.com</title>
</svelte:head>

<div class="text-white">
    <div class="flex justify-between">
        <h1 class="text-[20px] font-bold">Blogs</h1>
        <Button on:click={addBlog} title="Add Blog"/>
    </div>
    <div class="mt-[10px]">
        {#each blogs as blog}
            <BlogEntry {blog}/>
        {/each}
    </div>
</div>