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
	const groups = await getProjectGroups();
	if (groups) {
		return new Response(JSON.stringify(groups), { status: 200, headers: jsonResp });
	}
	return new Response("not found", { status: 404 });
};

export const POST: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", { status: 401 });
	}
	const pg: ProjectGroup = await ev.request.json();
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
