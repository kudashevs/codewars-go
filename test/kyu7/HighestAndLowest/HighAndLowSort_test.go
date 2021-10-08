package highestandlowest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu7/HighestAndLowest"
)

var _ = Describe("HighAndLow with sorting", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(HighAndLowSort("8 3 -5 42 -1 0 0 -9 4 7 4 -4")).To(Equal("42 -9"))
	})
})
