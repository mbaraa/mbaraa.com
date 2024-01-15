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
					Roles:       []string{"Paris", "Texas"},
				},

				{
					Title:       "ProgressSoft",
					Description: "JAVAAAAAa",
					Location:    "Amman, Jordan",
					StartDate:   "2022",
					EndDate:     "2023",
					Roles:       []string{"Margaret"},
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
					Roles:       []string{"Taco Truck", "Venice Bitch"},
				},
			},
		},
	})
	_, err := io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
	}
}

func handleAboutPage(w http.ResponseWriter, r *http.Request) {
	page := tmplrndr.NewAbout().Render(tmplrndr.AboutProps{
		PrerenderedMarkdown: "Hello there! My name is Baraa Al-Masri and I like building software, most likely web apps, some of the stuff I build are opensource, here's a list of my [projects](/projects) that saw the light of day.\n\nMy interst in computers started when I was a kid when my computer broke down and I was like, hell I should be able to fix that, and so I've completed this dark path, where I was to fix computers and phones of my family's.\n\nI started my Linux journey when I was 13 where we had a computer course, and in the OS section Ubuntu was mentioned, I thought well I should try that on my computer, one thing led to another and I use Gentoo now :)\n\nI'm currently studying Software Engineering in my collage, because I thought collage would give me what I wanted, but it wasn't enough for me so I started building useless daily apps, till I found  myself in web development, and so I stayed in that area.\n\nAfter trying different languages and frameworks, two of them really hooked me in, Go was the best programming language I've ever used concidering it's simplicity, cleanliness, and rich standard library, so I kept using it for backend and system applications, I've tried a lot of frontend frameworks like, Go templates, Vue.js, React.js, Next.js and the GOAT SvelteKit which became my favourite frontend framework because of its clean and light architecture for an SSR frontend framework.\n\nI started my professional work at ProgressSoft as a Java SpringBoot developer, before that I used to build web apps for my university students, and I used to work as a freelancer since I was 2nd year in collage, most of the stuff I made were proprietary software, so here's that.",
		Technologies: []string{
			"Go",
			"Rust",
			"TypeScript",
			"Java",
			"Elixir",
			"Bash",
			"SvelteKit",
			"Nuxt",
			"Yew",
			"SpringBoot",
			"Phoenix",
			"Linux",
			"Google Cloud Platform",
			"Embedded Programming",
			"Vim",
			"LaTex",
			"Git",
			"Nginx",
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
	http.HandleFunc("/about", handleAboutPage)

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infoln("server started at port 3000")
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":3000", nil))
}
