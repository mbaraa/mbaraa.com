<script lang="ts">
	import type Experience from "$lib/models/Experience";
	import Experiences from "$lib/components/home/Experiences.svelte";

	export let data: {
		work: Experience[];
		volunteering: Experience[];
	};
</script>

<svelte:head>
	<title>Contributed by Baraa</title>
</svelte:head>

<div class="font-[Vistol]">
	<Experiences xpName="Professional Work" xps={data.work} />
	<Experiences xpName="Volunteering" xps={data.volunteering} />
</div>
