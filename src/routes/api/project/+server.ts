import type { RequestEvent, RequestHandler } from "@sveltejs/kit";
import { getProjectGroups } from "$lib/db/Getters";
import { isAuth } from "../_auth";
import { updateProjectGroup } from "$lib/db/Updaters";
import { deleteProjectGroup } from "$lib/db/Deleters";
import { insertProjectsGroup } from "$lib/db/Creators";
import type ProjectGroup from "$lib/models/ProjectGroup";

const jsonResp = {
	"Content-Type": "application/json",
	"Access-Control-Allow-Headers": "Content-Type"
};

export const GET: RequestHandler = async (ev: RequestEvent) => {
	const groups = (await getProjectGroups()) as ProjectGroup[];
	if (groups) {
		groups.forEach((pg) => {
			pg.projects.forEach((p) => {
				p.startYear = new Date(p.startYear);
				if (p.endYear) {
					p.endYear = new Date(p.endYear);
				}
			});
		});
		return new Response(JSON.stringify(groups), { status: 200, headers: jsonResp });
	}
	return new Response("not found", { status: 404 });
};

export const POST: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", { status: 401 });
	}
	const pg: ProjectGroup = await ev.request.json();
	pg.projects.forEach((p) => {
		p.startYear = new Date(p.startYear);
		if (p.endYear) {
			p.endYear = new Date(p.endYear);
		}
	});
	if (await insertProjectsGroup(pg)) {
		return new Response("ok", { status: 200 });
	}
	return new Response(null, { status: 500 });
};

export const PUT: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", { status: 401 });
	}
	const pg: ProjectGroup = await ev.request.json();
	pg.projects.forEach((p) => {
		p.startYear = new Date(p.startYear);
		if (p.endYear) {
			p.endYear = new Date(p.endYear);
		}
	});
	if (await updateProjectGroup(pg.publicId ?? "", pg)) {
		return new Response("ok", { status: 200 });
	}
	return new Response(null, { status: 500 });
};

export const DELETE: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", { status: 401 });
	}
	if (await deleteProjectGroup(ev.url.searchParams.get("id") as string)) {
		return new Response("ok", { status: 200 });
	}
	return new Response(null, { status: 500 });
};
