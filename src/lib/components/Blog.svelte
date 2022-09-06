<script lang="ts">
    import type Blog from "$lib/models/Blog";

    export let blog: Blog;

    function formateDate(date: Date): string {
        const _date = new Date(date).toLocaleTimeString("en-US", {
            year: "numeric",
            day: "numeric",
            month: "short",
        });

        const dateStr = _date.toString();
        return dateStr.substring(0, dateStr.lastIndexOf(","));
    }
</script>

<a href={`/blog/${blog.id}`}>
    <div
        class="transform transition hover:scale-[105%] w-[80vw] p-[30px] m-[30px] bg-white rounded-[16px] font-[Vistol] "
    >
        <h1 class="font-[1000] text-[30px] ">{blog.name}</h1>
        <h1 class="text-[20px]">{@html blog.description}</h1>
        <p class="pt-[5px]">
            {formateDate(blog.createdAt)}
        </p>
    </div>
</a>
