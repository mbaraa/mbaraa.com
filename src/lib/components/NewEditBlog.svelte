<script lang="ts">
    import type Blog from "$lib/models/Blog";
    import BlogRequests from "$lib/utils/requests/BlogRequests";

    export let blog: Blog;

    const styles = `<style>
        code {
            width: 100%;
            background-color: #696969;
            border-radius: 8px;
            padding: 5px;
            color: white;
        }
</style>`;

    async function saveBlog() {
        if (blog.id) {
            await BlogRequests.updateBlog(blog);
        } else {
            await BlogRequests.newBlog(blog);
        }
    }
</script>

{#if blog}
    <div
        class="font-[Vistol] absolute left-[50%] translate-x-[-50%] py-[50px] w-auto "
    >
        <div
            class="p-[45px] text-[20px] m-[20px] bg-white rounded-[32px] w-[90vw] "
        >
            <table class="w-[100%]">
                <tbody>
                    <tr>
                        <td>
                            <h1 class="font-[600] px-[10px]">Title:</h1>
                        </td>
                        <td>
                            <textarea
                                bind:value={blog.name}
                                class="w-[200px] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "
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
                                class="w-[200px] h-[50px] p-[3px] text-[15px] rounded-[8px] border-[1px] border-[#000] "
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
                                class="w-[100%] h-[1000px] p-[3px] rounded-[8px] border-[1px] border-[#000] "
                            />
                        </td>
                    </tr>
                </tbody>
            </table>
            <button
                on:click={saveBlog}
                class="border-[1px] text-black border-[#000] p-[5px] rounded-[8px] "
                >Save</button
            >
        </div>
    </div>
{/if}
