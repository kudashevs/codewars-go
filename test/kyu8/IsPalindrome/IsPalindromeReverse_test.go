package ispalindrome_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/IsPalindrome"
)

var _ = Describe("IsPalindromeReverse", func() {
	It("tests basic strings", func() {
		Expect(IsPalindromeReverse("a")).To(Equal(true))
		Expect(IsPalindromeReverse("aba")).To(Equal(true))
		Expect(IsPalindromeReverse("Abba")).To(Equal(true))
		Expect(IsPalindromeReverse("hello")).To(Equal(false))
	})
})
