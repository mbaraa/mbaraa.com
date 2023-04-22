import type { RequestEvent, RequestHandler } from "@sveltejs/kit";
import * as jwt from "jsonwebtoken";
import config from "$lib/config";
import bcrypt from "bcrypt";
import { isAuth } from "../../../_auth";

export const GET: RequestHandler = async (ev: RequestEvent) => {
    if (!isAuth(ev)) {
        return new Response("oops", {status: 401});
    }
    return new Response("ok", {status: 200});
};

export const POST: RequestHandler = async (ev: RequestEvent) => {
    const password = (await ev.request.json().catch(() => {return {password: ""}}) as { password: string }).password;

	const ok = await bcrypt.compare(password, config.adminPassword);

	if (!ok) {
		return new Response("oops :)", {status: 401});
	}

    const token = jwt.sign({
        "my_name_is": Math.random().toString(),
        "walter_hartwell_white": Math.random().toString(),
        "i_live_in_308": "negra_aroya_lane",
        "albuquerque_new_mexico_87104": Math.random().toString()
    }, config.jwtSecret);

    return new Response(JSON.stringify({token: token}), {status: 200});
};