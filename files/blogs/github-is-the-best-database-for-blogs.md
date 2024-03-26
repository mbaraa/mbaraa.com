If you're reading this, well, this website itself is the cause of me writing this post!

So, let's begin!

### How Much Rewrites Did This Website Witness?

Four, four rewrites for a simple portfolio and blog, why? you might ask, I'll give you a very long and detailed answer.

#### First write (early web days)

When I wrote a portfolio for the first time, it was because I bought the domain `mbaraa.fun` so I made a [Vue](https://vuejs.org/) + [Vutify](https://vuetifyjs.com/en/) back in 2021 when I was setting up my resume using a template that had a website placeholder, so I thought, "Well, I'm a web developer, why not make a portfolio website?", and I made something that did the job, and I kept it on, until the domain expired.

#### First rewrite (Next, Vercel and PlanetScale)

After some while, I was doing the first rewrite, and I took the thing a bit more serious, that I bought the domain `mbaraa.com`, no fun anymore. and wrote my portfolio with [Next.js](https://nextjs.org/), deployed it using [Vercel](https://vercel.com), and the database was hosted on [PlanetScale](https://planetscale.com/), and this whole mumbo jumbo of a stack is too much for a simple portfolio and blog, it's like strapping a nuclear reactor into a bike. the site stayed as a portfolio for a while, until I had the idea of making a blog, which was really horrible, it didn't have a dashboard, I used to edit the records manually via SQL queries, so yeah, I thought of making a fancy dashboard, with a fancy database.

#### Second rewrite (SvelteKit, Firebase, and my independent server)

I discovered [SvelteKit](https://kit.svelte.dev) in July/2022 and it was my goto frontend JS framework to prototype an idea or something that would be an overkill to use a fancy thing like [Yew](https://yew.rs), [htmx](https://htmx.org), or [templ](https://github.com/a-h/templ), since then, React was a bit too much, since I really like server side rendered stuff, that I have to do super client specific things, like scroll events, pop state, media, and other stuff...

SvelteKit gave me the satisfaction of both worlds, I get a type safeish frontend and a good performant server side rendering backend, I pretty much used SvelteKit for every project and tool, and to prototype ideas at work, so down this road, I decided to do a rewrite with it. I was still using Vercel for deployment, and same for PlanetScale (cuz data migration is a pain in the ass) so what I did, is that I dumped the database into [JSON](https://www.json.org/json-en.html) files, which must've been a sign to stop there, and to do this properly I moved my website to my server, which is a [Compute Engine](https://cloud.google.com/products/compute/?hl=en) VM, so that I could manage the JSON files easily on the server, one thing led to another and I went super fancy, out of the blue, I decided to use [FireStore](https://firebase.google.com/docs/firestore/) as my database, and I gotta tell you, IT WAS A STUPID IDIOTIC DECISION, why? well, when my aim was towards speed and SEO, FireStore kinda broke that (I think if I'd hosted the website on FireStore it would've been a different story), but when a request reaches the website's server, it makes another request to FireStore to fetch the data, and that really bottlenecked all requests that needed to fill a template from the database. I did do a pre-render, but I didn't like restarting the website's container every time I made a change, so I disabled it. I did add a data proxy on the server, to reduce FireStore calls, but again I didn't like it much, so my solution was to rewrite the data layer, since it was really easy to replace, but little I knew that data migration from FireStore was almost impossible on the database level, so, I setup some endpoints to return a JSON response of the wanted document's collection, or whatever it's called, I dumped the stuff into a [MongoDB](https://mongodb.com), thinking that I'll just use mongo where the database will be on the same server, so no network bottlenecks, no problem whatsoever, but no, half-way through, TypeScript really drove me crazy, with it's ambiguous errors, and mongo's weren't helping at all, so I thought a [Go](https://golang.org) rewrite will be it, and I'll just write it and that'll be it, since Go is the only language I can write code with, without looking up errors, or begging for a help on Stack Overflow.

#### Last rewrite (Go, and the GitHub dashboard)

I started this rewrite in January/2024 as some part of my extreme activity, that month I did the most contributions on GitHub, where I rewrote my website, created two projects and a half [GitHub Graph Drawer](https://github-graph-drawer.mbaraa.com), [htmx.pics](https://htmx.pics), and [SlideMD](https://slidemd.com). back to the website, somehow, someway, I thought, splitting the dashboard, and the website into Go's [Multi-module workspaces](https://go.dev/doc/tutorial/workspaces), where I can have some shared code, make it look fancy, and other stuff, but little I knew that I was gonna leave this dashboard behind, and get into a burnout phase, where the dashboard wasn't finished, and I wanted to write two blog posts, but I couldn't, since I didn't finish the dashboard, so like yesterday (use the post's date in the blogs page as a reference point), I thought of this great idea, which is "The best dashboard is no dashboard", and let me tell you, I finished implementing it in like 2 hours or something, what really pushed me into doing it, was the fact that I have created this tool [Rex](https://github.com/mbaraa/rex), that is a super minimal and fast, CI/CD thing, to deploy a GitHub (yes only GitHub, I didn't make actions for other Git things), so basically, when I push, It deploys the application into my server, which is super neat. I'm using the website's Git repository as a database, you can see for example. my blogs [here](https://github.com/mbaraa/mbaraa.com/tree/main/files/blogs), all that I have to do to make a change on my website now, or if you liked it and want to fork it, is to update the wanted file under [files](https://github.com/mbaraa/mbaraa.com/tree/main/files), and just push the changes, and Rex takes care of the rest for me.

So for now I'm gonna stick with this solution, until I go crazy and do something else, that will waste my time again.

### Conclusion from Each Rewrite

Each rewrite did teach me a valuable lesson, and I did learn a lot along the way. when I first made my portfolio, I had been writing web for 7 months, so I was really new into the stuff, the last rewrite I had been a web developer for 3 years, 2 of them I worked with professionals, where I learned so much stuff, when I saw stuff from a different POV, and given the fact, that this is a side project, and had a little focus from me, and it got the most wisdom, since, apparently, whenever I take a break, I make a modification to it!

### GitHub is the Best Database for Blogs

So why did I say that GitHub is the best database for blogs, well, let me tell you a little story, a personal blog is a program that's edited by a single person, hence the name "Personal Blog", so as long as there's only a single person modifying data for the site, and, yes, I did think of using a CMS to manage all that stuff for me, but it's stupid for my case, since I'M A SINGLE PERSON MODIFYING MY WEBSITE'S DATA, and most CMSs take a lot of resources, and this is an overkill for a portfolio and blog, FYI the Go binary running my website is only utilizing 8 MiB, and that should be enough reason, why I'm not using a CMS here, and here's a screenshot of the memory utilization from `docker stats`

![mbaraa.com memory usage](/img/mbaraacom_memory_usage.png "mbaraa.com memory usage from docker stats")
And for the record, this will be the first blog post, that I'll publish using this GitHub dashboard thingy.

The funny thing, is that, I could've used GitHub pages, and just slapped markdown files, but I wanted my website to still look the same, and there are pages that are not juts markdown files, such as [projects](https://mbaraa.com/projects), and I had the [Sunk Cost Fallacy](https://www.grammarly.com/blog/sunk-cost-fallacy/), where, I wrote a fancy wrapper for Go's html templates, and had all those organized components and templates, that I didn't want to replace it with boring markdown pages, and also to flex the website's speed, which is BLAZINGLY FAST, well, the Go server doesn't do much to render a page (since all the site's data are in memory :)), but, eh, if it's fast, it's fast.

### What About Multiple Writers?

Actually, this can be managed easily, like any other Git repository, a maintainer(s) will manage what to be merged into production, so unless you're hoping to build something like [devto](https://dev.to) or [Medium](https://medium.com/), GitHub with Rex is your best choice.

#### Team Structure (hypothetical team)

1. Lead editor, or in GitHub terms, repository owner:
   - This person's responsibility is to either directly write a blog post and push it to the repository, or accept (merge) other people's posts into the main branch, and do some administration on the repository.
2. Core editors, or in GitHub terms, repository collaborator:
   - Those have the same power as the Lead, with less privileges that the Lead cuts, so their job, is only to write, or accept other people's work.
3. Outside editors, or in GitHub terms, outside collaborator:
   - Those people suggest a blog post, and the Lead or a Core Editor accepts the post, or declines it.

This whole GitHub dashboard took me 2 hours to write (excluding the templates writing time), so it's doable, and it saves resources like crazy, and actually, I'm thinking of doing this for a couple of other places where I have authority on.

#### Open Data?

Well, since I'm an opensource supporter and contributer, and that I work for [JOSA](https://josa.ngo) where we promote opensource software, and how valuable your data is.

This method of hosting a blog is as open as data can be, you can literally see what data you're reading, and how your click is processed on the site [here's the repo's link if you're too lazy to click my GitHub profile on the bottom](https://github.com/mbaraa/mbaraa.com), and if you open the network inspector, you'll see that no fishy data is going in and out, and again the website is LITERALLY the same as in GitHub, nothing is hidden, nothing get recorded without you knowing.

If you're really interested in this open data thing, email me at [pub@mbaraa.com](mailto:pub@mbaraa.com), we could start a campaign or something.

### Simplicity is Really Underrated These Days

As far as I'm seeing, people are going toward complexity (I was going for complexity), but I think it's time to sit down, and think about how much people will edit content on a website, and within what period of time, the smaller the number of people, and the larger the number of separating time periods, the less complex the project should be, like a blog for example, why pay Medium whatever amount of money, to get "premium features", or host a giga chungus wordpress, or use Strapi with Next.js, and whole Infrastructure, that's serving 50 people a day, simplicity is always the key, where you could get those SEO numbers, by learning advanced web topics, and stop just slapping text into a file, until a page appears!

I have killed over 30 projects, because their complexity chunked up and really burnt me out, prototypes always help demonstrate the size of the project, the expected concurrent users, and most importantly, is it worth it being a big chungus, fancy domain-hex-multi-module-monorepo project, well, I guess that's it, thanks for sticking to the end.

### Quote of the day

"No matter how many people give me advice, I am going to do what my heart tells me to do."
\- [Lana Del Rey](https://en.wikipedia.org/wiki/Lana_Del_Rey)
