package splitstrings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu6/SplitStrings"
)

var _ = Describe("SplitString with regex", func() {
	It("Basic tests", func() {
		Expect(SplitStringRegex("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(SplitStringRegex("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(SplitStringRegex("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})
