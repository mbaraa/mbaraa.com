<script lang="ts">
	import MarkdownViewer from "$lib/components/MarkdownViewer.svelte";
	import type Info from "$lib/models/Info";

	export let data: { info: Info };
	let techSizePerChunk = Math.ceil(data.info.technologies.length / 3);
	const bigScreenTechLists = new Array<Array<string>>();
	for (let i = 0; i < data.info.technologies.length; i += techSizePerChunk) {
		const chunk = new Array<string>();
		for (let j = i; j < i + techSizePerChunk; j++) {
			chunk.push(data.info.technologies[j]);
		}
		bigScreenTechLists.push(chunk);
	}
</script>

<svelte:head>
	<title>About Baraa</title>
</svelte:head>

<div class="font-[Vistol]">
	{#if data}
		<div class="h-[90vh] sm:h-auto">
			<div class="px-10 py-5 lg:p-20">
				<h1 class="text-[25px] font-bold text-white">About Me</h1>
				<p class="text-[#20DB8F] text-[20px]"><MarkdownViewer text={data.info.about} /></p>
				<hr class="my-10" />
				<h2 class="text-[#ABABAB] text-[17px] mb-3">Some of the technologies that I've used:</h2>
				<ul class="list-disc block lg:hidden">
					{#each data.info.technologies as technology}
						<li class="ml-[15px] text-white">{technology}</li>
					{/each}
				</ul>
				<div class="hidden lg:flex gap-32">
					{#each bigScreenTechLists as technologyChunk}
						<ul class="list-disc inline">
							{#each technologyChunk as technology}
								<li class="ml-[15px] text-white">{technology}</li>
							{/each}
						</ul>
					{/each}
				</div>
			</div>
		</div>
	{/if}
</div>
