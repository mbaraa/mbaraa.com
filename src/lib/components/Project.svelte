<script lang="ts">
    import type Project from "$lib/models/Project";
    import Link from "./Link.svelte";

    export let project: Project;
</script>

<div class="p-[30px] flex ">
    <img
        src={project.imagePath}
        alt={project.name}
        class="rounded-[100%] w-[100px] h-[100px] "
    />
    <div class="px-[20px] md:px-[30px] ">
        <span class="text-[25px] font-[1000] block">
            {project.name}
            <span class="text-[#20db8f] font-[500] ">
                ({project.startYear}-{project.endYear ?? "Present"})
            </span>
        </span>
        <span class="text-[15px] md:text-[15px] block py-[10px] ">
            {@html project.description}
        </span>
        <span>
            {#if project.website}
                <Link
                    link={{
                        name: "View Website...",
                        link: project.website,
                    }}
                />
                &nbsp;
            {/if}
            {#if project.sourceCode}
                <Link
                    link={{
                        name: "Check the madness behind this...",
                        link: project.sourceCode,
                    }}
                />
            {/if}
        </span>
    </div>
</div>
