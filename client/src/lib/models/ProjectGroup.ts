import type Project from "./Project"

export default interface ProjectGroup {
    name: string;
    description: string;
    projects: Project[]
}
