package findsmallestinteger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/FindTheSmallestIntegerInArray"
)

var _ = Describe("FindTheSmallestIntegerInArray with min", func() {
	It("should work for sample tests", func() {
		Expect(Expect(SmallestIntegerFinderMin([]int{34, 15, 88, 2})).To(Equal(2)))
		Expect(Expect(SmallestIntegerFinderMin([]int{34, -345, -1, 100})).To(Equal(-345)))
	})
})
