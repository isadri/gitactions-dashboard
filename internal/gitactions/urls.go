package gitactions

import "fmt"

func getWorkflowsUrl(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows",
		owner, repo)
}
