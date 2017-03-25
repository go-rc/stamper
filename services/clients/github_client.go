package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL           = "https://api.github.com"
	previewMediaType = "application/vnd.github.korra-preview"
)

type UserPermission struct {
	Permission string `json:"permission"`
}

type GitHubClient struct {
	token string
}

func NewGitHubClient(token string) *GitHubClient {
	return &GitHubClient{token: token}
}

func (c *GitHubClient) GetUserPermissions(nwo, login string) (string, error) {
	url := fmt.Sprintf("%s/repos/%s/collaborators/%s/permission", apiURL, nwo, login)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.token))
	req.Header.Add("Accept", previewMediaType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var payload UserPermission

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return "", err
	}

	return payload.Permission, nil
}
