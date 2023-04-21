import type Blog from "$lib/models/Blog";
import { db, toKebab } from "./index";
import type Experience from "$lib/models/Experience";

export async function insertBlog(blog: Blog): Promise<unknown> {
	blog.publicId = toKebab(blog.name);
	blog.createdAt = new Date();
	blog.updatedAt = new Date();

	const document = db.doc(`blogs/${blog.publicId}`);
	const status = await document.set(blog);

	if (!status) {
		return null;
	}

	return blog;
}

async function insertXP(xp: Experience, xpName: string): Promise<unknown> {
	xp.publicId = toKebab(xp.name);
	const document = db.doc(`${xpName}/${xp.publicId}`);
	const status = await document.set(xp);
	if (!status) {
		return null;
	}
	return xp;
}

export async function insertWorkXP(xp: Experience): Promise<unknown> {
	return await insertXP(xp, "work");
}

export async function insertVolunteeringXP(xp: Experience): Promise<unknown> {
	return await insertXP(xp, "volunteering");
}