<script lang="ts">
	import Header from "$lib/components/Header.svelte";
	import Bio from "$lib/components/home/Bio.svelte";
	import type Link from "$lib/models/Link";
	import type Info from "$lib/models/Info";

	export let data: { links: Link[]; info: Info };
</script>

<svelte:head>
	<title>Home of Baraa</title>
</svelte:head>

<div class="">
	<Header title={data.info.name} showTitle={false} showSeparator={false} links={data.links} />
	<Bio name={data.info.name} brief={data.info.brief} />
</div>
