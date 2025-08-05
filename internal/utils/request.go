package utils

import (
	"net/http"
	"os"
)

func SetGitHubHeaders(req *http.Request) {
	token := os.Getenv("GITHUB_TOKEN")
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
}
