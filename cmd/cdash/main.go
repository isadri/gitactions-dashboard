package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/isadri/cicd-dashboard/internal/gitactions"

	"github.com/joho/godotenv"
)

func run(w http.ResponseWriter, r *http.Request) {
	log.Printf("+ accept connection %s", r.RemoteAddr)
	var workflows *gitactions.Workflows
	var err error

	if workflows, err = gitactions.GetWorkflows("isadri", "eval"); err != nil {
		log.Fatal(err)
	}

	log.Print("create dashboard.html template")
	templ := template.Must(template.New("dashboard.html").
		ParseFiles("web/template/dashboard.html"))
	log.Print("execute dashboard.html template")
	if err := templ.Execute(w, workflows); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	log.Print("loading environment variables in .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/", run)
	log.Print("start server")
	log.Print("server listening on localhost:8000")
	http.ListenAndServe("localhost:8000", nil)
}
