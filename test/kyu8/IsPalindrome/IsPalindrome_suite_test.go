package ispalindrome_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIsPalindrome(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsPalindrome Suite")
}
