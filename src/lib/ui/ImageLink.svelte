<script lang="ts">
	export let href: string;
	export let imgSrc: string;
</script>

<a {href} target="_blank">
	<img src={imgSrc} alt={href} class="w-6 h-6 inline mx-1" />
</a>
