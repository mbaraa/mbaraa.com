package tmplrndr

import (
	"io"
)

type ProjectGroup struct {
	Title       string
	Description string
	Order       int
	PublicId    string
	Projects    []struct {
		Name        string
		Description string
		LogoUrl     string
		SourceCode  string
		Website     string
		StartYear   string
		EndYear     string
		ComingSoon  bool
	}
}

type ProjectsProps struct {
	Groups []ProjectGroup
}

type projectsTemplate struct{}

func NewProjects() Template[ProjectsProps] {
	return &projectsTemplate{}
}

func (p *projectsTemplate) Render(props ProjectsProps) io.Reader {
	for i, group := range props.Groups {
		for j, project := range group.Projects {
			props.Groups[i].Projects[j].ComingSoon = project.SourceCode == "" && project.Website == ""
		}
	}
	out, _ := renderTemplate("projects", props)
	return out
}
