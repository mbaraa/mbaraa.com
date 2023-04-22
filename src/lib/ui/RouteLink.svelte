<script lang="ts">
	import { page } from "$app/stores";

	export let title: string;
	export let href = "#";
	export let _class: string = "";

	let isActive = $page.url.pathname.includes(href);
</script>

<a {href} class:active={isActive} class={_class}>
    {#if isActive}|> {/if}{title}
</a>

<style>
    a.active {
        color: #999;
        font-weight: bold;
    }
</style>
