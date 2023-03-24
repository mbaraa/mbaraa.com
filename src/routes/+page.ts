import ProjectRequests from "$lib/utils/requests/ProjectRequests";
import type { Load } from "@sveltejs/kit";

export const ssr = true;
export const prerender = true;

export const load: Load = async () => {
  return {
    groups: await ProjectRequests.getProjectGroups(),
  };
};
