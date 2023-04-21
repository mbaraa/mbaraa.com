import { db } from "$lib/db";
import type Info from "$lib/models/Info";
import { type Load } from "@sveltejs/kit";

export let ssr = true;

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
	return {
		info: await getInfo()
	};
};
