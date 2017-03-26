package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tombell/stamper/services/clients"
)

type GitHubService struct {
	Logger *log.Logger
}

func SetupGitHubService(id, cert string, l *log.Logger) {
	Service = &GitHubService{Logger: l}
}

var Service *GitHubService

func (s *GitHubService) HandleEvent(event string, body []byte) error {
	var payload EventPayload

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	// TODO: using env temporarily
	client := clients.NewGitHubClient(os.Getenv("GITHUB_API_TOKEN"))

	permission, err := client.GetUserPermissions(payload.Repository.FullName, payload.Sender.Login)
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
