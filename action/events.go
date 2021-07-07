package action

import "time"

// Event represents the GitHub Actions event trigger.
type Event struct {
	Action string `json:"action"`
	After  string `json:"after"`
	Before string `json:"before"`
	Number int    `json:"number"`

	PullRequest *PullRequest `json:"pull_request"`
}

// PullRequest represents the payload for a pull request GitHub Actions event.
// NOTE: not all the fields have been added.
type PullRequest struct {
	Body         string    `json:"body"`
	ChangedFiles int       `json:"changed_files"`
	Commits      int       `json:"commits"`
	CreatedAt    time.Time `json:"created_at"`
	Draft        bool      `json:"draft"`
	HTMLURL      string    `json:"html_url"`
	ID           int       `json:"id"`
	Number       int       `json:"number"`
	Title        string    `json:"title"`
}
