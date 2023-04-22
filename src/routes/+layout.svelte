<script lang="ts">
	import "../app.css";
	import Header from "$lib/components/Header.svelte";
	import type Link from "$lib/models/Link";
	import { data } from "autoprefixer";

	let links: Link[] = [
		{ name: "Blog", link: "/blog", target: "" },
		{ name: "Projects", link: "/projects", target: "" },
		{ name: "Work", link: "/work", target: "" },
		{ name: "About", link: "/about", target: "" }
	];
</script>

{#if data}
<Header title="mbaraa.com" showTitle showSeparator={false} {links} />
{/if}
<slot />
