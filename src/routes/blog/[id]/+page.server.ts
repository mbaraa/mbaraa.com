import type { Load } from "@sveltejs/kit";
import {getBlog} from "$lib/db/Getters";

export const ssr = true;

export const load: Load = async ({ params }: any) => {
	const id = params.id;
	return {
		blog: await getBlog(id)
	};
};
