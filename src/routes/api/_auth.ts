import jwt from "jsonwebtoken";
import config from "$lib/config";
import type { RequestEvent } from "@sveltejs/kit";

const { verify } = jwt;

export function isAuth(ev: RequestEvent): boolean {
	const token = ev.request.headers.get("Authorization") ?? "";
	try {
		verify(token, config.jwtSecret);
	} catch {
		return false;
	}
	return true;
}
