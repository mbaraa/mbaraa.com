<script lang="ts">
	import type Info from "$lib/models/Info";

	export let data: { info: Info };
</script>

<svelte:head>
	<title>About Baraa</title>
</svelte:head>

<div class="font-[Vistol]">
	{#if data}
		<div
			class="h-[90vh] sm:h-auto my-[30px] sm:my-[70px] py-[70px] sm:py-[100px] w-[90%] sm:w-[93%] xl:w-[70%] absolute left-[50%] translate-x-[-50%] top-[50%] translate-y-[-50%] z-[-1]"
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
	{/if}
</div>
