package mumbling_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMumbling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mumbling Suite")
}
