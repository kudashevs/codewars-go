package vowelcount_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVowelCount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VowelCount Suite")
}
