package stringendswith_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu7/stringendswith"
)

var _ = Describe("StringEndsWith", func() {
	It("Should work for fixed tests", func() {
		Expect(StringEndsWith("", "")).To(Equal(true))
		Expect(StringEndsWith(" ", "")).To(Equal(true))
		Expect(StringEndsWith("abc", "c")).To(Equal(true))
		Expect(StringEndsWith("banana", "ana")).To(Equal(true))
		Expect(StringEndsWith("a", "z")).To(Equal(false))
		Expect(StringEndsWith("", "t")).To(Equal(false))
		Expect(StringEndsWith("$a = $b + 1", "+1")).To(Equal(false))
		Expect(StringEndsWith("    ", "   ")).To(Equal(true))
		Expect(StringEndsWith("1oo", "100")).To(Equal(false))
	})
})
