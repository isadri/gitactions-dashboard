package server

import (
	"html/template"
	"log"
	"net/http"
	
	"github.com/isadri/cicd-dashboard/internal/gitactions"
)

func RegisterFuncs() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/customer", customerHandler)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)
	log.Print("creating home.index template")
	templ := template.Must(template.New("home.html").
		ParseFiles("web/template/home.html"))
	log.Print("executing home.html template")
	if err := templ.Execute(w, nil); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func customerHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)

	var workflows *gitactions.Workflows

	workflows, err := gitactions.GetWorkflows("fasgo-app", "Fasgo_Customer")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print("creating customer_workflows.html template")
	templ := template.Must(template.New("customer_workflows.html").
		ParseFiles("web/template/customer_workflows.html"))
	log.Print("executing customer_workflows.html")
	if err := templ.Execute(w, workflows); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
