package server

import (
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/isadri/cicd-dashboard/internal/gitactions"
	"github.com/isadri/cicd-dashboard/internal/repos"
	"github.com/isadri/cicd-dashboard/internal/utils"
)

func RegisterFuncs() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/repo", repoHandler)
	http.HandleFunc("/jobs", jobHandler)

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	log := utils.GetLogger()
	log.Infof("%s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)

	log.Info("get fasgo-app repositories")
	repos, err := repos.GetRepos(os.Getenv("ORG_NAME"))
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

	repoName := req.URL.Query().Get("name")
	if repoName == "" {
		log.Error("repository name path parameter is required")
		w.Write([]byte("repository name path parameter is required"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	workflows, err := gitactions.GetWorkflowRuns(os.Getenv("ORG_NAME"),
		repoName)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("creating workflows.html template")
	templ := template.Must(template.New("workflows.html").
		Funcs(template.FuncMap{
			"dateFormat": dateFormat,
			"replace":    strings.Replace,
		}).
		ParseFiles("web/template/workflows.html"))
	log.Info("executing workflows.html")
	if err := templ.Execute(w, *workflows); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func jobHandler(w http.ResponseWriter, req *http.Request) {
	log := utils.GetLogger()
	log.Infof("%s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)

	repoName := req.URL.Query().Get("repo")
	if repoName == "" {
		log.Error("missing repository name path paramter")
		w.Write([]byte("missing repository name path paramter"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	runId := req.URL.Query().Get("runid")
	if repoName == "" {
		log.Error("missing workflow run id paramter")
		w.Write([]byte("missing workflow run id paramter"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jobs, err := gitactions.GetWorkflowRunJobs(os.Getenv("ORG_NAME"),
		repoName, runId)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("create jobs.html template")
	templ := template.Must(template.New("jobs.html").
		Funcs(template.FuncMap{
			"dateFormat": dateFormat,
			"replace":    strings.Replace,
		}).ParseFiles("web/template/jobs.html"))
	log.Info("executing jobs.html template")
	if err := templ.Execute(w, *jobs); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func dateFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}
