package main

import (
	"github.com/weaveworks/release-note-action/action"
)

func main() {
	action.Log("running release note action")
	ctx, err := action.NewContextFromEnv()
	if err != nil {
		action.LogErrorAndExit("failed to create action context: %s", err.Error())
	}

	a := action.New()
	exists, err := a.ReleaseNoteExists(ctx)
	if err != nil {
		action.LogErrorAndExit("failed checking for release note: %s", err.Error())
	}

	if exists {
		action.Log("release-note code block found")
	} else {
		action.LogErrorAndExit("release not not found. Add a release-note code block to your PR description")
	}
}
