<script lang="ts">
    import type Work from "$lib/models/Work";
	import { onMount } from "svelte";
	import Button from "$lib/ui/Button.svelte";
	import WorkEntry from "$lib/components/dashboard/WorkXPEntry.svelte";

	let works: Work[] = [];

	function addWork() {
		works.push({roles: []})
        works = works
    }

	onMount(async () => {
		works = await fetch("/api/work", {
			method: "GET",
        })
            .then((resp) => resp.json())
            .then((works) => works)
            .catch(() => []);
	});
</script>

<svelte:head>
    <title>Dashboard - mbaraa.com</title>
</svelte:head>

<div class="text-white">
    <div class="flex justify-between">
        <h1 class="text-[20px] font-bold">Works</h1>
        <Button on:click={addWork} title="Add Work"/>
    </div>
    <div class="mt-[10px]">
        {#each works as work}
            <WorkEntry {work}/>
        {/each}
    </div>
</div>