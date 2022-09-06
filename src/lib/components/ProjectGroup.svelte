<script lang="ts">
    import type ProjectGroup from "$lib/models/ProjectGroup";
    import Project from "./Project.svelte";
    import Seperator from "./Seperator.svelte";

    export let group: ProjectGroup;
</script>

<div class="font-[Vistol] text-white p-[40px]">
    <h1 class="uppercase font-[1000] text-[30px] ">
        {group.name}
    </h1>
    <p class="text-[#20db8f] italic pb-[10px] text-[16px] ">
        {group.description}
    </p>
    {#each group.projects as project}
        <Project {project} />
    {/each}
</div>
<div class="last:hidden">
    <Seperator />
</div>
