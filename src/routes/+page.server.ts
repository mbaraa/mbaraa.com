import type { Load } from "@sveltejs/kit";
import {getInfo} from "$lib/db/Getters";

export let ssr = true;

export const load: Load = async () => {
	return {
		info: await getInfo()
	};
};
