package main

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"internal/db"
	"internal/log"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"internal/config"

	"github.com/mbaraa/mbaraa.com/dashboard/tmplrndr"
	"golang.org/x/crypto/bcrypt"
)

var (
	//go:embed resources/*
	res embed.FS

	lastLoginToken = ""
)

func main() {
	generateToken()
	db.Init(config.Config().DbUri)

	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/login", handleLogin)

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infof("dashboard's server started at port %s\n", config.Config().DashboardPort)
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":"+config.Config().DashboardPort, nil))
}

func generateToken() {
	rawThing := time.Now().String()
	sha256 := sha256.New()
	sha256.Write([]byte(rawThing))
	lastLoginToken = hex.EncodeToString(sha256.Sum(nil))
}

func checkAuthority(w http.ResponseWriter, r *http.Request) bool {
	token, err := r.Cookie("token")
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("you can't do that!"))
		return false
	}
	if !time.Now().After(token.Expires) {
		log.Errorln("token expired")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("you can't do that!"))
		return false
	}
	return token.Value == lastLoginToken
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "robots.txt") {
		robotsFile, _ := os.ReadFile("./resources/robots.txt")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
		return
	}
	if !checkAuthority(w, r) {
		return
	}

	page := tmplrndr.NewIndex().Render(tmplrndr.IndexProps{})
	_, _ = io.Copy(w, page)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		page := tmplrndr.NewLogin().Render(tmplrndr.LoginProps{})
		_, _ = io.Copy(w, page)
	case http.MethodPost:
		handleDoLogin(w, r)
	}
}

func handleDoLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	password := r.FormValue("password")
	hashed, err := bcrypt.GenerateFromPassword([]byte(config.Config().DashboardPassword), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Errorf("Unauthorized login attempt at %s\n", time.Now().String())
		return
	}

	generateToken()
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    lastLoginToken,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(time.Hour),
	})
}
