<script lang="ts">
	import Projects from "$lib/components/home/Projects.svelte";
	import type ProjectGroup from "$lib/models/ProjectGroup";

	export let data: { groups: ProjectGroup[] };
</script>

<svelte:head>
	<title>Made by Baraa</title>
</svelte:head>

<div class="font-[Vistol]">
	{#if data}
		<Projects groups={data.groups} />
	{/if}
</div>
