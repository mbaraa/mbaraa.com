<script lang="ts">
    import type Link from "$lib/models/Link";
    import LinksRequests from "$lib/utils/requests/LinksRequests";
    import { onMount } from "svelte";
    import Seperator from "./Seperator.svelte";
    import strings from "$lib/strings";
    import { default as LinkView } from "./Link.svelte";

    export let scrollY: number;

    let links = new Array<Link>();

    $: showName = scrollY > 120;
    $: showSeparator = scrollY > 320;

    onMount(async () => {
        links = await LinksRequests.getLinks();
    });
</script>

<div class="fixed">
    <header
        class="w-[100vw] bg-black flex justify-center sm:justify-between text-white font-[Vistol] p-[30px] text-[20px] "
    >
        <a class="font-[1000] hidden sm:block text-[30px]" href="/">
            {#if showName}
                {strings.en.name}
            {/if}
        </a>

        <nav>
            <ul class="m-0 list-none flex gap-[20px] ">
                {#each links as link}
                    <li>
                        <LinkView {link} />
                    </li>
                {/each}
            </ul>
        </nav>
    </header>

    {#if showSeparator}
        <div class="relative top-[-10px] ">
            <Seperator />
        </div>
    {/if}
</div>
