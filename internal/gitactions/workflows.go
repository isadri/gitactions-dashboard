package gitactions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/isadri/cicd-dashboard/internal/urls"
	"github.com/isadri/cicd-dashboard/internal/utils"
)

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Commit struct {
	Message string
}

type Repo struct {
	Name    string
	HTMLURL string `json:"html_url"`
}

type Workflow struct {
	ID         int64
	Repository Repo
	Name       string
	Branch     string `json:"head_branch"`
	Conclusion string
	HTMLURL    string    `json:"html_url"`
	RunAttempt int       `json:"run_attempt"`
	StartedAt  time.Time `json:"run_started_at"`
	Actor      User      `json:"triggering_actor"`
	HeadCommit Commit    `json:"head_commit"`
}

type WorkflowRuns struct {
	Workflows []Workflow `json:"workflow_runs"`
}

func GetWorkflowRuns(owner, repo string) (*WorkflowRuns, error) {
	log := utils.GetLogger()
	reqUrl := urls.GetWorkflowsRunsUrl(owner, repo)
	log.Infof("request to %s", reqUrl)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	utils.SetGitHubHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get workflow runs failed: %s", resp.Status)
	}

	var workflows WorkflowRuns

	if err := json.NewDecoder(resp.Body).Decode(&workflows); err != nil {
		return nil, err
	}
	return &workflows, nil
}
