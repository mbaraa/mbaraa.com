<script lang="ts">
    import type Link from "$lib/models/Link";
    export let link: Link;
</script>

<div
    class="italic font-[SFUI] text-[20px] text-[white] font-bold inline-block underline"
>
    <a href={link.link} target="_blank">{link.name}</a>
</div>
