import type { Load } from "@sveltejs/kit";
import {getVolunteeringExperiences, getWorkExperiences} from "$lib/db/Getters";

export const ssr = true;

export const load: Load = async () => {
	return {
		work: await getWorkExperiences(),
		volunteering: getVolunteeringExperiences()
	};
};
