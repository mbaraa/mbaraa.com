<script lang="ts">
	import type ProjectGroup from "$lib/models/ProjectGroup";
	import Button from "$lib/ui/Button.svelte";

	export let groups: ProjectGroup[];
	export let group: ProjectGroup;
	let editMode = false;

	function getCurrentGroupOrder(): number {
		return groups.findIndex((g) => g.publicId === group.publicId) + 1;
	}

	function handleChangeOrder(e: Event) {
		group.order = Number((e.target as HTMLOptionElement).value);
	}

	async function saveProjectGroup(): Promise<void> {
		let method = "POST";
		if (group.publicId) {
			method = "PUT";
		}
		await fetch("/api/project", {
			method: method,
			mode: "cors",
			headers: {
				Authorization: localStorage.getItem("token") ?? ""
			},
			body: JSON.stringify(group)
		});
	}

	async function deleteProjectGroup(): Promise<void> {
		await fetch(`/api/project?id=${group.publicId}`, {
			method: "DELETE",
			headers: {
				Authorization: localStorage.getItem("token") ?? ""
			}
		});
	}
</script>

<div class="text-black bg-[#CBCBCB] block rounded-[10px] mb-[10px] last:mb-0 p-[15px]">
	<div class="flex justify-between">
		<h3>{group.name}</h3>
		<h3
			class="block font-bold cursor-pointer"
			on:click={() => {
				editMode = !editMode;
			}}
		>
			{#if editMode}{"ᐯ"}{:else}{"ᐳ"}{/if}
		</h3>
	</div>
	{#if editMode}
		<div class="block">
			<div class="font-[Vistol] w-auto">
				<div class="p-[45px] text-[20px] m-[20px] bg-white rounded-[32px]">
					<table class="w-[100%]">
						<tbody>
							<tr>
								<td>
									<h1 class="font-[600] px-[10px]">Title:</h1>
								</td>
								<td>
									<textarea
										bind:value={group.name}
										class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
									/>
								</td>
							</tr>
							<tr>
								<td>
									<h1 class="font-[600] px-[10px]">Description:</h1>
								</td>
								<td>
									<textarea
										bind:value={group.description}
										class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
									/>
								</td>
							</tr>
							<tr>
								<td>
									<h1 class="font-[600] px-[10px]">Order:</h1>
								</td>
								<td>
									<select name="Order" id="project.order" on:change={handleChangeOrder}>
										{#each groups as _, orderBetweenAll}
											<option
												selected={getCurrentGroupOrder() === orderBetweenAll + 1}
												value={orderBetweenAll + 1}>{orderBetweenAll + 1}</option
											>
										{/each}
									</select>
								</td>
							</tr>
							<tr>
								<td>
									<h1 class="font-[600] px-[10px]">Projects:</h1>
								</td>
								<td>
									<Button
										_class="bg-white"
										on:click={() => {
											group.projects.push({});
											group = group;
										}}
										title="+"
									/>
									{#each group.projects as project, i}
										<div class="flex justify-between">
											<div>
												<label for="project.name">Name:</label>
												<textarea
													id="project.name"
													bind:value={project.name}
													class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
												/>
												<label for="project.description">Description:</label>
												<textarea
													id="project.description"
													bind:value={project.description}
													class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
												/>
												<label for="project.sourceCode">Source Code:</label>
												<textarea
													id="project.sourceCode"
													bind:value={project.sourceCode}
													class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
												/>
												<label for="project.website">Website:</label>
												<textarea
													id="project.website"
													bind:value={project.website}
													class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
												/>
												<label for="project.startDate">Start Year:</label>
												<textarea
													id="project.startDate"
													bind:value={project.startYear}
													class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
												/>
												<label for="project.endDate">End Year:</label>
												<textarea
													id="project.endDate"
													bind:value={project.endYear}
													class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
												/>
											</div>
											<Button
												_class=""
												on:click={() => {
													group.projects = group.projects.filter(
														(p) => p.publidId !== project.publicId
													);
												}}
												title="-"
											/>
										</div>
										<hr class="pb-[20px] mt-[20px]" />
									{/each}
								</td>
							</tr>
						</tbody>
					</table>
					<div class="relative float-right flex justify-between">
						<Button _class="bg-white ml-[10px]" on:click={saveProjectGroup} title="Save" />
						<Button _class="bg-white ml-[10px]" on:click={deleteProjectGroup} title="Delete" />
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
