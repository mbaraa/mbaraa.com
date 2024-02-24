package tmplrndr

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"

	"io"
)

type MarkdownViewerProps struct {
	Markdown string
}

type markdownViewerTemplate struct{}

func NewMarkdownViewer() Template[MarkdownViewerProps] {
	return &markdownViewerTemplate{}
}

func (m *markdownViewerTemplate) Render(props MarkdownViewerProps) io.Reader {
	opts := html.RendererOptions{
		Flags: html.CommonFlags | html.TOC | html.LazyLoadImages,
	}
	renderer := html.NewRenderer(opts)
	props.Markdown = string(markdown.ToHTML([]byte(props.Markdown), nil, renderer))

	out, _ := renderTemplate("markdown-viewer", props)
	return out
}
