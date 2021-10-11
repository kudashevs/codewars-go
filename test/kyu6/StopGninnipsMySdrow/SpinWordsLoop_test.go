package stopgninnipsmysdrow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu6/StopGninnipsMySdrow"
)

var _ = Describe("SpinWordsLoop with accumulating loop", func() {
	It("should test that the solution returns the correct value for single word inputs", func() {
		Expect(SpinWordsLoop("Welcome")).To(Equal("emocleW"))
		Expect(SpinWordsLoop("to")).To(Equal("to"))
		Expect(SpinWordsLoop("CodeWars")).To(Equal("sraWedoC"))
	})

	It("should test that the solution returns the correct value for multiple word outputs", func() {
		Expect(SpinWordsLoop("Hey fellow warriors")).To(Equal("Hey wollef sroirraw"))
		Expect(SpinWordsLoop("Burgers are my favorite fruit")).To(Equal("sregruB are my etirovaf tiurf"))
		Expect(SpinWordsLoop("Pizza is the best vegetable")).To(Equal("azziP is the best elbategev"))
	})
})
