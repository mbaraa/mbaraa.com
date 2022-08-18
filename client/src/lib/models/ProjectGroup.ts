import type Project from "./Project"

export default interface ProjectGroup {
    name: string;
    projects: Project[]
}
