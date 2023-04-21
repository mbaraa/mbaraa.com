import { db } from "$lib/db";
import type { Load } from "@sveltejs/kit";
import type ProjectGroup from "$lib/models/ProjectGroup";
import type Project from "$lib/models/Project";

export const ssr = true;

async function getProjectGroups(): Promise<ProjectGroup[]> {
	const projectGroupsRef = db.collection("projectGroups");
	const snapshot = await projectGroupsRef.get();

	let projectGroups: ProjectGroup[] = [];
	snapshot.forEach((projectGroupDoc) => {
		const projectGroup = projectGroupDoc.data() as ProjectGroup;

		projectGroup.projects = projectGroup.projects.sort((projectI: Project, projectJ: Project) => {
			return -(+projectI.startYear - +projectJ.startYear);
		});

		projectGroups.push(projectGroup);
	});

	return projectGroups;
}

export const load: Load = async () => {
	return {
		groups: await getProjectGroups()
	};
};
