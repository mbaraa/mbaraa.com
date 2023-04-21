<script lang="ts">
	import type Experience from "$lib/models/Experience";

	export let xps: Experience[];
	export let xpName: string;
</script>

<div class="text-white mb-[50px] mx-[30px]">
	<h1 class="font-bold text-[30px] mt-[20px] mb-[30px]">{xpName}:</h1>
	<ul class="text-[#20DB8F] text-[25px]">
		{#each xps as xp}
			<li class="relative flex gap-[20px]">
				<div class="font-[Space]">
					<h3 class="relative top-[-6px] text-[15px] text-center leading-[15px]">
						{xp.startDate}<br />-<br />{#if xp.endDate}{xp.endDate}{:else}....{/if}
					</h3>
				</div>
				<div class="border-l-[5px] border-l-[white] min-h-auto py-[20px]">{""}</div>
				<div class="text-[15px] pb-[20px] text-white">
					<h1 class="text-[25px]">{xp.name}</h1>
					{#if xp.description}
						<h3>{xp.description}</h3>
					{/if}
					{#if xp.roles}
						<ul class="list-disc">
							{#each xp.roles as role}
								<li class="ml-[12px]">{role}</li>
							{/each}
						</ul>
					{/if}
				</div>
			</li>
		{/each}
	</ul>
</div>

<style>
	ul li::before {
		content: "";
		position: absolute;
		height: 20px;
		width: 20px;
		border-radius: 50%;
		background-color: #20db8f;
		top: 7px;
		left: 48.5px;
	}
</style>
