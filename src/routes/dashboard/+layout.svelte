<script lang="ts">
	import { onMount } from "svelte";
	import Sidebar from "$lib/components/dashboard/Sidebar.svelte";
	import { goto } from "$app/navigation";

	let authed = false;
	let password = "";

	async function login(): Promise<void> {
		authed = await fetch("/api/admin/auth/session", {
			method: "POST",
			mode: "cors",
			body: JSON.stringify({password: password})
		})
			.then((resp) => resp.json())
			.then(resp => {
				localStorage.setItem("token", resp.token);
				return true
			})
			.catch(() => false);
	}

	async function handleEnter(e: KeyboardEvent) {
		if (e.key == "Enter") {
			await login()
		}
	}

	onMount(async () => {
		authed = await fetch("/api/admin/auth/session", {
			method: "GET",
			headers: {
				"Authorization": localStorage.getItem("token") ?? ""
			}
		})
			.then((resp) => resp.ok)
			.catch(() => false);

		if (authed) {
			await goto("/dashboard/info");
		}
	});
</script>
<div class="font-[Vistol]">
    {#if authed}
        <div class="flex justify-between font-[Vistol]">
            <Sidebar/>
            <div class="w-full h-full p-[20px] md:pr-[50px]">
                <slot/>
            </div>
        </div>
    {:else}
        <div class="mx-[30px]">
            <input class="p-[5px] rounded-[5px]" bind:value={password} on:keypress={handleEnter}
                   placeholder="Enter Password"/>
            <button class="text-white p-[5px] border-[1px] border-[white] rounded-[5px]"
                    on:click={login}>Login
            </button>
        </div>
    {/if}
</div>