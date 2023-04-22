import type { Load } from "@sveltejs/kit";
import { getAllBlogs, getInfo } from "$lib/db/Getters";

export const ssr = true;

export const load: Load = async () => {
	return {blogs: await getAllBlogs(), info: await getInfo()};
};
