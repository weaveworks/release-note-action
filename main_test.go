package main

import (
	"os"
	"testing"

	"github.com/weaveworks/release-note-action/action"
)

func TestSuccess(t *testing.T) {
	os.Setenv("GITHUB_SHA", "0df36e327f56c8897894b9f52734c69b56ff6e79")
	os.Setenv("GITHUB_REF", "refs/pull/5/merge")
	os.Setenv("GITHUB_EVENT_NAME", "pull_request")
	os.Setenv("GITHUB_WORKFLOW", "pr_size")
	os.Setenv("GITHUB_ACTION", "richardcaserelease-note-action")
	os.Setenv("GITHUB_ACTOR", "actionsbot")
	os.Setenv("GITHUB_JOB", "labeler")
	os.Setenv("GITHUB_RUN_NUMBER", "5")
	os.Setenv("GITHUB_RUN_ID", "982249112")
	os.Setenv("GITHUB_REPOSITORY", "test/repo")

	eventPath := "./sample/pr.json"
	os.Setenv("GITHUB_EVENT_PATH", eventPath)

	a, err := action.New(releaseNotAction)
	if err != nil {
		t.Errorf("failed to create action: %w", err)
	}

	if err := a.Do(); err != nil {
		t.Errorf("action failed: %w", err)
	}
}
