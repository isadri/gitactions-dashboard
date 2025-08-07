package repos

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/isadri/gitactions-dashboard/internal/urls"
	"github.com/isadri/gitactions-dashboard/internal/utils"
)

type user struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Repo struct {
	Name     string
	FullName string `json:"full_name"`
	Owner    user
	HTMLURL  string `json:"html_url"`
}

func GetRepos(org string) ([]Repo, error) {
	req, err := http.NewRequest(http.MethodGet, urls.GetReposUrl(org), nil)

	if err != nil {
		return nil, err
	}
	utils.SetGitHubHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer utils.Close(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get %s repositories failed: %s",
			org, resp.Status)
	}

	var repos []Repo

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}
