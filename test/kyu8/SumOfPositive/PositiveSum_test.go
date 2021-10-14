package sumofpositive_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/sumofpositive"
)

var _ = Describe("PositiveSum", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(PositiveSum([]int{1, 2, 3, 4, 5})).To(Equal(15))
		// Expect(PositiveSum([]int{1, -2, 3, 4, 5})).To(Equal(13))
		// Expect(PositiveSum([]int{})).To(Equal(0))
		// Expect(PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
		// Expect(PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))
	})
})
