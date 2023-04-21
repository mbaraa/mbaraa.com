import { db } from "$lib/db";
import type ProjectGroup from "$lib/models/ProjectGroup";
import type Link from "$lib/models/Link";
import { type Load } from "@sveltejs/kit";
import type Info from "$lib/models/Info";

export let ssr = true;

async function getLinks(): Promise<Link[]> {
	const linksRef = db.collection("links");
	const snapshot = await linksRef.get();

	let links: Link[] = [];
	snapshot.forEach((linkDoc) => {
		links.push(linkDoc.data() as Link);
	});

	return links;
}

async function getProjectGroups(): Promise<ProjectGroup[]> {
	const projectGroupsRef = db.collection("projectGroups");
	const snapshot = await projectGroupsRef.get();

	let projectGroups: ProjectGroup[] = [];
	snapshot.forEach((projectGroupDoc) => {
		projectGroups.push(projectGroupDoc.data() as ProjectGroup);
	});

	return projectGroups;
}

async function getInfo(): Promise<Info> {
	const infoRef = db.collection("info");
	const snapshot = await infoRef.get();

	let info: Info = { name: "", about: "", blogIntro: "" };
	snapshot.forEach((infoDoc) => {
		info = infoDoc.data() as Info;
	});

	return info;
}

export const load: Load = async () => {
	return {
		groups: await getProjectGroups(),
		links: await getLinks(),
		info: await getInfo()
	};
};
