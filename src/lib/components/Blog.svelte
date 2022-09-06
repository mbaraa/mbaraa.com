<script lang="ts">
    import type Blog from "$lib/models/Blog";

    export let blog: Blog;
</script>

<a href={`/blog/${blog.id}`}>
    <div
        class="transform transition hover:scale-[105%] w-[80vw] p-[30px] m-[30px] bg-white rounded-[16px] font-[Vistol] "
    >
        <h1 class="font-[1000] text-[30px] ">{blog.name}</h1>
    </div>
</a>
