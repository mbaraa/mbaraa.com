<script lang="ts">
	import MarkdownViewer from "$lib/components/MarkdownViewer.svelte";
	import type Blog from "$lib/models/Blog";
	export let data: { blog: Blog };
	const blog = data.blog;
</script>

<svelte:head>
	<title>
		{blog?.name}
	</title>
	<meta property="og:title" content={blog?.name} />
</svelte:head>

{#if blog}
	<div class="font-[Vistol] absolute left-[50%] translate-x-[-50%] py-[50px] w-auto">
		<h1 class="text-white w-[100%] text-center text-[30px] md:text-[40px] font-[1000]">
			{blog.name}
		</h1>
		<div class="p-[45px] text-[18px] m-[18px] bg-white rounded-[32px] w-[90vw]">
			<MarkdownViewer text={blog.content} />
		</div>
	</div>
{/if}
