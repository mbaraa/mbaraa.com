import type { RequestEvent, RequestHandler } from "@sveltejs/kit";
import { getBlogImage } from "$lib/db/Getters";

const jsonResp = {
	"Content-Type": "application/json",
	"Access-Control-Allow-Headers": "Content-Type"
};
export const GET: RequestHandler = async (ev: RequestEvent) => {
	const id = ev.params["id"] as string;
	const imageB64 = await getBlogImage(id);
	if (!imageB64) {
		return new Response("not found", { status: 404 });
	}
	return new Response(JSON.stringify({ imageB64 }), { status: 200, headers: jsonResp });
};
