import auth from "$lib/auth";
import type Blog from "$lib/models/Blog";
import { v4 as uuidv4 } from "uuid";
import { get, getAll } from "$lib/db";

export async function GET({ url }: any): Promise<Response> {
  const id = url.searchParams.get("id");
  if (id) {
    return new Response(JSON.stringify(await get(id)), { status: 200 });
  }

  const blogs = (await getAll())
    .map((blog) => {
      blog.content = "";
      return blog;
    })
    .sort((blogI: Blog, blogJ: Blog) => {
      return -(
        new Date(blogI.createdAt).getTime() -
        new Date(blogJ.createdAt).getTime()
      );
    });

  return new Response(JSON.stringify(blogs));
}

// export async function POST({ request }: any): Promise<Response> {
//   if (!auth(request)) {
//     return new Response(null, { status: 401 });
//   }
//
// }
//
// export async function PUT({ request }: any): Promise<Response> {
//   if (!auth(request)) {
//     return new Response(null, {
//       status: 401,
//     });
//   }
// }
//
// export async function DELETE({ request, url }: any): Promise<Response> {
//   if (!auth(request)) {
//     return new Response(null, {
//       status: 401,
//     });
//   }
// }

function toKebab(s: string): string {
  return s
    .toLowerCase()
    .replaceAll(" ", "-")
    .split("")
    .filter((s: string) => s.match(/^[a-z0-9-]+$/i))
    .join("");
}
