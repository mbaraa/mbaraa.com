<script lang="ts">
    import type ProjectGroup from "$lib/models/ProjectGroup";
    import Project from "./Project.svelte";
    export let group: ProjectGroup;
</script>

<div class="mt-[50px]">
    <h1
        class="text-white font-[SFUI] ml-[15px] md:ml-[50px] text-[35px] italic text-bold"
    >
        {group.name}
    </h1>
    {#each group.projects as project}
        <Project {project} />
    {/each}
</div>
