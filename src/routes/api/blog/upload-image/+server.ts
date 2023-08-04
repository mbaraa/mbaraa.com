import type { RequestEvent, RequestHandler } from "@sveltejs/kit";
import { isAuth } from "../../_auth";
import { uploadBlogImage } from "../../../../lib/db/Creators";

const jsonResp = {
	"Content-Type": "application/json",
	"Access-Control-Allow-Headers": "Content-Type"
};

export const POST: RequestHandler = async (ev: RequestEvent) => {
	if (!isAuth(ev)) {
		return new Response("oops", { status: 401 });
	}
	const image = (await ev.request.formData()).get("image") as File;
	const imageName = ev.request.headers.get("IMAGE_NAME") ?? "";
	if (!image || !imageName) {
		return new Response(null, { status: 400 });
	}
	let err: unknown;
	const fileName = (await uploadBlogImage(imageName, image).catch((_err) => {
		err = _err;
	})) as string;

	return new Response(JSON.stringify({ imageId: fileName }), { status: 200, headers: jsonResp });
};
