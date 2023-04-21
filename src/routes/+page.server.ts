import { db } from "$lib/db";
import type Link from "$lib/models/Link";
import type Info from "$lib/models/Info";
import { type Load } from "@sveltejs/kit";

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

async function getInfo(): Promise<Info> {
	const infoRef = db.collection("info");
	const snapshot = await infoRef.get();

	let info: Info = { name: "", about: "", blogIntro: "", brief: "" };
	snapshot.forEach((infoDoc) => {
		info = infoDoc.data() as Info;
	});

	return info;
}

export const load: Load = async () => {
	return {
		links: await getLinks(),
		info: await getInfo()
	};
};
