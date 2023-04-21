import { db } from "$lib/db";
import type { Load } from "@sveltejs/kit";
import type ProjectGroup from "$lib/models/ProjectGroup";

async function getProjectGroups(): Promise<ProjectGroup[]> {
	const projectGroupsRef = db.collection("projectGroups");
	const snapshot = await projectGroupsRef.get();

	let projectGroups: ProjectGroup[] = [];
	snapshot.forEach((projectGroupDoc) => {
		projectGroups.push(projectGroupDoc.data() as ProjectGroup);
	});

	return projectGroups.sort((projectI: Project, projectJ: Project) => {
		return -(+projectI.startYear - +projectJ.startYear);
	});
}

export const load: Load = async () => {
	return {
		groups: await getProjectGroups()
	};
};
