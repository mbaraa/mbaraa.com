<script lang="ts">
	import Bio from "$lib/components/home/Bio.svelte";
	import type Info from "$lib/models/Info";

	export let data: { info: Info };
</script>

<svelte:head>
	<title>Home of Baraa</title>
</svelte:head>

<Bio name={data.info.name} brief={data.info.brief} />
