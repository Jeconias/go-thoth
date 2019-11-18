package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HSL", func() {

	It("should empty validation", func() {
		s := "hsl(0, 100%, 50%)"
		m := models.HSLValidate{
			HSL:     s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate hsl when hsl invalid", func() {
		s := "453e2b"
		m := models.HSLValidate{
			HSL:     "hsl(0, 100%, 50%)",
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("hsl"))
	})

	It("should fail to validate without field `HSL`", func() {
		s := "hsl(0, 100%, 50%)"
		m := models.HSLValidate{
			// HSL: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("HSL"))
		Expect(errs[0].Tag()).To(Equal("hsl"))
	})

	It("should fail to validater without field `Pointer`", func() {
		s := "hsl(0, 100%, 50%)"
		m := models.HSLValidate{
			HSL: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("hsl"))
	})

})
