<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import Bio from "$lib/components/Bio.svelte";
    import Projects from "$lib/components/Projects.svelte";
    import Seperator from "$lib/components/Seperator.svelte";
    import  type { Load} from "@sveltejs/kit"
    import type ProjectGroup from "$lib/models/ProjectGroup";

    let scrollY = 0;
    export let data: Load;
    
</script>

<svelte:head>
    <title>Made By Baraa</title>
</svelte:head>

<Header {scrollY} />
<div class="pb-[85px] block" />
<div class="font-[Vistol]">
    <Bio />
    <Seperator />
    <Projects groups={data.groups}/>
</div>

<svelte:window bind:scrollY />
