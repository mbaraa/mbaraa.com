import auth from "$lib/auth";
import type Blog from "$lib/models/Blog";
import { v4 as uuidv4 } from "uuid";
import { db } from "$lib/db";
import { RequestEvent, type RequestHandler } from "@sveltejs/kit";

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
      return -(blogI.createdAt.getTime() - blogJ.createdAt.getTime());
    });

  return new Response(JSON.stringify(blogs));
}

export const POST: RequestHandler = async (ev: RequestEvent) => {
  if (!auth(ev.request)) {
    return new Response(null, { status: 401 });
  }

  const blog: Blog = await ev.request.json();

  const status = await insert(blog);

  if (!status) {
    return new Response(null, { status: 500 });
  }

  return new Response(null);
};

export const PUT: RequestHandler = async (ev: RequestEvent) => {
  if (!auth(ev.request)) {
    return new Response(null, { status: 401 });
  }

  const blog: Blog = await ev.request.json();

  const status = await update(blog.publicId as string, blog);

  if (!status) {
    return new Response(null, { status: 500 });
  }

  return new Response(null);
};

export const DELETE: RequestHandler = async (ev: RequestEvent) => {
  if (!auth(ev.request)) {
    return new Response(null, { status: 401 });
  }

  const status = await delete_(ev.url.searchParams.get("id") as string);

  if (!status) {
    return new Response(null, { status: 500 });
  }

  return new Response(null);
};

async function get(id: string): Promise<unknown> {
  const blogsRef = db.collection("blogs");
  const snapshot = await blogsRef.where("publicId", "==", id).get();

  if (snapshot.empty) {
    return null;
  }

  let blog: Blog | undefined = undefined;

  snapshot.forEach((doc) => {
    blog = doc.data() as Blog;
    blog.createdAt = doc.data().createdAt.toDate();
    blog.updatedAt = doc.data().updatedAt.toDate();
  });

  return blog;
}

async function getAll(): Promise<Blog[]> {
  const blogsRef = db.collection("blogs");
  const snapshot = await blogsRef.get();

  let blogs: Blog[] = [];

  snapshot.forEach((doc) => {
    const blog = doc.data() as Blog;
    blog.createdAt = doc.data().createdAt.toDate();
    blog.updatedAt = doc.data().updatedAt.toDate();
    blogs.push(blog);
  });

  return blogs;
}

async function insert(blog: Blog): Promise<unknown> {
  blog.publicId = toKebab(blog.name);
  blog.createdAt = new Date();
  blog.updatedAt = new Date();

  const document = db.doc(`blogs/${blog.publicId}`);
  const status = await document.set(blog);

  if (!status) {
    return null;
  }

  return blog;
}

async function update(id: string, blog: Blog): Promise<unknown> {
  const document = db.doc(`blogs/${blog.publicId}`);
  const status = await document.update({
    content: blog.content,
    description: blog.description,
    name: blog.name,
    updatedAt: new Date(),
  });

  if (!status) {
    return null;
  }

  return blog;
}

async function delete_(id: string): Promise<unknown> {
  const document = db.doc(`blogs/${id}`);
  return await document.delete();
}

function toKebab(s: string): string {
  return s
    .toLowerCase()
    .replaceAll(" ", "-")
    .split("")
    .filter((s: string) => s.match(/^[a-z0-9-]+$/i))
    .join("");
}
