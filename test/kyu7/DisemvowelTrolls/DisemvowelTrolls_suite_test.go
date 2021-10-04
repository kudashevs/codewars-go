package disemvoweltrolls_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDisemvowelTrolls(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DisemvowelTrolls Suite")
}
