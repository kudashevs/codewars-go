package DisemvowelTrolls_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kudashevs/codewars-go/kata/kyu7/DisemvowelTrolls"
)

var _ = Describe("Disemvowel", func() {
	It("Testing for fixed tests", func() {
		Expect(Disemvowel("This website is for losers LOL!")).To(Equal("Ths wbst s fr lsrs LL!"))
		Expect(Disemvowel("No offense but,\nYour writing is among the worst I've ever read")).To(Equal("N ffns bt,\nYr wrtng s mng th wrst 'v vr rd"))
		Expect(Disemvowel("What are you, a communist?")).To(Equal("Wht r y,  cmmnst?"))
	})
})
