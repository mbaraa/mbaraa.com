import type { Load } from "@sveltejs/kit";
import {getProjectGroups} from "$lib/db/Getters";

export const ssr = true;

export const load: Load = async () => {
	return {
		groups: await getProjectGroups()
	};
};
