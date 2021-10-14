package sumofpositive_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSumOfPositive(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumOfPositive Suite")
}
