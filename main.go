package main

import (
	"embed"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/mbaraa/mbaraa.com/config"
	"github.com/mbaraa/mbaraa.com/data"
	"github.com/mbaraa/mbaraa.com/log"
	"github.com/mbaraa/mbaraa.com/tmplrndr"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

var (
	//go:embed resources/*
	res embed.FS
)

func main() {
	timer := time.NewTicker(time.Hour)
	go func() {
		for range timer.C {
			err := data.UpdateBlogsMeta()
			if err != nil {
				log.Errorln(err)
			}
		}
	}()

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	appHandler := http.NewServeMux()
	appHandler.HandleFunc("/", handleHomePage)
	appHandler.HandleFunc("/projects", handleProjectsPage)
	appHandler.HandleFunc("/xp", handleXpPage)
	appHandler.HandleFunc("/about", handleAboutPage)
	appHandler.HandleFunc("/blog", handleBlogsPage)
	appHandler.HandleFunc("/blog/", handleBlogPostPage)
	// the blogs' images thing
	appHandler.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir(config.Config().FilesDir+"/images/"))))

	appHandler.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infof("website's server started at port %s\n", config.Config().WebsitePort)
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":"+config.Config().WebsitePort, m.Middleware(appHandler)))

	timer.Stop()
}

func handleErrorPage(w http.ResponseWriter, r *http.Request) {
	log.Errorf("Error happened when calling: %s %s?%s\n", r.Method, r.URL.Path, r.URL.Query().Encode())
	page := tmplrndr.NewError().Render(tmplrndr.ErrorProps{})
	_, _ = io.Copy(w, page)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "robots.txt") {
		robotsFile, _ := res.ReadFile("resources/robots.txt")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
		return
	}
	if strings.Contains(r.URL.Path, "sitemap.xml") {
		sitemapFile, _ := res.ReadFile("resources/sitemap.xml")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(sitemapFile)
		return
	}

	info, err := data.GetInfo()
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
		return
	}
	page := tmplrndr.NewIndex().Render(tmplrndr.IndexProps{
		Name:  "Baraa Al-Masri",
		Brief: info.BriefAbout,
	})
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
	}
}

func handleProjectsPage(w http.ResponseWriter, r *http.Request) {
	pgs, err := data.GetProjectGroups()
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
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
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
	}
}

func handleXpPage(w http.ResponseWriter, r *http.Request) {
	work, err := data.GetWorkXP()
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
		return
	}
	vol, err := data.GetVolunteeringXP()
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
		return
	}

	workXpView := tmplrndr.ExperienceGroup{
		Name: "ProfessionalWork",
		Xps:  nil,
	}
	for _, work := range work {
		startYear := ""
		if work.StartDate != 0 {
			startYear = time.Unix(work.StartDate, 0).Format("Jan/2006")
		}
		endYear := ""
		if work.EndDate != 0 {
			endYear = time.Unix(work.EndDate, 0).Format("Jan/2006")
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
			startYear = time.Unix(vol.StartDate, 0).Format("Jan/2006")
		}
		endYear := ""
		if vol.EndDate != 0 {
			endYear = time.Unix(vol.EndDate, 0).Format("Jan/2006")
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
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
	}
}

func handleAboutPage(w http.ResponseWriter, r *http.Request) {
	info, err := data.GetInfo()
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
		return
	}

	page := tmplrndr.NewAbout().Render(tmplrndr.AboutProps{
		PrerenderedMarkdown: info.FullAbout,
		Technologies:        info.Technologies,
	})
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, page)
	if err != nil {
		handleErrorPage(w, r)
		log.Errorln(err)
	}
}

func handleBlogsPage(w http.ResponseWriter, r *http.Request) {
	blogs, err := data.GetBlogs()
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
		return

	}
	info, err := data.GetInfo()
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
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
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, page)
	if err != nil {
		handleErrorPage(w, r)
		log.Errorln(err)
	}
}

func handleBlogPostPage(w http.ResponseWriter, r *http.Request) {
	blogId := r.URL.Path[len("/blog/"):]
	blog, err := data.GetBlogByPublicId(blogId)
	if err != nil {
		log.Errorln(err)
		handleErrorPage(w, r)
		return
	}

	err = data.IncrementBlogReads(blog.PublicId)
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
	w.Header().Set("Content-Type", "text/html")
	_, err = io.Copy(w, page)
	if err != nil {
		handleErrorPage(w, r)
	}
}
