import type Experience from "$lib/models/Experience";
import {db} from "./index";
import type Project from "$lib/models/Project";
import type ProjectGroup from "$lib/models/ProjectGroup";
import type Blog from "$lib/models/Blog";
import type Info from "$lib/models/Info";

// TODO
// remove duplicate code and return empty records

export async function getProjectGroups(): Promise<ProjectGroup[]> {
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

export async function getWorkExperiences(): Promise<Experience[]> {
    return await getExperiences("work");
}

export async function getVolunteeringExperiences(): Promise<Experience[]> {
    return await getExperiences("volunteering");
}

export async function getAllBlogs(): Promise<Blog[]> {
    const blogsRef = db.collection("blogs");
    const snapshot = await blogsRef.get();

    let blogs: Blog[] = [];

    snapshot.forEach((doc) => {
        const blog = doc.data() as Blog;
        blog.createdAt = doc.data().createdAt.toDate();
        blog.updatedAt = doc.data().updatedAt.toDate();
        blogs.push(blog);
    });
    return blogs
        .map((blog) => {
            blog.content = "";
            return blog;
        })
        .sort((blogI: Blog, blogJ: Blog) => {
            return -(blogI.createdAt.getTime() - blogJ.createdAt.getTime());
        });
}

export async function getBlog(id: string): Promise<unknown> {
    const blogsRef = db.collection("blogs");
    const snapshot = await blogsRef.where("publicId", "==", id).get();

    if (snapshot.empty) {
        return null;
    }

    let blog: Blog | undefined = undefined;

    snapshot.forEach((doc) => {
        blog = doc.data() as Blog;
        blog.createdAt = doc.data().createdAt.toDate();
        blog.updatedAt = doc.data().updatedAt.toDate();
    });

    return blog;
}

export async function getInfo(): Promise<Info> {
    const infoRef = db.collection("info");
    const snapshot = await infoRef.get();

    let info: Info = { name: "", about: "", blogIntro: "", brief: "", technologies: [] };
    snapshot.forEach((infoDoc) => {
        info = infoDoc.data() as Info;
    });

    return info;
}
