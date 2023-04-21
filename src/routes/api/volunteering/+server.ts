import type { RequestHandler, RequestEvent } from "@sveltejs/kit";
import { getVolunteeringExperiences } from "$lib/db/Getters";
import { isAuth } from "../_auth";
import type Experience from "$lib/models/Experience";
import { insertVolunteeringXP } from "$lib/db/Creators";
import { updateVolunteeringXP } from "$lib/db/Updaters";

const jsonResp = {"Content-Type": "application/json", "Access-Control-Allow-Headers": "Content-Type"}

export const GET: RequestHandler = async (ev: RequestEvent) => {
	const volunteeringXP = await getVolunteeringExperiences();
	if (!volunteeringXP) {
		return new Response("not found", {status: 404});
	}
	return new Response(JSON.stringify(volunteeringXP), { status: 200, headers: jsonResp });
};

export const POST: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	const volunteeringXP: Experience = await ev.request.json();
	if (await insertVolunteeringXP(volunteeringXP)) {
		return new Response("ok", {status: 200});
	}
	return new Response(null, { status: 500 });
};

export const PUT: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	const volunteeringXP: Experience = await ev.request.json();
	if (await updateVolunteeringXP(volunteeringXP)) {
		return new Response("ok", {status: 200});
	}
	return new Response(null, { status: 500 });
};

export const DELETE: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	return new Response("ok", { status: 200 });
};
