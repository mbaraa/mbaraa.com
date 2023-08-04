import process from "process";
import * as dotenv from "dotenv";
import bcrypt from "bcrypt";
import { getBlogImages } from "$lib/db/Getters";
import { writeFile } from "fs/promises";

async function hashPassword(): Promise<void> {
	process.env["ADMIN_PASSWORD"] = await bcrypt
		.hash(process.env["ADMIN_PASSWORD"] ?? "", 4)
		.then((hashed: string) => hashed);
}

function stop() {
	console.log("stopped!");
	process.exit();
}

async function downloadImages(): Promise<void> {
	console.log("downloading images from fire store...");
	const images = await getBlogImages();
	if (!images) {
		return;
	}

	for await (const image of images) {
		const keys = Object.keys(image);
		if (!keys.includes("imageName") || !keys.includes("base64")) {
			continue;
		}
		console.log(`copying file: ${image.imageName}...`);
		await writeFile("./static/" + image.imageName, image.base64, { encoding: "base64" });
		console.log(`copied file: ${image.imageName} 🤘`);
	}
	console.log("all good 🤘");
}

dotenv.config();
hashPassword();
downloadImages();

process.on("SIGINT", stop); // Ctrl+C
process.on("SIGTERM", stop); // docker stop
