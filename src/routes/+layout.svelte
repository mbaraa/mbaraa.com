<script lang="ts">
	import "../app.css";
	import Header from "$lib/components/Header.svelte";
	import type Link from "$lib/models/Link";

	let links: Link[] = [
		{ name: "Blog", link: "/blog", target: "" },
		{ name: "Work", link: "/work", target: "" },
		{ name: "About", link: "/about", target: "" },
		{ name: "Contact", link: "/contact", target: "" }
	];
</script>

<Header title="mbaraa.com" showTitle showSeparator={false} {links} />
<slot />
