package removefirstlastchar_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRemoveFirstLastCharacter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RemoveFirstLastCharacter Suite")
}
