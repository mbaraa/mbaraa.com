<script lang="ts">
	import type Project from "$lib/models/Project";
	import Link from "$lib/ui/Link.svelte";

	export let projects: Project[];
</script>

<div class="container">
	<div class="timeline">
		<ul>
			{#each projects as project}
				<li>
					<div class="timeline-content">
						<h3 class="date">
							({project.startYear}-{#if project.endYear}{project.endYear}{:else}Present{/if})
						</h3>
						<h1 class="text-[#20DB8F]">{project.name}</h1>
						<p>{project.description}</p>
						<div class="float-right">
							{#if project.website}
								<Link link={{ name: "Website", link: project.website }} />
							{/if}
							{#if project.sourceCode}
								<Link link={{ name: "Source Code", link: project.sourceCode }} />
							{/if}
							{#if (!project.website && !project.sourceCode) || project.comingSoon}
								<span class="text-[15px] italic">Coming Soon...</span>
							{/if}
						</div>
					</div>
				</li>
			{/each}
		</ul>
	</div>
</div>

<style>
	.container {
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 0 auto;
	}
	.timeline {
		width: 80%;
		height: auto;
		max-width: 800px;
		margin: 0 auto;
		position: relative;
	}

	.timeline ul {
		list-style: none;
	}
	.timeline ul li {
		padding: 20px;
		background-color: #1e1f22;
		color: white;
		border-radius: 10px;
		margin-bottom: 20px;
	}
	.timeline ul li:last-child {
		margin-bottom: 0;
	}
	.timeline-content h1 {
		font-weight: 500;
		font-size: 25px;
		line-height: 30px;
		margin-bottom: 10px;
	}
	.timeline-content p {
		font-size: 16px;
		line-height: 30px;
		font-weight: 300;
	}
	.timeline-content .date {
		font-size: 12px;
		font-weight: 300;
		margin-bottom: 10px;
		letter-spacing: 2px;
	}
	@media only screen and (min-width: 768px) {
		.timeline:before {
			content: "";
			position: absolute;
			top: 0;
			left: 50%;
			transform: translateX(-50%);
			width: 2px;
			height: 100%;
			background-color: gray;
		}
		.timeline ul li {
			width: 50%;
			position: relative;
			margin-bottom: 50px;
		}
		.timeline ul li:nth-child(odd) {
			float: left;
			clear: right;
			transform: translateX(-30px);
			border-radius: 20px 0px 20px 20px;
		}
		.timeline ul li:nth-child(even) {
			float: right;
			clear: left;
			transform: translateX(30px);
			border-radius: 0px 20px 20px 20px;
		}
		.timeline ul li::before {
			content: "";
			position: absolute;
			height: 20px;
			width: 20px;
			border-radius: 50%;
			background-color: gray;
			top: 0px;
		}
		.timeline ul li:nth-child(odd)::before {
			transform: translate(50%, -50%);
			right: -30px;
		}
		.timeline ul li:nth-child(even)::before {
			transform: translate(-50%, -50%);
			left: -30px;
		}
		.timeline-content .date {
			position: absolute;
			top: -30px;
		}
		.timeline ul li:hover::before {
			background-color: #20db8f;
		}
	}
</style>
