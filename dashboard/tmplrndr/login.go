package tmplrndr

import "io"

type LoginProps struct {
}

type loginTemplate struct{}

// NewLogin returns a new login template instance.
func NewLogin() Template[LoginProps] {
	return &loginTemplate{}
}

func (i *loginTemplate) Render(props LoginProps) io.Reader {
	// error is ignored, because it's impossible to happen!
	out, _ := renderTemplate("login", props)
	return out
}
