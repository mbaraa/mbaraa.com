<script lang="ts">
	import type ProjectGroup from "$lib/models/ProjectGroup";
	import ProjectGroups from "./ProjectGroups.svelte";

	export let groups: ProjectGroup[];
</script>

<div class="pb-[50px]">
	{#each groups as group}
		<div class="">
			<div
				class="mx-[10%] my-[20px] md:mt-[50px] md:mb-[70px] bg-white p-[10px] rounded-tr-[15px] rounded-bl-[15px] text-center"
			>
				<h1 class="text-[25px] font-bold">{group.name}</h1>
				<h2 class="text-[15px]">{group.description}</h2>
			</div>
			<ProjectGroups projects={group.projects} />
		</div>
	{/each}
</div>
