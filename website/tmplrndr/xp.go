package tmplrndr

import (
	"io"
)

type Experience struct {
	Title       string
	Description string
	Location    string
	StartDate   string
	EndDate     string
	Roles       []string
}

type ExperienceGroup struct {
	Name string
	Xps  []Experience
}

type XPsProps struct {
	ProfessionalWork ExperienceGroup
	Volunteering     ExperienceGroup
}

type xpsTemplate struct{}

func NewXPs() Template[XPsProps] {
	return &xpsTemplate{}
}

func (p *xpsTemplate) Render(props XPsProps) io.Reader {
	out, _ := renderTemplate("xp", props)
	return out
}
