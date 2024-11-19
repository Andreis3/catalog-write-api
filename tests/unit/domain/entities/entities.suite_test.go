//go:build unit

package entities_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_EntitiesSuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = false
	reporterConfig.Verbose = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Entities tests", suiteConfig, reporterConfig)

}
