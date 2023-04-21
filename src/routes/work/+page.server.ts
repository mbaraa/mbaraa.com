import { db } from "$lib/db";
import type Experience from "$lib/models/Experience";
import { type Load } from "@sveltejs/kit";

async function getExperiences(xpName: string): Promise<Experience[]> {
	const experienceRef = db.collection(xpName);
	const snapshot = await experienceRef.get();

	let experience: Experience[] = [];
	snapshot.forEach((projectGroupDoc) => {
		experience.push(projectGroupDoc.data() as Experience);
	});

	return experience.sort((xpI: Experience, xpJ: Experience) => {
		return -(+xpI.startDate - +xpJ.startDate);
	});
}

async function getWorkExperiences(): Promise<Experience[]> {
	return await getExperiences("work");
}

async function getVolunteelingExperiences(): Promise<Experience[]> {
	return await getExperiences("volunteering");
}

export const load: Load = async () => {
	return {
		work: await getWorkExperiences(),
		volunteering: getVolunteelingExperiences()
	};
};
