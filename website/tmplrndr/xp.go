package tmplrndr

import (
	"bytes"
	"io"
	"mbaraacom/log"
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
	out, err := renderTemplate("xp", props)
	if err != nil {
		log.Errorln(err)
		return bytes.NewBuffer([]byte{})
	}
	return out
}
