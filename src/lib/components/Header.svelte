<script lang="ts">
	import type Link from "$lib/models/Link";
	import Seperator from "$lib/ui/Seperator.svelte";
	import { default as LinkView } from "$lib/ui/Link.svelte";

	export let showTitle: boolean;
	export let showSeparator: boolean;
	export let links = new Array<Link>();
	export let title: string;
</script>

<div class="fixed z-[50]">
	<header
		class="w-[100vw] bg-black flex justify-center sm:justify-between text-white font-[Vistol] p-[40px] text-[20px]"
	>
		<a class="font-[1000] hidden sm:block text-[30px] hover:text-[#20db8f]" href="/">
			{#if showTitle}
				{title}
			{/if}
		</a>

		<nav>
			<ul class="m-0 list-none flex gap-[20px]">
				{#each links as link}
					<li>
						<LinkView {link} />
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
</div>
