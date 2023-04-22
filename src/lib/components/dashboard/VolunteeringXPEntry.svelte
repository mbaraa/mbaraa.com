<script lang="ts">
	import type Volunteering from "$lib/models/Volunteering";
	import Button from "$lib/ui/Button.svelte";

	export let volunteering: Volunteering;
	let editMode = false;

	async function saveVolunteering(): Promise<void> {
		let method = "POST";
		if (volunteering.publicId) {
			method = "PUT";
        }
		await fetch("/api/volunteering", {
			method: method,
			mode: "cors",
			headers: {
				"Authorization": localStorage.getItem("token") ?? ""
			},
			body: JSON.stringify(volunteering)
		})
	}

	async function deleteVolunteering(): Promise<void> {
		await fetch(`/api/volunteering?id=${volunteering.publicId}`, {
			method: "DELETE",
			headers: {
				"Authorization": localStorage.getItem("token") ?? ""
			},
		})
	}
</script>

<div class="text-black bg-[#CBCBCB] block rounded-[10px] mb-[10px] last:mb-0 p-[15px] ">
    <div class="flex justify-between">
        <h3>{volunteering.name}</h3>
        <h3 class="block font-bold cursor-pointer" on:click={() => {editMode = !editMode}}>
            {#if editMode}{"ᐯ"}{:else}{"ᐳ"}{/if}
        </h3>
    </div>
    {#if editMode}
        <div class="block">
            <div
                    class="font-[Vistol] w-auto "
            >
                <div
                        class="p-[45px] text-[20px] m-[20px] bg-white rounded-[32px] "
                >
                    <table class="w-[100%]">
                        <tbody>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">Name:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={volunteering.name}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">Description:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={volunteering.description}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">Location:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={volunteering.location}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">Start Year:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={volunteering.startDate}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">End Year:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={volunteering.endDate}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">Roles:</h1>
                            </td>
                            <td>
                                <Button _class="bg-white" on:click={() => {volunteering.roles.push(""); volunteering = volunteering}}
                                        title="+"/>
                                {#each volunteering.roles as role}
                                    <div class="flex justify-between">
                                    <textarea
                                            bind:value={role}
                                            class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                                        <Button _class="" on:click={() => {volunteering.roles = volunteering.roles.filter((_role) => _role !== role)}}
                                                title="-"/>
                                    </div>
                                {/each}
                            </td>
                        </tr>
                        </tbody>
                    </table>
                    <div class="relative float-right flex justify-between">
                        <Button _class="bg-white ml-[10px]" on:click={saveVolunteering} title="Save"/>
                        <Button _class="bg-white ml-[10px]" on:click={deleteVolunteering} title="Delete"/>
                    </div>
                </div>
            </div>
        </div>
    {/if}
</div>
