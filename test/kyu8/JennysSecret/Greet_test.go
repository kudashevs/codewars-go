package jenny_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu8/JennysSecret"
)

var _ = Describe("JennysSecret", func() {
	It("should work for sample test cases", func() {
		Expect(Greet("Alfred")).To(Equal("Hello, Alfred!"))
		Expect(Greet("Johnny")).To(Equal("Hello, my love!"))
	  })
})
