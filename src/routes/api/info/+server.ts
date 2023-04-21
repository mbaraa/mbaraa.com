import type {RequestEvent, RequestHandler} from "@sveltejs/kit";
import {getInfo} from "$lib/db/Getters";
import type Info from "$lib/models/Info";
import {updateInfo} from "$lib/db/Updaters";
import {isAuth} from "../_auth";

const jsonResp = {"Content-Type": "application/json", "Access-Control-Allow-Headers": "Content-Type"}

export const GET: RequestHandler = async (ev: RequestEvent) => {
	const info = await getInfo();
	if (!info) {
		return new Response("not found", {status: 404})
	}

	return new Response(JSON.stringify(info), {status: 200, headers: jsonResp});
};

export const PUT: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", {status: 401});
	}
	const newInfo = await ev.request.json() as Info;

	if (await updateInfo(newInfo)) {
		return new Response("ok", {status: 200});
	}

	return new Response("oops", {status: 500})
};