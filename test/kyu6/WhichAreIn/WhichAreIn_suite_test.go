package WhichAreIn_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWhichAreIn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WhichAreIn Suite")
}
