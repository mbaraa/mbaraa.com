import { Firestore } from "@google-cloud/firestore";

export const db = new Firestore({
	projectId: process.env["FIREBASE_PROJECT_ID"],
	keyFilename: "./firebase-key.json"
});

export function toKebab(s: string): string {
	return s
		.toLowerCase()
		.replaceAll(" ", "-")
		.split("")
		.filter((s: string) => s.match(/^[a-z0-9-]+$/i))
		.join("");
}