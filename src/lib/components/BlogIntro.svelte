<script lang="ts">
    import strings from "$lib/strings";
    import Seperator from "./Seperator.svelte";
</script>

<div class="p-[40px] font-[Vistol]">
    <span
        class="uppercase font-[1000] text-white text-[36px] md:text-[100px] block"
    >
        {strings.en.firstName}&lsquo;s
    </span>
    <span
        class="uppercase text-[36px] md:text-[100px] font-[1000] text-white blog relative left-[-18px] top-[-10px] md:left-[-50px] md:top-[-44px] "
        >log</span
    >
    <p
        class="text-[#20db8f] font-[1000] text-[22px] w-[280px] md:w-[450px] relative left-[50px] top-[-70px] md:left-[100px] md:top-[-200px] "
    >
        {strings.en.blogIntro}
    </p>

    <div class="relative top:top[-40px] md:top-[-120px]">
        <Seperator />
    </div>
</div>

<style>
    .blog {
        writing-mode: vertical-rl;
    }
</style>
