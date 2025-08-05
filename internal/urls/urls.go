package urls

import "fmt"

func GetReposUrl(org string) string {
	return fmt.Sprintf("https://api.github.com/orgs/%s/repos",
		org)
}

func GetWorkflowsUrl(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows",
		owner, repo)
}
