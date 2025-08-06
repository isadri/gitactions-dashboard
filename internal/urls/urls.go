package urls

import (
	"fmt"
	"os"
)

func GetReposUrl(org string) string {
	if os.Getenv("ORG_TYPE") == "user" {
		return fmt.Sprintf("https://api.github.com/users/%s/repos",
			org)
	}
	return fmt.Sprintf("https://api.github.com/orgs/%s/repos",
		org)
}

func GetWorkflowsUrl(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows",
		owner, repo)
}

func GetWorkflowsRunsUrl(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runs",
		owner, repo)
}

func GetWorkflowRunJobsUrl(owner, repo string, runId string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runs/%s/jobs",
		owner, repo, runId)
}

func GetJobLogsUrl(owner, repo string, jobId string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/jobs/%s/logs",
		owner, repo, jobId)
}
