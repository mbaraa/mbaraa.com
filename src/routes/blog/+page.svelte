<script lang="ts">
	import Blogs from "$lib/components/blog/Blogs.svelte";
	import Intro from "$lib/components/blog/Intro.svelte";
	import type Blog from "$lib/models/Blog";
	import type Info from "$lib/models/Info";

	export let data: { blogs: Blog[]; info: Info };
</script>

<svelte:head>
	<title>Blogs By Baraa</title>
</svelte:head>

<div class="font-[Vistol]">
	{#if data}
		<Intro
				name={data.info.name.substring(0, data.info.name.indexOf(" "))}
				blogIntro={data.info.blogIntro}
		/>
		<div class="md:relative md:top-[-100px]">
			<Blogs blogs={data.blogs} />
		</div>
	{/if}
</div>
