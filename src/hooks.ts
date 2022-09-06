import process from "process";
import * as dotenv from "dotenv";
import db from "$lib/db";

dotenv.config()

db.$disconnect().catch(err => {console.error(err)});

process.on('SIGINT', function () {process.exit();}); // Ctrl+C
process.on('SIGTERM', function () {process.exit();}); // docker stop
