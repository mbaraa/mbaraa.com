package tmplrndr

import (
	"bytes"
	"internal/log"
	"io"
)

type BlogPostPreview struct {
	Title       string
	Description string
	PublicId    string
	VisitTimes  uint
	WrittenAt   string
}

type BlogsProps struct {
	BlogIntro string
	Blogs     []BlogPostPreview
}

type blogsTemplate struct{}

func NewBlogs() Template[BlogsProps] {
	return &blogsTemplate{}
}

func (a *blogsTemplate) Render(props BlogsProps) io.Reader {
	out, err := renderTemplate("blogs", props)
	if err != nil {
		log.Errorln(err)
		return bytes.NewBuffer([]byte{})
	}
	return out
}
