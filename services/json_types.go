package services

type EventPayload struct {
	Action       string       `json:"action"`
	Repository   Repository   `json:"repository"`
	Sender       Sender       `json:"sender"`
	Issue        Issue        `json:"issue,omitempty"`
	PullRequest  Issue        `json:"pull_request,omitempty"`
	Comment      Comment      `json:"comment,omitempty"`
	Installation Installation `json:"installation,omitempty"`
}

type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type Sender struct {
	Login string `json:"login"`
}

type Issue struct {
	Number int    `json:"number"`
	Body   string `json:"body"`
}

type Comment struct {
	Body string `json:"body"`
}

type Installation struct {
	ID int `json:"id"`
}
