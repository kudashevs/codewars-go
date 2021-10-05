package squarensum_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSquareNSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SquareNSum Suite")
}
