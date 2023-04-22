<script lang="ts">
	import type Link from "$lib/models/Link";
	import Seperator from "$lib/ui/Seperator.svelte";

	export let showTitle: boolean;
	export let showSeparator: boolean;
	export let links: Link[];
	export let title: string;

	function formatNumber(num: number): string {
		return num < 10 ? `0${num}` : `${num}`;
	}
</script>

<header
	class="w-[100vw] bg-black flex justify-center sm:justify-between text-white font-[Vistol] p-[30px] text-[20px]"
>
	<a class="font-[1000] hidden sm:block text-[30px] hover:text-[#20db8f]" href="/">
		{#if showTitle}
			{title}
		{/if}
	</a>

	<nav>
		<ul class="m-0 list-none flex gap-[20px]">
			{#each links as link, i}
				<li>
					<a
						class="underline hover:text-[#20DB8F]"
						target={link.target ?? "_blank"}
						href={link.link}
						><span class="text-[#20DB8F]">{formatNumber(i + 1)}.&nbsp;</span>{link.name}</a
					>
				</li>
			{/each}
		</ul>
	</nav>
</header>

{#if showSeparator}
	<div class="relative top-[-10px]">
		<Seperator />
	</div>
{/if}
