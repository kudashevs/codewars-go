package jenny_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJennysSecret(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JennysSecret Suite")
}
