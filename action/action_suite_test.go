package action_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMetalJanitorAction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Release Note Suite")
}
