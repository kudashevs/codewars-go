package highestandlowest_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHighestAndLowest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HighestAndLowest Suite")
}
