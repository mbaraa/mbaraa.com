<script lang="ts">
	import type Volunteering from "$lib/models/Volunteering";
	import { onMount } from "svelte";
	import Button from "$lib/ui/Button.svelte";
	import VolunteeringEntry from "$lib/components/dashboard/VolunteeringXPEntry.svelte";

	let volunteerings: Volunteering[] = [];

	function addVolunteering() {
		volunteerings.push({roles: []})
        volunteerings = volunteerings
    }

	onMount(async () => {
		volunteerings = await fetch("/api/volunteering", {
			method: "GET",
        })
            .then((resp) => resp.json())
            .then((volunteerings) => volunteerings)
            .catch(() => []);
	});
</script>

<svelte:head>
    <title>Dashboard - mbaraa.com</title>
</svelte:head>

<div class="text-white">
    <div class="flex justify-between">
        <h1 class="text-[20px] font-bold">Volunteerings</h1>
        <Button on:click={addVolunteering} title="Add Volunteering"/>
    </div>
    <div class="mt-[10px]">
        {#each volunteerings as volunteering}
            <VolunteeringEntry {volunteering}/>
        {/each}
    </div>
</div>