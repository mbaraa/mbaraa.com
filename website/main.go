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
	page := tmplrndr.NewIndex().Render(tmplrndr.IndexProps{})
	_, err := io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
	}
}

func main() {
	http.HandleFunc("/", handleHomePage)
	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infoln("server started at port 3000")
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":3000", nil))
}
