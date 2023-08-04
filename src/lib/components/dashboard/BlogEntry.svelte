<script lang="ts">
	import type Blog from "$lib/models/Blog";
	import Button from "$lib/ui/Button.svelte";

	export let blog: Blog;
	let editMode = false;

	async function saveBlog(): Promise<void> {
		let method = "POST";
		if (blog.publicId) {
			method = "PUT";
		}
		await fetch("/api/blog", {
			method: method,
			mode: "cors",
			headers: {
				Authorization: localStorage.getItem("token") ?? ""
			},
			body: JSON.stringify(blog)
		});
	}

	async function deleteBlog(): Promise<void> {
		await fetch(`/api/blog?id=${blog.publicId}`, {
			method: "DELETE",
			headers: {
				Authorization: localStorage.getItem("token") ?? ""
			}
		});
	}

	let fileName: string, imageFile: File;
	let status: string = "nothing";

	async function uploadImage(): Promise<void> {
		status = `uploading`;
		const formData = new FormData();
		formData.append("image", imageFile);
		const imageId = await fetch(`/api/blog/upload-image/${fileName}`, {
			method: "POST",
			headers: {
				Authorization: localStorage.getItem("token") ?? ""
			},
			body: formData
		})
			.then((resp) => resp.json())
			.then((data) => {
				console.log(data);
				return data["imageId"];
			});
		status = `uploaded file: /img/${imageId}`;
	}
</script>

<div class="text-black bg-[#CBCBCB] block rounded-[10px] mb-[10px] last:mb-0 p-[15px]">
	<div class="flex justify-between">
		<h3>{blog.name}</h3>
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
										bind:value={blog.name}
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
										bind:value={blog.description}
										class="w-[100%] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000]"
									/>
								</td>
							</tr>
							<tr>
								<td>
									<h1 class="font-[600] px-[10px]">Content:</h1>
								</td>
								<td>
									<textarea
										bind:value={blog.content}
										class="w-[100%] h-[300px] p-[3px] rounded-[8px] border-[1px] border-[#000]"
									/>
								</td>
							</tr>
							<tr>
								<td>
									<h1 class="font-[600] px-[10px]">Images:</h1>
								</td>
								<td>
									<input
										type="file"
										on:change={(e) => {
											const inputEv = e.target;
											if (inputEv.files) {
												imageFile = inputEv.files[0];
											}
										}}
									/>
									<br />
									<input type="text" bind:value={fileName} placeholder="Image file name" />
									<br />
									<button on:click={uploadImage}>Upload</button>
									<br />
									<span>Status: {status}</span>
								</td>
							</tr>
						</tbody>
					</table>
					<div class="relative float-right flex justify-between">
						<Button _class="bg-white ml-[10px]" on:click={saveBlog} title="Save" />
						<Button _class="bg-white ml-[10px]" on:click={deleteBlog} title="Delete" />
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
