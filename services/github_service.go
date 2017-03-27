package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/tombell/stamper/services/clients"
)

type GitHubServiceConfig struct {
	IntegrationID string
	Cert          string
	Needle        string
	Label         string
	Logger        *log.Logger
}

type GitHubService struct {
	integrationID string
	needle        string
	label         string
	logger        *log.Logger
	privateKey    []byte
}

func SetupGitHubService(cfg *GitHubServiceConfig) error {
	key, err := ioutil.ReadFile(cfg.Cert)
	if err != nil {
		return err
	}

	Service = &GitHubService{
		integrationID: cfg.IntegrationID,
		needle:        cfg.Needle,
		label:         cfg.Label,
		logger:        cfg.Logger,
		privateKey:    key,
	}

	Service.logger.Printf("GitHub service ready...")

	return nil
}

func (srv *GitHubService) HandleEvent(event string, body []byte) error {
	var payload EventPayload

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	var str string
	var num int

	if event == "issues" && payload.Action == "opened" {
		str = payload.Issue.Body
		num = payload.Issue.Number
	} else if event == "pull_request" && payload.Action == "opened" {
		str = payload.PullRequest.Body
		num = payload.PullRequest.Number
	} else if event == "issue_comment" && payload.Action == "created" {
		str = payload.Comment.Body
		num = payload.Issue.Number
	} else {
		srv.logger.Printf("event/action not handled")
		return nil
	}

	if !strings.Contains(str, srv.needle) {
		srv.logger.Printf("needle not found in body: %s", srv.needle)
		return nil
	}

	client, err := clients.NewGitHubClient(payload.Installation.ID, srv.integrationID, srv.privateKey)
	if err != nil {
		return err
	}

	permission, err := client.GetPermissionLevel(payload.Repository.FullName, payload.Sender.Login)
	if err != nil {
		return err
	}

	fmt.Println(permission)

	if permission != "admin" && permission != "write" {
		srv.logger.Printf("sender does not have permission: %s", permission)
		return nil
	}

	err = client.AddLabelsToIssue(payload.Repository.FullName, num, []string{srv.label})
	if err != nil {
		return err
	}

	srv.logger.Printf("added label %s to issue", srv.label)
	return nil
}

var Service *GitHubService
