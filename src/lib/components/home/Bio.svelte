<script lang="ts">
	export let name: string = "";
	export let brief: string = "";
</script>

<div class="w-[100%] h-[100%] font-[Vistol]">
	<div
		class="w-[80%] md:w-[50%] absolute left-[50%] translate-x-[-50%] top-[50%] translate-y-[-50%]"
	>
		<h3 class="text-[#64FFDA] text-[20px]">Hi, I am</h3>
		<h1 class="text-white font-bold text-[35px] sm:text-[45px] lg:text-[70px]">{name}.</h1>
		<h2 class="text-[#ABABAB] font-bold text-[30px]">I build stuff, mostly web apps.</h2>
		<p class="text-[#20DB8F] text-[20px]">{brief}</p>
		<button
			class="text-[#64FFDA] rounded-[4px] border-[2px] border-[#64FFDA] p-[5px] hover:bg-[#64FFDA] hover:text-black mt-[10px]"
		>
			<a href="/projects">Check out my projects!</a>
		</button>
	</div>
</div>
