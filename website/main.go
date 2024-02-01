package main

import (
	"embed"
	"fmt"
	"internal/db"
	"internal/log"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"internal/config"

	"github.com/mbaraa/mbaraa.com/website/tmplrndr"
)

var (
	//go:embed resources/*
	res embed.FS
)

func main() {
	db.Init(config.Config().DbUri)

	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/projects", handleProjectsPage)
	http.HandleFunc("/xp", handleXpPage)
	http.HandleFunc("/about", handleAboutPage)
	http.HandleFunc("/blog", handleBlogsPage)
	http.HandleFunc("/blog/", handleBlogPostPage)
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./_uploads/"))))

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infof("website's server started at port %s\n", config.Config().WebsitePort)
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":"+config.Config().WebsitePort, nil))
}

func handelErrorPage(w http.ResponseWriter, r *http.Request) {
	log.Errorf("Error happended when calling: %s %s?%s\n", r.Method, r.URL.Path, r.URL.Query().Encode())
	page := tmplrndr.NewError().Render(tmplrndr.ErrorProps{})
	_, _ = io.Copy(w, page)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "robots.txt") {
		robotsFile, _ := os.ReadFile("./resources/robots.txt")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
		return
	}

	info, err := db.GetInfo()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}
	page := tmplrndr.NewIndex().Render(tmplrndr.IndexProps{
		Name:  "Baraa Al-Masri",
		Brief: info.BriefAbout,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
	}
}

func handleProjectsPage(w http.ResponseWriter, r *http.Request) {
	pgs, err := db.GetProjectGroups()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}
	var viewGroups []tmplrndr.ProjectGroup
	for _, pg := range pgs {
		vg := tmplrndr.ProjectGroup{
			Title:       pg.Title,
			Description: pg.Description,
			Order:       int(pg.Order),
		}
		for _, project := range pg.Projects {
			vg.Projects = append(vg.Projects, struct {
				Name        string
				Description string
				LogoUrl     string
				SourceCode  string
				Website     string
				StartYear   string
				EndYear     string
				ComingSoon  bool
			}{
				Name:        project.Name,
				Description: project.Description,
				LogoUrl:     project.LogoUrl,
				SourceCode:  project.SourceCode,
				Website:     project.Website,
				StartYear:   project.StartYear,
				EndYear:     project.EndYear,
			})
		}
		viewGroups = append(viewGroups, vg)
	}

	page := tmplrndr.NewProjects().Render(tmplrndr.ProjectsProps{
		Groups: viewGroups,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
	}
}

func handleXpPage(w http.ResponseWriter, r *http.Request) {
	work, err := db.GetWorkXP()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}
	vol, err := db.GetVolunteeringXP()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	workXpView := tmplrndr.ExperienceGroup{
		Name: "ProfessionalWork",
		Xps:  nil,
	}
	for _, work := range work {
		startYear := ""
		if work.StartDate != 0 {
			startYear = fmt.Sprint(time.Unix(work.StartDate, 0).Year())
		}
		endYear := ""
		if work.EndDate != 0 {
			endYear = fmt.Sprint(time.Unix(work.EndDate, 0).Year())
		}
		workXpView.Xps = append(workXpView.Xps, tmplrndr.Experience{
			Title:       work.Title,
			Description: work.Description,
			Location:    work.Location,
			StartDate:   startYear,
			EndDate:     endYear,
			Roles:       work.Roles,
		})
	}

	volXpView := tmplrndr.ExperienceGroup{
		Name: "Volunteering",
		Xps:  nil,
	}
	for _, vol := range vol {
		startYear := ""
		if vol.StartDate != 0 {
			startYear = fmt.Sprint(time.Unix(vol.StartDate, 0).Year())
		}
		endYear := ""
		if vol.EndDate != 0 {
			endYear = fmt.Sprint(time.Unix(vol.EndDate, 0).Year())
		}
		volXpView.Xps = append(volXpView.Xps, tmplrndr.Experience{
			Title:       vol.Title,
			Description: vol.Description,
			Location:    vol.Location,
			StartDate:   startYear,
			EndDate:     endYear,
			Roles:       vol.Roles,
		})
	}

	page := tmplrndr.NewXPs().Render(tmplrndr.XPsProps{
		ProfessionalWork: workXpView,
		Volunteering:     volXpView,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
	}
}

func handleAboutPage(w http.ResponseWriter, r *http.Request) {
	info, err := db.GetInfo()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	page := tmplrndr.NewAbout().Render(tmplrndr.AboutProps{
		PrerenderedMarkdown: info.FullAbout,
		Technologies:        info.Technologies,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
		log.Errorln(err)
	}
}

func handleBlogsPage(w http.ResponseWriter, r *http.Request) {
	blogs, err := db.GetBlogs()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return

	}
	info, err := db.GetInfo()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	blogViews := make([]tmplrndr.BlogPostPreview, 0)
	for _, blog := range blogs {
		blogViews = append(blogViews, tmplrndr.BlogPostPreview{
			Title:       blog.Title,
			Description: blog.Description,
			PublicId:    blog.PublicId,
			VisitTimes:  blog.VisitTimes,
			WrittenAt:   time.Unix(blog.WrittenAt, 0).Format("Jan 02, 2006"),
		})
	}

	page := tmplrndr.NewBlogs().Render(tmplrndr.BlogsProps{
		BlogIntro: info.BlogIntro,
		Blogs:     blogViews,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
		log.Errorln(err)
	}
}

func handleBlogPostPage(w http.ResponseWriter, r *http.Request) {
	blogId := r.URL.Path[len("/blog/"):]
	blog, err := db.GetBlogByPublicId(blogId)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	blog.VisitTimes++
	err = db.UpdateBlog(blog.Id, db.Blog{VisitTimes: blog.VisitTimes})
	if err != nil {
		log.Errorln(err)
	}

	page := tmplrndr.NewBlogPost().Render(tmplrndr.BlogPostProps{
		BlogPostPreview: tmplrndr.BlogPostPreview{
			Title:       blog.Title,
			Description: blog.Description,
			PublicId:    blog.PublicId,
			VisitTimes:  blog.VisitTimes,
			WrittenAt:   time.Unix(blog.WrittenAt, 0).Format("Jan 02, 2006"),
		},
		Content: blog.Content,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
	}
}
