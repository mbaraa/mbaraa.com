package tmplrndr

import "io"

type IndexProps struct {
	Name  string
	Brief string
}

type indexTemplate struct{}

// NewIndex returns a new index template instance.
func NewIndex() Template[IndexProps] {
	return &indexTemplate{}
}

func (i *indexTemplate) Render(props IndexProps) io.Reader {
	// error is ignored, because it's impossible to happen!
	out, _ := renderTemplate("index", props)
	return out
}
