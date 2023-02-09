<script lang="ts">
    import { page } from "$app/stores";
    import BlogRequests from "$lib/utils/requests/BlogRequests";
    import { onMount } from "svelte";
    import type Blog from "$lib/models/Blog";
    import MarkdownIt from "markdown-it";
    import hljs from "highlight.js"
    import "highlight.js/styles/base16/gruvbox-dark-hard.css"

  const markdown = new MarkdownIt({
    highlight: function (str, lang) {
          return hljs.highlightAuto(str).value
    }
  });

    let blog: Blog;

    onMount(async () => {
        blog = await BlogRequests.getBlog($page.params.id);
    });

    const style = `<style scoped>
        h1 {
              display: block;
                font-size: 2em;
                  font-weight: bold;
                    margin-block-start: .67em;
                      margin-block-end: .67em;
        }

        h2,
        :-moz-any(article, aside, nav, section)
        h1 {
              display: block;
                font-size: 1.5em;
                  font-weight: bold;
                    margin-block-start: .83em;
                      margin-block-end: .83em;
        }

        h3,
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        h1 {
              display: block;
                font-size: 1.17em;
                  font-weight: bold;
                    margin-block-start: 1em;
                      margin-block-end: 1em;
        }

        h4,
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        h1 {
              display: block;
                font-size: 1.00em;
                  font-weight: bold;
                    margin-block-start: 1.33em;
                      margin-block-end: 1.33em;
        }

        h5,
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        h1 {
              display: block;
                font-size: 0.83em;
                  font-weight: bold;
                    margin-block-start: 1.67em;
                      margin-block-end: 1.67em;
        }

        h6,
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        :-moz-any(article, aside, nav, section)
        h1 {
              display: block;
                font-size: 0.67em;
                  font-weight: bold;
                    margin-block-start: 2.33em;
                      margin-block-end: 2.33em;
        }

        p {
            line-height: 2.3;
        }

        pre {
            background-color: #282828;
            border-radius: 8px;
            padding: 10px;
            color: white;
            line-height: 1.6;
        }

        code {
            background-color: #282828;
            border-radius: 8px;
            padding: 3px;
            color: white;
        }

        :scope a {
            color: #20db8f;
            text-decoration-line: underline;
        }

        a:hover {
            text-decoration-line: none;
            color: #10ca7e;
            mouse-cursor: pointer;
        }
</style>`;

</script>

<svelte:head>
    <title>
        {blog?.name}
    </title>
    <meta property="og:title" content={blog?.name} />
</svelte:head>

{#if blog}
    <div
        class="font-[Vistol] absolute left-[50%] translate-x-[-50%] py-[50px] w-auto "
    >
        <h1
            class="text-white w-[100%] text-center text-[30px] md:text-[40px] font-[1000] "
        >
            {blog.name}
        </h1>
        <div
            class="p-[45px] text-[18px] m-[18px] bg-white rounded-[32px] w-[90vw]"
        >
        <div>
            {@html (style + markdown.render(blog.content))}
        </div>

        </div>
    </div>
{/if}
