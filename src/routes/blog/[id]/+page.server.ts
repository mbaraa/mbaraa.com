import type { Load } from "@sveltejs/kit";
import { getBlog } from "$lib/db/Getters";
import { updateBlog } from "$lib/db/Updaters";
import type Blog from "$lib/models/Blog";

export const ssr = true;

export const load: Load = async ({ params }: any) => {
	const id = params.id;
	const blog = await getBlog(id) as Blog;
	if (blog.readTimes) {
		blog.readTimes++;
	} else {
		blog.readTimes = 1;
	}
	await updateBlog(id, blog);
	return {
		blog: blog,
	};
};
