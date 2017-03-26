package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	apiURL = "https://api.github.com"

	// https://developer.github.com/changes/2016-09-14-Integrations-Early-Access/
	mediaTypeIntegrationPreview = "application/vnd.github.machine-man-preview+json"

	// https://developer.github.com/changes/2016-11-28-preview-org-membership/
	mediaTypeOrgMembershipPreview = "application/vnd.github.korra-preview+json"
)

type GitHubClient struct {
	token string
}

func NewGitHubClient(installationID int, integrationID string, privateKey []byte) (*GitHubClient, error) {
	integrationToken, err := generateJWT(integrationID, privateKey)
	if err != nil {
		return nil, err
	}

	installationToken, err := createToken(installationID, integrationToken)
	if err != nil {
		return nil, err
	}

	return &GitHubClient{token: installationToken}, nil
}

type RepositoryPermissionLevel struct {
	Permission string `json:"permission,omitempty"`
}

func (c *GitHubClient) GetPermissionLevel(nwo, login string) (string, error) {
	url := fmt.Sprintf("%s/repos/%s/collaborators/%s/permission", apiURL, nwo, login)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.token))
	req.Header.Add("Accept", mediaTypeOrgMembershipPreview)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var payload RepositoryPermissionLevel

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return "", err
	}

	return payload.Permission, nil
}

func generateJWT(integrationID string, privateKey []byte) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute).Unix(),
		"iss": integrationID,
	})

	return t.SignedString(key)
}

func createToken(installationID int, integrationToken string) (string, error) {
	url := fmt.Sprintf("%s/installations/%d/access_tokens", apiURL, installationID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", integrationToken))
	req.Header.Add("Accept", mediaTypeIntegrationPreview)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var payload map[string]interface{}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return "", err
	}

	return payload["token"].(string), nil
}
