import type Blog from "$lib/models/Blog";
import Requests from "./Requests";

export default class BlogRequests {
    static async newBlog(blog: Blog) {
        return Requests.makeAuthRequest("POST", "blog", blog, {}, {'Content-Type': 'application/json'})
    }

    static async updateBlog(blog: Blog) {
        return Requests.makeAuthRequest("PUT", "blog", blog, {}, {'Content-Type': 'application/json'})
    }

    static async deleteBlog(id: string) {
        return Requests.makeAuthRequest("DELETE", `blog/${id}`, null)
    }

    static async getBlogs(): Promise<Blog[]> {
        return Requests.makeRequest("GET", "blog", null)
            .then(resp => resp.json())
            .then(blogs => blogs)
    }

    static async getBlog(id: string): Promise<Blog> {
        return Requests.makeRequest("GET", `blog/${id}`, null)
            .then(resp => resp.json())
            .then(blogs => blogs as Blog)
    }
}
