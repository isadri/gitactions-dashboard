package server

import (
	"html/template"
	"net/http"

	"github.com/isadri/cicd-dashboard/internal/gitactions"
	"github.com/isadri/cicd-dashboard/internal/repos"
	"github.com/isadri/cicd-dashboard/internal/utils"
)

func RegisterFuncs() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/repo", repoHandler)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	log := utils.GetLogger()
	log.Infof("%s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)

	log.Info("get fasgo-app repositories")
	repos, err := repos.GetRepos("fasgo-app")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Info("creating home.index template")
	templ := template.Must(template.New("home.html").
		ParseFiles("web/template/home.html"))
	log.Info("executing home.html template")
	if err := templ.Execute(w, repos); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func repoHandler(w http.ResponseWriter, req *http.Request) {
	log := utils.GetLogger()
	log.Infof("%s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)

	var workflows *gitactions.Workflows

	repoName := req.URL.Query().Get("name")
	if repoName == "" {
		log.Error("repository name path parameter is required")
		w.Write([]byte("repository name path parameter is required"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	workflows, err := gitactions.GetWorkflows("fasgo-app", repoName)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("creating customer_workflows.html template")
	templ := template.Must(template.New("customer_workflows.html").
		ParseFiles("web/template/customer_workflows.html"))
	log.Info("executing customer_workflows.html")
	if err := templ.Execute(w, workflows); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
