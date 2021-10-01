package vowelcount_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu7/VowelCount"
)

var _ = Describe("GetCount with intersection", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(GetCountIntersection("abracadabra")).To(Equal(5))
	})
})
