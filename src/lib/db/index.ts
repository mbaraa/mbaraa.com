import type Blog from "$lib/models/Blog";
import { Firestore } from "@google-cloud/firestore";

const db = new Firestore({
  projectId: "",
  keyFilename: "./firebase-key.json",
});

export async function get(id: string): Promise<unknown> {
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

export async function getAll(): Promise<Blog[]> {
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
