import process from "process";
import * as dotenv from "dotenv";
import bcrypt from "bcrypt";

async function hashPassword(): Promise<void> {
	process.env["ADMIN_PASSWORD"] = await bcrypt
		.hash(process.env["ADMIN_PASSWORD"] ?? "", 4)
		.then((hashed: string) => hashed);
}

function stop(): void {
	console.log("stopped!");
	process.exit();
}

dotenv.config();
hashPassword();

process.on("SIGINT", stop); // Ctrl+C
process.on("SIGTERM", stop); // docker stop
