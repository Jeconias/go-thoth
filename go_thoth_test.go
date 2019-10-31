package thoth_test

import (
	myasthurts "github.com/lab259/go-my-ast-hurts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoThoth", func() {

	Context("parsing interfaces of struct with thoth tag", func() {

		It("should show func Home", func() {
			env, err := myasthurts.NewEnvironment()
			Expect(err).ToNot(HaveOccurred())

			pkg, err := env.Parse("./tests/gen")
			Expect(err).ToNot(HaveOccurred())

			Expect(pkg.Name).To(Equal("any"))

			refType, ok := pkg.RefTypeByName("Home")
			Expect(ok).To(BeTrue())
			Expect(refType.Name()).To(Equal("Home"))
		})

		It("should show func User", func() {
			env, err := myasthurts.NewEnvironment()
			Expect(err).ToNot(HaveOccurred())

			pkg, err := env.Parse("./tests/gen")
			Expect(err).ToNot(HaveOccurred())

			Expect(pkg.Name).To(Equal("any"))

			refType, ok := pkg.RefTypeByName("User")
			Expect(ok).To(BeTrue())
			Expect(refType.Name()).To(Equal("User"))
		})
	})
})
