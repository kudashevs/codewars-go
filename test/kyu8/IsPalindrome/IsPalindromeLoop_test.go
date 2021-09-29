package ispalindrome_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/IsPalindrome"
)

var _ = Describe("IsPalindrome with loop", func() {
	It("tests basic strings", func() {
		Expect(IsPalindromeLoop("a")).To(Equal(true))
		Expect(IsPalindromeLoop("aba")).To(Equal(true))
		Expect(IsPalindromeLoop("Abba")).To(Equal(true))
		Expect(IsPalindromeLoop("hello")).To(Equal(false))
	})
})
