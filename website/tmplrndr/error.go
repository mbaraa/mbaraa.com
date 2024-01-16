package tmplrndr

import "io"

type ErrorProps struct {
}

type errorTemplate struct{}

// NewError returns a new error template instance.
func NewError() Template[ErrorProps] {
	return &errorTemplate{}
}

func (i *errorTemplate) Render(props ErrorProps) io.Reader {
	// error is ignored, because it's impossible to happen!
	out, _ := renderTemplate("error", props)
	return out
}
