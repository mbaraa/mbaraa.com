import type { Load } from "@sveltejs/kit";
import type Blog from "$lib/models/Blog";
import { db } from "$lib/db";

export const ssr = true;

async function get(id: string): Promise<unknown> {
	const blogsRef = db.collection("blogs");
	const snapshot = await blogsRef.where("publicId", "==", id).get();

	if (snapshot.empty) {
		return null;
	}

	let blog: Blog | undefined = undefined;

	snapshot.forEach((doc) => {
		blog = doc.data() as Blog;
		blog.createdAt = doc.data().createdAt.toDate();
		blog.updatedAt = doc.data().updatedAt.toDate();
	});

	return blog;
}

export const load: Load = async ({ params }: any) => {
	const id = params.id;
	return {
		blog: await get(id)
	};
};
