<script lang="ts">
	import type Info from "$lib/models/Info";

	export let data: { info: Info };
</script>

<svelte:head>
	<title>About Baraa</title>
</svelte:head>

<div class="font-[Vistol]">
	<div
		class="w-[80%] md:w-[50%] absolute left-[50%] translate-x-[-50%] top-[50%] translate-y-[-50%]"
	>
		<h1 class="text-[25px] font-bold text-white">About Me</h1>
		<p class="text-[#20DB8F] text-[20px]">{data.info.about}</p>
		<h2 class="text-[#ABABAB] text-[17px] mt-[10px]">Some of the technologies that I've used:</h2>
		<ul class="list-disc">
			{#each data.info.technologies as technology}
				<li class="ml-[15px] text-white">{technology}</li>
			{/each}
		</ul>
	</div>
</div>
