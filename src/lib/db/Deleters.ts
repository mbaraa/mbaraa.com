import { db } from "./index";

export async function deleteBlog(id: string): Promise<unknown> {
	const document = db.doc(`blogs/${id}`);
	return await document.delete();
}

export async function deleteProjectGroup(id: string): Promise<unknown> {
	const document = db.doc(`projectGroups/${id}`);
	return await document.delete();
}

export async function deleteWorkXP(id: string): Promise<unknown> {
	const document = db.doc(`work/${id}`);
	return await document.delete();
}

export async function deleteVolunteeringXP(id: string): Promise<unknown> {
	const document = db.doc(`volunteering/${id}`);
	return await document.delete();
}
