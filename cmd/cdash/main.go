package main

import (
	// "encoding/json"
	// "fmt"
	"html/template"
	"log"
	"os"

	"github.com/isadri/cicd-dashboard/internal/gitactions"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var workflows *gitactions.Workflows

	if workflows, err = gitactions.GetWorkflows("isadri", "eval"); err != nil {
		log.Fatal(err)
	}
	templ := template.Must(template.New("dashboard.html").
		ParseFiles("web/templates/dashboard.html"))
	if err := templ.Execute(os.Stdout, workflows); err != nil {
		log.Fatal(err)
	}
}
