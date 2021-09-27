package ispalindrome_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/IsPalindrome"
)

var _ = Describe("IsPalindrome", func() {
	It("tests basic strings", func() {
		Expect(IsPalindrome("a")).To(Equal(true))
		Expect(IsPalindrome("aba")).To(Equal(true))
		Expect(IsPalindrome("Abba")).To(Equal(true))
		Expect(IsPalindrome("hello")).To(Equal(false))
	})
})
