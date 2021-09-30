package vowelcount_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu7/VowelCount"
)

var _ = Describe("GetCount with loop", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(GetCountLoop("abracadabra")).To(Equal(5))
	})
})
