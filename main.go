package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/weaveworks/release-note-action/action"
)

func main() {
	a, err := action.New(releaseNotAction)
	if err != nil {
		fmt.Fprintln(os.Stdout, "failed to create action: %w", err)
		os.Exit(1)
	}

	if err := a.Do(); err != nil {
		fmt.Printf("action failed: %s\n", err.Error())
		os.Exit(1)
	}
}

func releaseNotAction(ctx *action.Context) error {
	fmt.Println("Running release note action")

	pr := ctx.Event.PullRequest
	if pr == nil {
		return errors.New("no pull request event data")
	}

	re := regexp.MustCompile(`(?s)(?:Release note\*\*:\s*(?:<!--[^<>]*-->\s*)?` + "```(?:release-note)?|```release-note)(.+?)```")
	potentialMatch := re.FindStringSubmatch(pr.Body)
	if potentialMatch == nil {
		return errors.New("no release-note code block found")
	} else {
		fmt.Fprintln(os.Stdout, "release note found")
	}

	return nil
}
