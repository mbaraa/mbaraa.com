<script lang="ts">
    import Requests from "$lib/utils/requests/Requests";
    import { onMount } from "svelte";
    import { default as LinkV } from "./Link.svelte";
    import type Link from "$lib/models/Link";

    let links: Link[] | never = [];

    async function fetchLinks() {
        links = [
            { name: "GitHub", link: "https://github.com/mbaraa" },
            { name: "Twitter", link: "https://twitter.com/mbaraa271" },
            { name: "LinkedIn", link: "https://linkedin.com/in/-mbaraa-" },
        ];
    }

    onMount(async () => {
        await fetchLinks();
    });
</script>

<div class="relative bg-[#2d333b] h-[100px]">
    <div class="absolute top-[35px] right-[20px]">
        {#each links as link (link.name)}
            <LinkV {link} />
            <span class="last:hidden text-white">&nbsp;•&nbsp;</span>
        {/each}
    </div>
</div>
