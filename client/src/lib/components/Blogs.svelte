<script lang="ts">
    import type Blog from "$lib/models/Blog";

    export let blogs: Blog[];
</script>

{#if blogs.length > 0}
    <div>
        <div class=" ">
            <h1
                class="text-white font-[SFUI] mx-[20px] md:mx-[40px] my-[20px] md:my-[25px] text-[40px] italic text-bold"
            >
                My Blogs...
            </h1>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2">
            {#each blogs as blog}
                <a href={`/blog/${blog.id}`}>
                    <div
                        class="flex justify-center text-[35px] text-white shadow-md shadow-[#979797] p-[35px] bg-[#212121] mx-[20px] md:mx-[40px] my-[20px] rounded-[8px] transform transition hover:translate-y-[-12px] "
                        title="Click to see the blog"
                    >
                        {blog.name}
                    </div>
                </a>
            {/each}
        </div>
    </div>
{:else}
    <h1
        class="text-white text-center font-[SFUI] ml-[20px] md:ml-[55px] mb-[-20px] py-[30px] text-[40px] text-bold h-[100vh]"
    >
        Well it's empty now... 🗿
    </h1>
{/if}
