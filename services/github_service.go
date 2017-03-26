package services

import (
	"encoding/json"
	"log"

	"github.com/tombell/stamper/services/clients"
)

type OpenedEvent struct {
	Action      string     `json:"action"`
	Repository  Repository `json:"repository"`
	Sender      Sender     `json:"sender"`
	Issue       Issue      `json:"issue"`
	PullRequest Issue      `json:"pull_request"`
}

type CommentEvent struct {
	Action     string     `json:"action"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
	Issue      Issue      `json:"issue"`
	Comment    Comment    `json:"comment"`
}

type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type Sender struct {
	Login string `json:"login"`
}

type Issue struct {
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

type Comment struct {
	Body string `json:"body"`
}

type Installation struct {
	ID int `json:"id"`
}

type GitHubService struct {
	Client *clients.GitHubClient
	Logger *log.Logger
}

func SetupGitHubService(token string, l *log.Logger) *GitHubService {
	Service = &GitHubService{
		Client: clients.NewGitHubClient(token),
		Logger: l,
	}
	return Service
}

var Service *GitHubService

func (s *GitHubService) HandleEvent(event string, body []byte) error {
	switch event {
	case "issues", "pull_request":
		return s.handleOpenedEvent(body)
	case "issue_comment":
		return s.handleCommentEvent(body)
	}

	return nil
}

func (s *GitHubService) handleOpenedEvent(body []byte) error {
	var payload OpenedEvent

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	s.Logger.Printf("issue/pr opened by %s\n", payload.Sender.Login)

	permission, err := s.Client.GetUserPermissions(
		payload.Repository.FullName,
		payload.Sender.Login,
	)
	if err != nil {
		return err
	}

	s.Logger.Printf(
		"sender %s has %s access to %s\n",
		payload.Sender.Login,
		permission,
		payload.Repository.FullName,
	)

	return nil
}

func (s *GitHubService) handleCommentEvent(body []byte) error {
	var payload CommentEvent

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	s.Logger.Printf("issue/pr comment by %s\n", payload.Sender.Login)

	permission, err := s.Client.GetUserPermissions(
		payload.Repository.FullName,
		payload.Sender.Login,
	)
	if err != nil {
		return err
	}

	s.Logger.Printf(
		"sender %s has %s access to %s\n",
		payload.Sender.Login,
		permission,
		payload.Repository.FullName,
	)

	return nil
}
