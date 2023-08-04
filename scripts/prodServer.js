import { handler } from "./handler.js";
import express from "express";

const app = express();
app.use("/img", express.static("static"));
console.log("oi hello there mate");
app.use(handler);

app.listen(3000, () => {
	console.log("listening on port 3000");
});
