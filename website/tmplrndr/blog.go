package tmplrndr

import (
	"bytes"
	"io"
	"mbaraacom/log"
)

type BlogPostProps struct {
	BlogPostPreview
	Content string
}

type blogPostTemplate struct{}

func NewBlogPost() Template[BlogPostProps] {
	return &blogPostTemplate{}
}

func (a *blogPostTemplate) Render(props BlogPostProps) io.Reader {
	md := NewMarkdownViewer().Render(MarkdownViewerProps{
		Markdown: props.Content,
	})
	buf := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(buf, md)
	props.Content = buf.String()
	buf.Reset()

	out, err := renderTemplate("blog", props)
	if err != nil {
		log.Errorln(err)
		return bytes.NewBuffer([]byte{})
	}
	return out
}
