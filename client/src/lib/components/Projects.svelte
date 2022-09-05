<script lang="ts">
    import type ProjectGroup from "$lib/models/ProjectGroup";
    import ProjectRequests from "$lib/utils/requests/ProjectRequests";
    import { onMount } from "svelte";
    import { default as ProjectGroupView } from "./ProjectGroup.svelte";

    export let groups = new Array<ProjectGroup>();

    onMount(async () => {
        groups = await ProjectRequests.getProjectGroups();
    });
</script>

{#each groups as group}
    <ProjectGroupView {group} />
{/each}
