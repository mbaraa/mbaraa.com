import type Experience from "$lib/models/Experience";
import { db } from "./index";
import type Project from "$lib/models/Project";
import type ProjectGroup from "$lib/models/ProjectGroup";
import type Blog from "$lib/models/Blog";
import type Info from "$lib/models/Info";

export async function getProjectGroups(): Promise<unknown> {
	const projectGroupsRef = db.collection("projectGroups");
	const snapshot = await projectGroupsRef.get();

	if (snapshot.empty) {
		return null;
	}

	let projectGroups: ProjectGroup[] = [];
	snapshot.forEach((projectGroupDoc) => {
		const projectGroup = projectGroupDoc.data() as ProjectGroup;

		projectGroup.projects = projectGroup.projects.map((p: any) => {
			const project = p as Project;
			if (p.startYear) {
				project.startYear = new Date(p.startYear._seconds * 1000);
			}
			if (p.endYear) {
				project.endYear = new Date(p.endYear._seconds * 1000);
			}
			return project;
		});

		projectGroup.projects = projectGroup.projects.sort((projectI: Project, projectJ: Project) => {
			return -(projectI.startYear.getTime() - projectJ.startYear.getTime());
		});

		projectGroups.push(projectGroup);
	});

	return projectGroups.sort((groupI: ProjectGroup, groupJ: ProjectGroup) => {
		return groupI.order - groupJ.order;
	});
}

async function getExperiences(xpName: string): Promise<unknown> {
	const experienceRef = db.collection(xpName);
	const snapshot = await experienceRef.get();

	if (snapshot.empty) {
		return null;
	}

	let experience: Experience[] = [];
	snapshot.forEach((projectGroupDoc) => {
		experience.push(projectGroupDoc.data() as Experience);
	});

	return experience.sort((xpI: Experience, xpJ: Experience) => {
		return -(+xpI.startDate - +xpJ.startDate);
	});
}

export async function getWorkExperiences(): Promise<unknown> {
	return await getExperiences("work");
}

export async function getVolunteeringExperiences(): Promise<unknown> {
	return await getExperiences("volunteering");
}

export async function getAllBlogs(): Promise<unknown> {
	const blogsRef = db.collection("blogs");
	const snapshot = await blogsRef.get();

	if (snapshot.empty) {
		return null;
	}

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

export async function getBlogImages(): Promise<{ imageName: string; base64: string }[]> {
	const blogImagesRef = db.collection("blog-images");
	const snapshot = await blogImagesRef.get();

	if (snapshot.empty) {
		return [];
	}

	let images: { imageName: string; base64: string }[] = [];

	snapshot.forEach((doc: any) => {
		const image = doc.data() as { imageName: string; base64: string };
		images.push(image);
	});

	return images;
}

export async function getInfo(): Promise<unknown> {
	const infoRef = db.collection("info");
	const snapshot = await infoRef.get();

	if (snapshot.empty) {
		return null;
	}

	let info: Info = { name: "", about: "", blogIntro: "", brief: "", technologies: [] };
	snapshot.forEach((infoDoc) => {
		info = infoDoc.data() as Info;
	});

	return info;
}
