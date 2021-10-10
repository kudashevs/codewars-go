package splitstrings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu6/SplitStrings"
)

var _ = Describe("SplitString with loop", func() {
	It("Basic tests", func() {
		Expect(SplitStringLoop("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(SplitStringLoop("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(SplitStringLoop("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})
