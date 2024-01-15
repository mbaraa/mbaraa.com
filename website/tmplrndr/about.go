package tmplrndr

import (
	"bytes"
	"io"
	"math"
	"mbaraacom/log"
)

type AboutProps struct {
	PrerenderedMarkdown   string
	RenderedMarkdown      string
	Technologies          []string
	BigScreenTechnologies [][]string
}

type aboutTemplate struct{}

func NewAbout() Template[AboutProps] {
	return &aboutTemplate{}
}

func (a *aboutTemplate) Render(props AboutProps) io.Reader {
	sizePerChunk := int(math.Ceil(float64(len(props.Technologies)) / 3.))
	for i := 0; i < len(props.Technologies); i += sizePerChunk {
		chunk := make([]string, 0)
		for j := i; j < i+sizePerChunk; j++ {
			chunk = append(chunk, props.Technologies[j])
		}
		props.BigScreenTechnologies = append(props.BigScreenTechnologies, chunk)
	}

	md := NewMarkdownViewer().Render(MarkdownViewerProps{
		Markdown: props.PrerenderedMarkdown,
	})
	buf := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(buf, md)

	props.RenderedMarkdown = buf.String()
	buf.Reset()
	props.PrerenderedMarkdown = ""

	out, err := renderTemplate("about", props)
	if err != nil {
		log.Errorln(err)
		return bytes.NewBuffer([]byte{})
	}
	return out
}
