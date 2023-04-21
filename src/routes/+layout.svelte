<script lang="ts">
	import "../app.css";
	import Header from "$lib/components/Header.svelte";
	import type Link from "$lib/models/Link";
	import type Info from "$lib/models/Info";

	export let data: { links: Link[]; info: Info };
</script>

<Header title={data.info.name} showTitle={false} showSeparator={false} links={data.links} />
<slot />
