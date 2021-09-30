package vowelcount_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu7/VowelCount"
)

var _ = Describe("GetCount with filter", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(GetCountFilter("abracadabra")).To(Equal(5))
	})
})
