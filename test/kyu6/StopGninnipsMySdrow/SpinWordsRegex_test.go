package stopgninnipsmysdrow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu6/StopGninnipsMySdrow"
)

var _ = Describe("SpinWordsLoop with regex", func() {
	It("should test that the solution returns the correct value for single word inputs", func() {
		Expect(SpinWordsRegex("Welcome")).To(Equal("emocleW"))
		Expect(SpinWordsRegex("to")).To(Equal("to"))
		Expect(SpinWordsRegex("CodeWars")).To(Equal("sraWedoC"))
	})

	It("should test that the solution returns the correct value for multiple word outputs", func() {
		Expect(SpinWordsRegex("Hey fellow warriors")).To(Equal("Hey wollef sroirraw"))
		Expect(SpinWordsRegex("Burgers are my favorite fruit")).To(Equal("sregruB are my etirovaf tiurf"))
		Expect(SpinWordsRegex("Pizza is the best vegetable")).To(Equal("azziP is the best elbategev"))
	})
})
