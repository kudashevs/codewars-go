package stopgninnipsmysdrow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu6/StopGninnipsMySdrow"
)

var _ = Describe("SpinWordsLoop with swapping loop", func() {
	It("should test that the solution returns the correct value for single word inputs", func() {
		Expect(SpinWordsLoopSwap("Welcome")).To(Equal("emocleW"))
		Expect(SpinWordsLoopSwap("to")).To(Equal("to"))
		Expect(SpinWordsLoopSwap("CodeWars")).To(Equal("sraWedoC"))
	})

	It("should test that the solution returns the correct value for multiple word outputs", func() {
		Expect(SpinWordsLoopSwap("Hey fellow warriors")).To(Equal("Hey wollef sroirraw"))
		Expect(SpinWordsLoopSwap("Burgers are my favorite fruit")).To(Equal("sregruB are my etirovaf tiurf"))
		Expect(SpinWordsLoopSwap("Pizza is the best vegetable")).To(Equal("azziP is the best elbategev"))
	})
})
