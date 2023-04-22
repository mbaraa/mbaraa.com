import type Project from "./Project";

export default interface ProjectGroup {
	publicId?: string;
	name: string;
	description: string;
	projects: Project[];
}
