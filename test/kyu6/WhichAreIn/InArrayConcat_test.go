package WhichAreIn_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu6/WhichAreIn"
)

func doInArrayConcatTest(array1 []string, array2 []string, exp []string) {
	var ans = InArrayLoop(array1, array2)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("InArray with concatenation", func() {
	It("should handle basic cases", func() {
		var a1 = []string{"live", "arp", "strong"}
		var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
		var r = []string{"arp", "live", "strong"}
		doInArrayConcatTest(a1, a2, r)

		a1 = []string{"cod", "code", "wars", "ewar", "ar"}
		a2 = []string{}
		r = []string{}
		doInArrayConcatTest(a1, a2, r)
	})
})
