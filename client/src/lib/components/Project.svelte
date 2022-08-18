<script lang="ts">
    import type Project from "$lib/models/Project";
    import Link from "./Link.svelte";

    export let project: Project;
</script>

<div
    class="text-white grid grid-cols-4 ml-[25px] sm:ml-[50px] mr-0 font-[SFUI] mt-[25px]"
>
    <div>
        <img
            alt={project.name}
            src={project.imagePath}
            class="rounded-[100%]"
            width={150}
            height={150}
        />
    </div>
    <div class="pl-[15px] w-[250%]">
        <span class="italic text-[22px] font-bold">
            {project.name} ({project.startYear}-{project.endYear ?? "Present"})
        </span>
        <div class="mt-[10px]" />
        <span class="w-[250%] text-[18px]">
            {project.description}
        </span>
        <br />
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
    </div>
</div>
