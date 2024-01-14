package tmplrndr

import "io"

// IndexProps it's just here, maybe it'll be of use in the future.
type IndexProps struct{}

type indexTemplate struct{}

// NewIndex returns a new index template instance.
func NewIndex() Template[IndexProps] {
	return &indexTemplate{}
}

func (i *indexTemplate) Render(_ IndexProps) io.Reader {
	// error is ignored, because it's impossible to happen!
	out, _ := renderTemplate("index", IndexProps{})
	return out
}
