import auth from "$lib/auth";
import db from "$lib/db";
import type Blog from "$lib/models/Blog";
import {v4 as uuidv4} from "uuid";

export async function GET({url}: any): Promise<Response> {
    const id = url.searchParams.get("id")
    if (id) {
        const blog = await db.blog.findFirst({
            where: {
                id: id,
            }
        })
        if (!blog) {
            return new Response(null, {status: 404});
        }

        await db.blog.update({
            where: {
                id: id,
            },
            data: {
                readTimes: blog.readTimes + 1,
            }
        });

        return new Response(JSON.stringify(blog))
    }

    const blogs = (await db.blog.findMany())
        .map(blog => {
            blog.content = ""
            return blog
        })
        .sort((blogI: Blog, blogJ: Blog) => {
            return -(blogI.createdAt.getTime() - blogJ.createdAt.getTime());
        });
    return new Response(JSON.stringify(blogs))
}

export async function POST({request}: any): Promise<Response> {
    if (!auth(request)) {
        return new Response(null, {status: 401})
    }

    const blog = await request.json();


    blog.id = toKebab(blog.name);
    blog.readTimes = 0;

    const existingBlog = await db.blog.findFirst({where: {id: blog.id}});

    if (existingBlog) {
        blog.id += "-" + uuidv4().substring(0, 8)
    }

    const newBlog = await db.blog.create({
        data: blog,
    })

    if (!newBlog) {
        return new Response(null, {status: 500});
    }

    return new Response(null, {status: 200});
}

export async function PUT({request}: any): Promise<Response> {
    if (!auth(request)) {
        return new Response(null, {
            status: 401
        })
    }

    const blog = await request.json();

    const updatedBlog = await db.blog.update({
        where: {
            id: blog.id,
        },
        data: blog,
    })

    if (!updatedBlog) {
        return new Response(null, {status: 500});
    }

    return new Response(null, {status: 200});
}

export async function DELETE({request, url}: any): Promise<Response> {
    if (!auth(request)) {
        return new Response(null, {
            status: 401
        })
    }
    const id = url.searchParams.get("id")

    const deletedBlog = await db.blog.delete({
        where: {
            id: id,
        },
    })

    if (!deletedBlog) {
        return new Response(null, {status: 500});
    }

    return new Response(null, {status: 200});
}

function toKebab(s: string): string {
    return s
        .toLowerCase()
        .replaceAll(" ", "-")
        .split("")
        .filter((s: string) => s.match(/^[a-z0-9-]+$/i))
        .join("");
}
