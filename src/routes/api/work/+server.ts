import type { RequestHandler, RequestEvent } from "@sveltejs/kit";
import { getWorkExperiences } from "$lib/db/Getters";
import { isAuth } from "../_auth";
import type Experience from "$lib/models/Experience";
import { insertWorkXP } from "$lib/db/Creators";
import { updateWorkXP } from "$lib/db/Updaters";
import { deleteWorkXP } from "$lib/db/Deleters";

const jsonResp = {"Content-Type": "application/json", "Access-Control-Allow-Headers": "Content-Type"}

export const GET: RequestHandler = async (ev: RequestEvent) => {
	const workXP = await getWorkExperiences();
	if (!workXP) {
		return new Response("not found", {status: 404});
	}
	return new Response(JSON.stringify(workXP), { status: 200, headers: jsonResp });
};

export const POST: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	const workXP: Experience = await ev.request.json();
	if (await insertWorkXP(workXP)) {
		return new Response("ok", {status: 200});
	}
	return new Response(null, { status: 500 });
};

export const PUT: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	const workXP: Experience = await ev.request.json();
	if (await updateWorkXP(workXP)) {
		return new Response("ok", {status: 200});
	}
	return new Response(null, { status: 500 });
};

export const DELETE: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	if (await deleteWorkXP(ev.url.searchParams.get("id") ?? "")) {
		return new Response("ok", { status: 200 });
	}
	return new Response("oops", { status: 500 });
};
