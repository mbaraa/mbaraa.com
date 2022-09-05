import type Link from "$lib/models/Link";

export default class LinksRequests {
    static async getLinks(): Promise<Link[]> {
        return [
            {name: "Blog", link: "/blog", target: ""},
            {name: "GitHub", link: "https://github.com/mbaraa"},
            {name: "Twitter", link: "https://twitter.com/mbaraa271"},
            {name: "LinkedIn", link: "https://linkedin.com/in/-mbaraa-"},
        ]
    }
}


