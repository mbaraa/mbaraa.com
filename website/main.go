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

func main() {
	http.HandleFunc("/", handleHomePage)
	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infoln("server started at port 3000")
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":3000", nil))
}
