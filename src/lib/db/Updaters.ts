import type Blog from "$lib/models/Blog";
import { db } from "./index";
import type Info from "$lib/models/Info";
import type Experience from "$lib/models/Experience";
import type ProjectGroup from "$lib/models/ProjectGroup";

export async function updateBlog(id: string, blog: Blog): Promise<unknown> {
	const document = db.doc(`blogs/${blog.publicId}`);
	const status = await document.update({
		content: blog.content,
		description: blog.description,
		name: blog.name,
		updatedAt: new Date()
	});

	if (!status) {
		return null;
	}
	return blog;
}

export async function updateProjectGroup(id: string, pg: ProjectGroup): Promise<unknown> {
	const document = db.doc(`projectGroups/${id}`);
	pg.publicId = id;
	const status = await document.update({ ...pg });

	if (!status) {
		return null;
	}
	return pg;
}

export async function updateInfo(info: Info): Promise<unknown> {
	const document = db.doc("info/info");
	const status = await document.update({ ...info });

	if (!status) {
		return null;
	}

	return info;
}

async function updateXP(xp: Experience, xpName: string): Promise<unknown> {
	const document = db.doc(`${xpName}/${xp.publicId}`);
	const status = await document.update({
		description: xp.description,
		name: xp.name,
		roles: xp.roles,
		startDate: xp.startDate,
		endDate: xp.endDate,
		location: xp.location
	});

	if (!status) {
		return null;
	}
	return xp;
}

export async function updateWorkXP(xp: Experience): Promise<unknown> {
	return await updateXP(xp, "work");
}

export async function updateVolunteeringXP(xp: Experience): Promise<unknown> {
	return await updateXP(xp, "volunteering");
}
