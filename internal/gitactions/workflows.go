package gitactions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/isadri/cicd-dashboard/internal/urls"
	"github.com/isadri/cicd-dashboard/internal/utils"
)

type step struct {
	Name   string
	Status string
	Number int
}

type Job struct {
	URL          string
	Status       string
	Name         string
	Steps        []step
	WorkflowName string
	HeadBranch   string
}

type workflow struct {
	ID        int
	Name      string
	Path      string
	State     string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string
	HTMLURL   string `json:"html_url"`
}

type Workflows struct {
	TotalCount int        `json:"total_count"`
	Items      []workflow `json:"workflows"`
}

func GetWorkflows(owner, repo string) (*Workflows, error) {
	log.Printf("server trying to fetch %s", urls.GetWorkflowsUrl(owner, repo))
	req, err := http.NewRequest(http.MethodGet, urls.GetWorkflowsUrl(owner, repo), nil)
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
		return nil, fmt.Errorf("fetch workflows failed: %s", resp.Status)
	}
	var workflows Workflows

	if err := json.NewDecoder(resp.Body).Decode(&workflows); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %s", err)
	}
	return &workflows, nil
}
