<script lang="ts">
	import type ProjectGroup from "$lib/models/ProjectGroup";
	import Button from "$lib/ui/Button.svelte";

	export let projectGroup: ProjectGroup;
	let editMode = false;

	async function saveProjectGroup(): Promise<void> {
		let method = "POST";
		if (projectGroup.publicId) {
			method = "PUT";
		}
		await fetch("/api/projectGroup", {
			method: method,
			mode: "cors",
            headers: {
				"Authorization": localStorage.getItem("token") ?? ""
            },
			body: JSON.stringify(projectGroup)
		})
	}

	async function deleteProjectGroup(): Promise<void> {
		await fetch(`/api/projectGroup?id=${projectGroup.publicId}`, {
			method: "DELETE",
			headers: {
				"Authorization": localStorage.getItem("token") ?? ""
			},
		})
	}
</script>

<div class="text-black bg-[#CBCBCB] block rounded-[10px] mb-[10px] last:mb-0 p-[15px] ">
    <div class="flex justify-between">
        <h3>{projectGroup.name}</h3>
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
                                <h1 class="font-[600] px-[10px]">Title:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={projectGroup.name}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">Description:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={projectGroup.description}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <h1 class="font-[600] px-[10px]">Content:</h1>
                            </td>
                            <td>
                            <textarea
                                    bind:value={projectGroup.content}
                                    class="w-[100%] h-[300px] p-[3px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                            </td>
                        </tr>
                        </tbody>
                    </table>
                    <div class="relative float-right flex justify-between">
                        <Button _class="bg-white ml-[10px]" on:click={saveProjectGroup} title="Save"/>
                        <Button _class="bg-white ml-[10px]" on:click={deleteProjectGroup} title="Delete"/>
                    </div>
                </div>
            </div>
        </div>
    {/if}
</div>
