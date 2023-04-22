<script lang="ts">
	import type Link from "$lib/models/Link";

	export let link: Link;
</script>

<a class="underline hover:text-[#20db8f]" target={link.target ?? "_blank"} href={link.link}
	>{link.name}</a
>
