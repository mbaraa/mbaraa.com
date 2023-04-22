import type { RequestEvent, RequestHandler } from "@sveltejs/kit";
import { getAllBlogs, getBlog } from "$lib/db/Getters";
import { isAuth } from "../_auth";
import type Blog from "$lib/models/Blog";
import { updateBlog } from "$lib/db/Updaters";
import { deleteBlog } from "$lib/db/Deleters";
import { insertBlog } from "../../../lib/db/Creators";

const jsonResp = {"Content-Type": "application/json", "Access-Control-Allow-Headers": "Content-Type"}

export const GET: RequestHandler = async (ev: RequestEvent) => {
	const id = ev.url.searchParams.get("id");
	if (id) {
		const blog = await getBlog(id);
		if (blog) {
			return new Response(JSON.stringify(blog), {status: 200, headers: jsonResp});
		}
		return new Response("not found", {status: 404});
	}
	const blogs = await getAllBlogs();
	if (blogs) {
		return new Response(JSON.stringify(blogs), {status: 200, headers: jsonResp});
	}
	return new Response("not found", {status: 404});
};

export const POST: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	const blog: Blog = await ev.request.json();
	if (await insertBlog(blog)) {
		return new Response("ok", {status: 200});
	}
	return new Response(null, { status: 500 });
};

export const PUT: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	const blog: Blog = await ev.request.json();
	if (await updateBlog(blog.publicId as string, blog)) {
		return new Response("ok", {status: 200});
	}
	return new Response(null, {status: 500});
};

export const DELETE: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	if (await deleteBlog(ev.url.searchParams.get("id") as string)) {
		return new Response("ok", {status: 200});
	}
	return new Response(null, {status: 500});
};
