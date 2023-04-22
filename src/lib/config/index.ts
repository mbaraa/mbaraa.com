import process from "process";

const config: { [keys: string]: string } = {
	firebaseProjectId: process.env["FIREBASE_PROJECT_ID"] ?? "",
	adminPassword: process.env["ADMIN_PASSWORD"] ?? "",
	jwtSecret: process.env["JWT_SECRET"] ?? "",
};

export default config;
