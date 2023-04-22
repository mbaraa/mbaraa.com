<script lang="ts">
	import type Info from "$lib/models/Info";
	import Button from "$lib/ui/Button.svelte";
	import { onMount } from "svelte";

	let info: Info;
	let editMode = false;

	async function saveInfo(): Promise<void> {
		await fetch("/api/info", {
			method: "PUT",
			mode: "cors",
			headers: {
				"Authorization": localStorage.getItem("token") ?? ""
			},
			body: JSON.stringify(info)
		})
	}

	onMount(async () => {
		info = await fetch("/api/info", {
			method: "GET",
			mode: "cors"
		})
			.then(resp => resp.json())
			.then(resp => resp)
			.catch(() => {
			});
	})
</script>

{#if info}
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
                                    bind:value={info.name}
                                    class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <h1 class="font-[600] px-[10px]">Brief:</h1>
                        </td>
                        <td>
                            <textarea
                                    bind:value={info.brief}
                                    class="w-[100%] h-[150px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <h1 class="font-[600] px-[10px]">About:</h1>
                        </td>
                        <td>
                            <textarea
                                    bind:value={info.about}
                                    class="w-[100%] h-[150px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <h1 class="font-[600] px-[10px]">Technologies:</h1>
                        </td>
                        <td>
                            <Button _class="bg-white" on:click={() => {info.technologies.push(""); info = info}}
                                    title="+"/>
                            {#each info.technologies as technology}
                                <div class="flex justify-between">
                                    <textarea
                                            bind:value={technology}
                                            class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "></textarea>
                                    <Button _class=""
                                            on:click={() => {info.technologies = info.technologies.filter((tech) => tech !== technology)}}
                                            title="-"/>
                                </div>
                            {/each}
                        </td>
                    </tr>

                    </tbody>
                </table>
                <div class="relative float-right flex justify-between">
                    <Button _class="bg-white ml-[10px]" on:click={saveInfo} title="Save"/>
                </div>
            </div>
        </div>
    </div>
{/if}