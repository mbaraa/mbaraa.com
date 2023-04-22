<script lang="ts">
    import type ProjectGroup from "$lib/models/ProjectGroup";
	import { onMount } from "svelte";
	import Button from "$lib/ui/Button.svelte";
	import ProjectGroupEntry from "../../../lib/components/dashboard/ProjectGroupEntry.svelte";

	let projectGroups: ProjectGroup[] = [];

	function addProjectGroup() {
		projectGroups.push({})
        projectGroups = projectGroups
    }

	onMount(async () => {
		projectGroups = await fetch("/api/project", {
			method: "GET",
        })
            .then((resp) => resp.json())
            .then((projectGroups) => projectGroups)
            .catch(() => []);

		for (const projectGroup: ProjectGroup of projectGroups) {
			projectGroup.content = await fetch(`/api/project?id=${projectGroup.publicId}`, {
				method: "GET",
			})
				.then((resp) => resp.json())
				.then((projectGroup) => projectGroup.content)
				.catch(() => "")
        }
	});
</script>

<svelte:head>
    <title>Dashboard - mbaraa.com</title>
</svelte:head>

<div class="text-white">
    <div class="flex justify-between">
        <h1 class="text-[20px] font-bold">ProjectGroups</h1>
        <Button on:click={addProjectGroup} title="Add ProjectGroup"/>
    </div>
    <div class="mt-[10px]">
        {#each projectGroups as projectGroup}
            <ProjectGroupEntry {projectGroup}/>
        {/each}
    </div>
</div>