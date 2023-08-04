import type Blog from "$lib/models/Blog";
import { db, toKebab } from "./index";
import type Experience from "$lib/models/Experience";
import type ProjectGroup from "$lib/models/ProjectGroup";
import { writeFile } from "fs/promises";

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

export async function uploadBlogImage(imageName: string, image: File): Promise<string> {
	const imgB64 = Buffer.from(await image.arrayBuffer()).toString("base64");

	const imageDocRaw = {
		imageName: `${(Math.random() * 100000).toString().substring(0, 4)}_${imageName}`,
		base64: imgB64
	};

	const imageDoc = db.doc(`blog-images/${imageDocRaw.imageName}`);
	const status = await imageDoc.set(imageDocRaw);
	if (!status) {
		return "";
	}
	await writeFile("./static/" + imageDocRaw.imageName, imageDocRaw.base64, { encoding: "base64" });

	return imageDocRaw.imageName;
}

export async function insertProjectsGroup(pg: ProjectGroup): Promise<unknown> {
	pg.publicId = toKebab(pg.name);

	const document = db.doc(`projectGroups/${pg.publicId}`);
	const status = await document.set(pg);

	if (!status) {
		return null;
	}

	return pg;
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
