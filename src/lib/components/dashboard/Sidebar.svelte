<script lang="ts">
	import RouteLink from "$lib/ui/RouteLink.svelte";
	import { page } from "$app/stores";

	$: path = $page.url.pathname;
</script>

<div class="hidden md:block w-[300px] h-[89.3vh] border-r-[1px] border-r-[#ddd] text-white">
    <div class="my-[20px] mx-[40px] ">
        {#key path}
        <span class="block font-bold text-[20px]  my-[10px] ">
            Pages
        </span>
            <span class="block text-[15px] my-[10px] ">
                <RouteLink
                        title="Blog"
                        href="/dashboard/blog"
                />
            </span>
            <span class="block text-[15px] my-[10px] ">
                <RouteLink
                        title="Info"
                        href="/dashboard/info"
                />
            </span>
            <span class="block text-[15px] my-[10px] ">
                <RouteLink
                        title="Work"
                        href="/dashboard/work"
                />
            </span>
            <span class="block text-[15px] my-[10px] ">
                <RouteLink
                        title="Volunteering"
                        href="/dashboard/volunteering"
                />
            </span>
        {/key}
    </div>
</div>
