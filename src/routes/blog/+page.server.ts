import type { Load } from "@sveltejs/kit";
import type Blog from "$lib/models/Blog";
import { db } from "$lib/db";

export const ssr = true;
export const prerender = true;

export const load: Load = async () => {
  const blogs = (await getAll())
    .map((blog) => {
      blog.content = "";
      return blog;
    })
    .sort((blogI: Blog, blogJ: Blog) => {
      return -(blogI.createdAt.getTime() - blogJ.createdAt.getTime());
    });
  return { blogs: blogs };
};

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
