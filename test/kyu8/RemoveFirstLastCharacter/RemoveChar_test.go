package removefirstlastchar_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/RemoveFirstLastCharacter"
)

var _ = Describe("RemoveChar", func() {
	Describe("Fixed Tests", func() {
		It("eloquent", func() {
			Expect(RemoveChar("eloquent")).To(Equal("loquen"))
		})

		It("country", func() {
			Expect(RemoveChar("country")).To(Equal("ountr"))
		})

		It("person", func() {
			Expect(RemoveChar("person")).To(Equal("erso"))
		})

		It("place", func() {
			Expect(RemoveChar("place")).To(Equal("lac"))
		})
	})
})
