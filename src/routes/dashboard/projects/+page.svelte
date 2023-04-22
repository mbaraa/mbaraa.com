<script lang="ts">
	import type ProjectGroup from "$lib/models/ProjectGroup";
	import { onMount } from "svelte";
	import Button from "$lib/ui/Button.svelte";
	import ProjectGroupEntry from "$lib/components/dashboard/ProjectGroupEntry.svelte";

	let projectGroups: ProjectGroup[] = [];

	function addProjectGroup() {
		projectGroups.push({ description: "", name: "", projects: [], publicId: "" });
		projectGroups = projectGroups;
	}

	onMount(async () => {
		projectGroups = await fetch("/api/project", {
			method: "GET"
		})
			.then((resp) => resp.json())
			.then((projectGroups) => projectGroups)
			.catch(() => []);
	});
</script>

<svelte:head>
	<title>Dashboard - mbaraa.com</title>
</svelte:head>

<div class="text-white">
	<div class="flex justify-between">
		<h1 class="text-[20px] font-bold">Project Groups</h1>
		<Button on:click={addProjectGroup} title="Add ProjectGroup" />
	</div>
	<div class="mt-[10px]">
		{#each projectGroups as projectGroup}
			<ProjectGroupEntry group={projectGroup} />
		{/each}
	</div>
</div>
