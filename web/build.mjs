import { readFile } from "node:fs/promises";

const html = await readFile(new URL("./index.html", import.meta.url), "utf8");
const script = await readFile(new URL("./app.js", import.meta.url), "utf8");
if (!html.includes("app.js") || !script.includes("/api/home") || !script.includes("/api/search")) {
  throw new Error("frontend entry points are incomplete");
}
console.log("frontend sources verified");
