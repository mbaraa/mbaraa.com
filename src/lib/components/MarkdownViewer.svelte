<script lang="ts">
	import MarkdownIt from "markdown-it";
	import hljs from "highlight.js";
	import "highlight.js/styles/base16/solarized-dark.css";

	export let text: string;

	const markdown = new MarkdownIt({
		highlight: function (str, lang) {
			return hljs.highlightAuto(str).value;
		}
	});

	const style = `<style scoped>
.markdown table,
tr,
td,
th {
	padding: 10px;
	border: 2px #dcdcdc solid;
}
.markdown a {
	color: #20db8f;
	text-decoration-line: underline;
}

.markdown a:hover {
	text-decoration-line: none;
	color: #10ca7e;
	mouse-cursor: pointer;
}
.markdown ol {
	list-style-type: upper-roman;
    margin-left: 12px;
}
.markdown ul {
	list-style-type: circle;
    margin-left: 12px;
}
.markdown h1 {
	display: block;
	font-size: 2em;
	font-weight: bold;
	margin-block-start: 0.67em;
	margin-block-end: 0.67em;
}
.markdown h2,
:-moz-any(article, aside, nav, section) h1 {
	display: block;
	font-size: 1.5em;
	font-weight: bold;
	margin-block-start: 0.83em;
	margin-block-end: 0.83em;
}
.markdown h3,
:-moz-any(article, aside, nav, section) :-moz-any(article, aside, nav, section) h1 {
	display: block;
	font-size: 1.17em;
	font-weight: bold;
	margin-block-start: 1em;
	margin-block-end: 1em;
}
.markdown h4,
:-moz-any(article, aside, nav, section)
	:-moz-any(article, aside, nav, section)
	:-moz-any(article, aside, nav, section)
	h1 {
	display: block;
	font-size: 1em;
	font-weight: bold;
	margin-block-start: 1.33em;
	margin-block-end: 1.33em;
}
.markdown h5,
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
.markdown h6,
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
.markdown p {
	line-height: 1.9;
}
.markdown pre {
	background-color: #282828;
	border-radius: 8px;
	padding: 10px;
	color: white;
	line-height: 1.6;
    width: inherit;
    overflow-x: scroll;
}
.markdown code {
	background-color: #282828;
	border-radius: 8px;
	padding: 3px;
	color: white;
    width: inherit;
    overflow-x: scroll;
}
</style>`;
	const copyScript = `function copyToClipboard(str) {
            let textArea = document.getElementById("dummyText");
            textArea.hidden = false;
            textArea.value = str;
            textArea.select();
            textArea.setSelectionRange(0, 99999);
            document.execCommand("copy");

            textArea.hidden = true;
        }`;

	function addCopyToCode(html: string): string {
		return html.replaceAll("<code", "<code copy");
	}
</script>

<div>
	{@html `${style} <div class="markdown">${markdown.render(text)}</div>`}
</div>
