import { type Load } from "@sveltejs/kit";
import type Blog from "$lib/models/Blog";
import type Info from "$lib/models/Info";
import { db } from "$lib/db";

export const ssr = true;

async function getAllBlogs(): Promise<Blog[]> {
	const blogsRef = db.collection("blogs");
	const snapshot = await blogsRef.get();

	let blogs: Blog[] = [];

	snapshot.forEach((doc) => {
		const blog = doc.data() as Blog;
		blog.createdAt = doc.data().createdAt.toDate();
		blog.updatedAt = doc.data().updatedAt.toDate();
		blogs.push(blog);
	});
	return blogs
		.map((blog) => {
			blog.content = "";
			return blog;
		})
		.sort((blogI: Blog, blogJ: Blog) => {
			return -(blogI.createdAt.getTime() - blogJ.createdAt.getTime());
		});
}

async function getInfo(): Promise<Info> {
	const infoRef = db.collection("info");
	const snapshot = await infoRef.get();

	let info: Info = { name: "", about: "", blogIntro: "", brief: "", technologies: [] };
	snapshot.forEach((infoDoc) => {
		info = infoDoc.data() as Info;
	});

	return info;
}

export const load: Load = async () => {
	return { blogs: await getAllBlogs(), info: await getInfo() };
};
