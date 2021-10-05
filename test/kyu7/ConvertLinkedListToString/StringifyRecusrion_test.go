package convertlinkedlisttostring_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu7/ConvertLinkedListToString"
)

var _ = Describe("Stringify with recursion", func() {
	It("should work for sample list", func() {
		Expect(StringifyRecursion(NewNode(1, NewNode(2, NewNode(3, nil))))).To(Equal("1 -> 2 -> 3 -> nil"))
	})
})
