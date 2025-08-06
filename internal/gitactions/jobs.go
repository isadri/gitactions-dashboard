package gitactions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/isadri/cicd-dashboard/internal/urls"
	"github.com/isadri/cicd-dashboard/internal/utils"
)

type Step struct {
	Name        string
	Conclusion  string
	Number      int
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type Job struct {
	ID           int64
	WorkflowName string `json:"workflow_name"`
	Branch       string `json:"head_branch"`
	Name         string
	HTMLURL      string `json:"html_url"`
	Status       string
	Conclusion   string
	StartedAt    time.Time `json:"started_at"`
	Steps        []Step
}

type WorkflowRunsJobs struct {
	Jobs []Job
}

func GetWorkflowRunJobs(owner, repo string, runId string) (*WorkflowRunsJobs, error) {
	log := utils.GetLogger()
	reqUrl := urls.GetWorkflowRunJobsUrl(owner, repo, runId)

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
		return nil, fmt.Errorf("get workflow run jobs failed: %s",
			resp.Status)
	}
	var jobs WorkflowRunsJobs

	if err := json.NewDecoder(resp.Body).Decode(&jobs); err != nil {
		return nil, err
	}
	return &jobs, nil
}
