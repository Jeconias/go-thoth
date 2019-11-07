package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Numbers", func() {
	Describe("Equal", func() {
		It("should fail to validate all numbers", func() {
			m := models.TypeEqNumber{}

			errs := m.Validate()
			Expect(errs).To(HaveLen(30))
			Expect(errs[0].Field()).To(Equal("Uint"))
			Expect(errs[0].Tag()).To(Equal("eq"))
		})
	})
})
