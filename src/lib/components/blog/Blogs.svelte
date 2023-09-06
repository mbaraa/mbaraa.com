<script lang="ts">
	import type Blog from "$lib/models/Blog";

	export let blogs: Blog[];

	function formatDate(date: Date): string {
		const _date = new Date(date).toLocaleTimeString("en-US", {
			year: "numeric",
			day: "numeric",
			month: "short"
		});

		const dateStr = _date.toString();
		return dateStr.substring(0, dateStr.lastIndexOf(","));
	}

    function formatNumber(n: number): string {
        return n.toString()
    }
</script>

<div class="w-auto absolute left-[50%] translate-x-[-50%]">
	{#each blogs as blog}
		<a href={`/blog/${blog.publicId}`}>
			<div
				class="transform transition hover:scale-[105%] w-[80vw] p-[30px] m-[30px] bg-[#212121] text-white rounded-[16px] font-[Vistol]"
			>
				<h1 class="font-[1000] text-[25px]">{blog.name}</h1>
				<h1 class="text-[15px]">{blog.description}</h1>
				<div class="py-[5px] absolute top-1 right-5 float-right hidden md:block">
					<span class="block">
						Visited {formatNumber(blog.readTimes)} times
					</span>
				</div>
				<div class="py-[5px] absolute bottom-1 right-5 float-right">
					<span class="block">
						{formatDate(blog.createdAt)}
					</span>
				</div>
			</div>
		</a>
	{/each}
</div>
