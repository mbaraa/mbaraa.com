import * as jwt from "jsonwebtoken";
import config from "$lib/config";
import type { RequestEvent } from "@sveltejs/kit";

export function isAuth(ev: RequestEvent): boolean  {
    const token = ev.request.headers.get("Authorization") ?? "";
    try {
        jwt.verify(token, config.jwtSecret);
    } catch {
        return false;
    }
    return true
}



