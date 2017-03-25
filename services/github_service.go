package services

import (
	"encoding/json"
	"fmt"
)

// parse incoming webhook type

// - issue/pull request opened
// - issue/pull request comment created

// parse the webhook body depending on type

// does the user have contributer or higher access to the repo?
// https://developer.github.com/v3/repos/collaborators/#review-a-users-permission-level

// does the body contain the specified string?

// add the label to the issue/pull request with the specified label
// https://developer.github.com/v3/issues/#edit-an-issue

// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------

func HandleEvent(event string, body []byte) error {
	switch event {
	case "issues", "pull_request":
		return HandleOpenedEvent(body)
	case "issue_comment":
		return HandleCommentEvent(body)
	default:
		panic("unrecognised incoming event")
	}
}

func HandleOpenedEvent(body []byte) error {
	var payload OpenedEvent

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	fmt.Println("Issue or Pull Request Opened:")
	fmt.Println(payload)
	fmt.Println("---")

	return nil
}

func HandleCommentEvent(body []byte) error {
	var payload CommentEvent

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	fmt.Println("Issue or Pull Request Comment:")
	fmt.Println(payload)
	fmt.Println("---")

	return nil
}
