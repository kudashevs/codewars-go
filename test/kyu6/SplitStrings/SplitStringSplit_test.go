package splitstrings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu6/SplitStrings"
)

var _ = Describe("SplitString with split", func() {
	It("Basic tests", func() {
		Expect(SplitStringSplit("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(SplitStringSplit("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(SplitStringSplit("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})
