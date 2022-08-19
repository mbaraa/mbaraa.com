<script lang="ts">
    import type Blog from "$lib/models/Blog";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import Separator from "./Separator.svelte";
    export let blog: Blog = { name: "New Blog", description: "" };
    export let isEdit: boolean = false;
    export let isNew: boolean = false;

    async function saveBlog() {
        if (isEdit) {
            await BlogRequests.updateBlog(blog);
        } else if (isNew) {
            await BlogRequests.newBlog(blog);
        }
    }
</script>

<div class="bg-[#2d333b] text-white h-[80vh] flex justify-center">
    <!-- <div class="hidden lg:block w-[779px] "> -->
    <!--     <img src="/images/blogImage.png" /> -->
    <!-- </div> -->
    <div class="text-center w-[90%] h-[100%] ">
        <div
            class="text-white shadow-md shadow-[#979797] p-[35px] bg-[#212121] mx-[20px] md:mx-[40px] mt-[33px] rounded-[8px] md:h-auto  h-[90%] "
        >
            <span
                class="text-center text-[30px] md:text-[40px] m-[30px] text-[#bdbdbd] "
            >
                {#if isEdit}
                    Editing:
                {:else if isNew}
                    New Blog:
                {/if}
                {blog.name}
            </span>
            <br />
            {#if isEdit || isNew}
                <Separator />
                <textarea
                    class="bg-[#2d333b] text-[20px] rounded-[8px] m-[15px] p-[15px] w-[80%]"
                    bind:value={blog.name}
                />
                <textarea
                    class="bg-[#2d333b] text-[20px] rounded-[8px] m-[15px] p-[15px] w-[80%]"
                    bind:value={blog.description}
                />
                <br />
                <button
                    on:click={saveBlog}
                    class="p-[15px] border-[#1b1b1b] border-[1px] rounded-[8px] bg-[#424242] hover:bg-[#6d6d6d] text-white w-[250px]"
                    >Save</button
                >
            {:else}
                <div class="flex justify-center mt-[15px] ">
                    <div class="border-b-[#979797] border-b-[1px] w-[250px] " />
                </div>
                <br />

                <span class="text-[35px] ">
                    {@html blog.description}
                </span>
            {/if}
        </div>
    </div>
</div>
