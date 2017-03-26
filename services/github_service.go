package services

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/tombell/stamper/services/clients"
)

type GitHubService struct {
	integrationID string
	privateKey    []byte
	logger        *log.Logger
}

func SetupGitHubService(id, cert string, l *log.Logger) error {
	key, err := ioutil.ReadFile(cert)
	if err != nil {
		return err
	}

	Service = &GitHubService{
		integrationID: id,
		privateKey:    key,
		logger:        l,
	}

	return nil
}

var Service *GitHubService

func (s *GitHubService) HandleEvent(event string, body []byte) error {
	var payload EventPayload

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	client, err := clients.NewGitHubClient(payload.Installation.ID, s.integrationID, s.privateKey)
	if err != nil {
		return err
	}

	permission, err := client.GetPermissionLevel(payload.Repository.FullName, payload.Sender.Login)
	if err != nil {
		return err
	}

	s.logger.Printf(
		"sender %s has %s access to %s\n",
		payload.Sender.Login,
		permission,
		payload.Repository.FullName,
	)

	return nil
}
