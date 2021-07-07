package action

import "errors"

var (
	ErrPullRequestEventRequired = errors.New("a pull request event is required")
	ErrNoReleaseNoteFound       = errors.New("no release-note code block found")
)
