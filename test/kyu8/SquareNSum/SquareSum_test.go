package squarensum_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/SquareNSum"
)

var _ = Describe("SquareSum", func() {
	It("Testing [1,2]", func() { Expect(SquareSum([]int{1, 2})).To(Equal(5)) })
	It("Testing [0,3,4,5]", func() { Expect(SquareSum([]int{0, 3, 4, 5})).To(Equal(50)) })
})
