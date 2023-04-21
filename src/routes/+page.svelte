<script lang="ts">
	import Header from "$lib/components/Header.svelte";
	//	import Bio from "$lib/components/Bio.svelte";
	//	import Projects from "$lib/components/Projects.svelte";
	import Seperator from "$lib/ui/Seperator.svelte";
	import type ProjectGroup from "$lib/models/ProjectGroup";
	import type Link from "$lib/models/Link";
	import type Info from "$lib/models/Info";

	export let data: { groups: ProjectGroup[]; links: Link[]; info: Info };

	$: showTitle = true; //scrollY > 120;
	$: showSeparator = true; //scrollY > 320;
	let scrollY = 0;
</script>

<svelte:head>
	<title>Made By Baraa</title>
</svelte:head>

<Header title={data.info.name} {showTitle} {showSeparator} links={data.links} />
<div class="pb-[85px] block" />
<div class="font-[Vistol]">
	<!-- <Bio name={data.info.name} about={data.info.about}/> -->
	<Seperator />
	<!-- <Projects {data.groups} /> -->
</div>

<svelte:window bind:scrollY />
