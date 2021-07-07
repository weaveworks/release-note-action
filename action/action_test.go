package action_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weaveworks/release-note-action/action"
)

const (
	prEventName = "pull_request"
)

var _ = Describe("TestReleaseNote", func() {
	var (
		err   error
		a     action.ReleaseNoteAction
		ctx   *action.Context
		found bool
	)

	BeforeEach(func() {
		a = action.New()
	})

	Describe("Processing a pull request with a release note", func() {
		BeforeEach(func() {
			setEnvVars(prEventName, "testdata/pr.json")
			ctx, err = action.NewContextFromEnv()
			Expect(err).NotTo(HaveOccurred())
			found, err = a.ReleaseNoteExists(ctx)
		})

		It("should have found a release note", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())
		})
	})

	Describe("Processing a pull request with a release note of NONE", func() {
		BeforeEach(func() {
			setEnvVars(prEventName, "testdata/pr_none.json")
			ctx, err = action.NewContextFromEnv()
			Expect(err).NotTo(HaveOccurred())
			found, err = a.ReleaseNoteExists(ctx)
		})

		It("should have found a release note", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())
		})
	})

	Describe("Processing a pull request with a action required release note", func() {
		BeforeEach(func() {
			setEnvVars(prEventName, "testdata/pr_action.json")
			ctx, err = action.NewContextFromEnv()
			Expect(err).NotTo(HaveOccurred())
			found, err = a.ReleaseNoteExists(ctx)
		})

		It("should have found a release note", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())
		})
	})

	Describe("Processing a pull request without a release note", func() {
		BeforeEach(func() {
			setEnvVars(prEventName, "testdata/pr_noreleasenote.json")
			ctx, err = action.NewContextFromEnv()
			Expect(err).NotTo(HaveOccurred())
			found, err = a.ReleaseNoteExists(ctx)
		})

		It("should have not found an error", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})

	Describe("Processing an event that isn't a pull request", func() {
		BeforeEach(func() {
			setEnvVars(prEventName, "testdata/non_pr.json")
			ctx, err = action.NewContextFromEnv()
			Expect(err).NotTo(HaveOccurred())
			found, err = a.ReleaseNoteExists(ctx)
		})

		It("should have raised a missing PR error", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})

func setEnvVars(eventName, eventPath string) {
	os.Setenv("GITHUB_SHA", "0df36e327f56c8897894b9f52734c69b56ff6e79")
	os.Setenv("GITHUB_REF", "refs/pull/5/merge")
	os.Setenv("GITHUB_EVENT_NAME", eventName)
	os.Setenv("GITHUB_WORKFLOW", "pr_size")
	os.Setenv("GITHUB_ACTION", "richardcaserelease-note-action")
	os.Setenv("GITHUB_ACTOR", "actionsbot")
	os.Setenv("GITHUB_JOB", "labeler")
	os.Setenv("GITHUB_RUN_NUMBER", "5")
	os.Setenv("GITHUB_RUN_ID", "982249112")
	os.Setenv("GITHUB_REPOSITORY", "test/repo")

	os.Setenv("GITHUB_EVENT_PATH", eventPath)
}
