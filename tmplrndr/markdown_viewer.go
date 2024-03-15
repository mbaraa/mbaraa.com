package tmplrndr

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma"
	chromahtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"

	"io"
)

var (
	htmlFormatter  *chromahtml.Formatter
	highlightStyle *chroma.Style
)

// the highlighter code was taken from this blog post https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html

func init() {
	htmlFormatter = chromahtml.New(chromahtml.WithClasses(true), chromahtml.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}
	styleName := "base16-snazzy"
	highlightStyle = styles.Get(styleName)
	if highlightStyle == nil {
		panic(fmt.Sprintf("didn't find style '%s'", styleName))
	}
}

// based on https://github.com/alecthomas/chroma/blob/master/quick/quick.go
func htmlHighlight(w io.Writer, source, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return htmlFormatter.Format(w, highlightStyle, it)
}

func renderCode(w io.Writer, codeBlock *ast.CodeBlock, _ bool) {
	defaultLang := ""
	lang := string(codeBlock.Info)
	htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(w, code, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

type MarkdownViewerProps struct {
	Markdown string
}

type markdownViewerTemplate struct{}

func NewMarkdownViewer() Template[MarkdownViewerProps] {
	return &markdownViewerTemplate{}
}

func (m *markdownViewerTemplate) Render(props MarkdownViewerProps) io.Reader {
	opts := html.RendererOptions{
		Flags:          html.CommonFlags | html.TOC | html.LazyLoadImages | html.UseXHTML,
		RenderNodeHook: myRenderHook,
	}

	buf := bytes.NewBuffer([]byte{})
	defer buf.Reset()

	buf.WriteString("<style scoped>")
	_ = htmlFormatter.WriteCSS(buf, highlightStyle)
	buf.WriteString("</style>")

	renderer := html.NewRenderer(opts)
	buf.Write(markdown.ToHTML([]byte(props.Markdown), nil, renderer))
	props.Markdown = buf.String()

	out, _ := renderTemplate("markdown-viewer", props)
	return out
}
