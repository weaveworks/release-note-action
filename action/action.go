package action

import (
	"regexp"
)

// ReleaseNoteAction represents the interface for a action to check for a release note
type ReleaseNoteAction interface {
	ReleaseNoteExists(ctx *Context) (bool, error)
}

// New creates a new instance of the release note action
func New() ReleaseNoteAction {
	return &action{}
}

// action contains the implementation of the release note action
type action struct{}

// ReleaseNoteExists will process a GitHub Actions event to see if a release note exists
func (a *action) ReleaseNoteExists(ctx *Context) (bool, error) {
	LogDebug("checking for existence of release note")

	pr := ctx.Event.PullRequest
	if pr == nil {
		return false, ErrPullRequestEventRequired
	}

	//nolint: lll
	re := regexp.MustCompile(`(?s)(?:Release note\*\*:\s*(?:<!--[^<>]*-->\s*)?` + "```(?:release-note)?|```release-note)(.+?)```")
	potentialMatch := re.FindStringSubmatch(pr.Body)
	if potentialMatch == nil {
		return false, nil
	}

	return true, nil
}
