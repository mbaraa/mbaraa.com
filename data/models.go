package data

type Blog struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	PublicId    string `json:"public_id"`
	VisitTimes  uint   `json:"visited_times"`
	WrittenAt   int64  `json:"written_at"`
}

type ContactLink struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Target   string `json:"target"`
	IconPath string `json:"icon_path"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LogoUrl     string `json:"logo_url"`
	SourceCode  string `json:"source_code"`
	Website     string `json:"website"`
	StartYear   string `json:"start_year"`
	EndYear     string `json:"end_year"`
}

type ProjectGroup struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Order       uint      `json:"order"`
	Projects    []Project `json:"projects"`
}

type WorkExperience experience

type VolunteeringExperience experience

type experience struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	StartDate   int64    `json:"start_date"`
	EndDate     int64    `json:"end_date"`
	Roles       []string `json:"roles"`
}

type Info struct {
	FullAbout    string   `json:"full_about"`
	BriefAbout   string   `json:"brief_about"`
	BlogIntro    string   `json:"blog_intro"`
	Technologies []string `json:"technologies"`
}
