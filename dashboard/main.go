package main

import (
	"embed"
	"internal/db"
	"internal/log"
	"net/http"
	"os"
	"strings"

	"internal/config"
)

var (
	//go:embed resources/*
	res embed.FS
)

func main() {
	db.Init(config.Config().DbUri)

	http.HandleFunc("/", handleHomePage)

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infof("dashboard's server started at port %s\n", config.Config().DashboardPort)
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":"+config.Config().DashboardPort, nil))
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "robots.txt") {
		robotsFile, _ := os.ReadFile("./resources/robots.txt")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
		return
	}
}
