package centuryfromyear_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/CenturyFromYear"
)

var _ = Describe("Century", func() {
	It("should return the correct values", func() {
		Expect(Century(int(1990))).To(Equal(20))
		Expect(Century(int(1705))).To(Equal(18))
		Expect(Century(int(1900))).To(Equal(19))
		Expect(Century(int(1601))).To(Equal(17))
		Expect(Century(int(2000))).To(Equal(20))
		Expect(Century(int(89))).To(Equal(1))
	})
})
