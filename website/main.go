package main

import (
	"embed"
	"fmt"
	"io"
	"mbaraacom/log"
	"mbaraacom/tmplrndr"
	"net/http"
	"time"
)

var (
	//go:embed resources/*
	res embed.FS
)

func handelErrorPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	page := tmplrndr.NewIndex().Render(tmplrndr.IndexProps{
		Name:  "Baraa Al-Masri",
		Brief: "I'm a software developer specializing in web development in various stacks, and a fresh embedded rustacean 🦀　 I pay rent by writing TypeScript full stack web apps @ Jordan Open Source Association.　 And in my free time I write more code, blog, and slack watching YT shorts.",
	})
	_, err := io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
	}
}

func handleProjectsPage(w http.ResponseWriter, r *http.Request) {
	page := tmplrndr.NewProjects().Render(tmplrndr.ProjectsProps{
		Groups: []tmplrndr.ProjectGroup{
			{
				Title:       "Bla Bla Bla",
				Description: "Blyaaaaaaaaaaaaaaaaaaaaaat",
				Order:       1,
				PublicId:    "bla-bla-bla",
				Projects: []struct {
					Name        string
					Description string
					LogoUrl     string
					SourceCode  string
					Website     string
					StartYear   string
					EndYear     string
					ComingSoon  bool
				}{
					{
						Name:        "Mbaraa.com",
						Description: "Website blyat",
						LogoUrl:     "https://mbaraa.com",
						SourceCode:  "https://github.com/mbaraa/madebybaraa",
						Website:     "https://mbaraa.com",
						StartYear:   "2021",
						EndYear:     "2023",
						ComingSoon:  false,
					},
					{
						Name:        "Mbaraa.com",
						Description: "Website blyat",
						LogoUrl:     "https://mbaraa.com",
						Website:     "https://mbaraa.com",
						StartYear:   "2021",
						EndYear:     "2023",
						ComingSoon:  false,
					},
					{
						Name:        "Mbaraa.com",
						Description: "Website blyat",
						LogoUrl:     "https://mbaraa.com",
						StartYear:   "2021",
						EndYear:     "2023",
						ComingSoon:  false,
					},
				},
			},
			{
				Title:       "Bla Bla Bla 2",
				Description: "Blyaaaaaaaaaaaaaaaaaaaaaat 2",
				Order:       1,
				PublicId:    "bla-bla-bla",
				Projects: []struct {
					Name        string
					Description string
					LogoUrl     string
					SourceCode  string
					Website     string
					StartYear   string
					EndYear     string
					ComingSoon  bool
				}{
					{
						Name:        "Mbaraa.com",
						Description: "Website blyat",
						LogoUrl:     "https://mbaraa.com",
						SourceCode:  "https://github.com/mbaraa/madebybaraa",
						Website:     "https://mbaraa.com",
						StartYear:   "2021",
						EndYear:     "2023",
						ComingSoon:  false,
					},
					{
						Name:        "Mbaraa.com",
						Description: "Website blyat",
						LogoUrl:     "https://mbaraa.com",
						Website:     "https://mbaraa.com",
						StartYear:   "2021",
						EndYear:     "2023",
						ComingSoon:  false,
					},
					{
						Name:        "Mbaraa.com",
						Description: "Website blyat",
						LogoUrl:     "https://mbaraa.com",
						StartYear:   "2021",
						EndYear:     "2023",
						ComingSoon:  false,
					},
				},
			},
		},
	})
	_, err := io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
	}
}

func handleXpPage(w http.ResponseWriter, r *http.Request) {
	page := tmplrndr.NewXPs().Render(tmplrndr.XPsProps{
		ProfessionalWork: tmplrndr.ExperienceGroup{
			Name: "Professional Work",
			Xps: []tmplrndr.Experience{
				{
					Title:       "ProgressSoft",
					Description: "JAVAAAAAa",
					Location:    "Amman, Jordan",
					StartDate:   "2022",
					EndDate:     "2023",
					Roles:       []string{"Nigga", "Paris"},
				},

				{
					Title:       "ProgressSoft",
					Description: "JAVAAAAAa",
					Location:    "Amman, Jordan",
					StartDate:   "2022",
					EndDate:     "2023",
					Roles:       []string{"Nigga", "Paris"},
				},
			},
		},
		Volunteering: tmplrndr.ExperienceGroup{
			Name: "Volunteering",
			Xps: []tmplrndr.Experience{
				{
					Title:       "ProgressHard",
					Description: "JAVAAAAAa",
					Location:    "Amman, Jordan",
					StartDate:   "2022",
					EndDate:     "2023",
					Roles:       []string{"Nigga", "Paris"},
				},
			},
		},
	})
	_, err := io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
	}
}

func main() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/projects", handleProjectsPage)
	http.HandleFunc("/xp", handleXpPage)

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infoln("server started at port 3000")
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":3000", nil))
}
