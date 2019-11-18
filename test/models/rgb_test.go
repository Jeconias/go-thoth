package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RGB", func() {

	It("should empty validation", func() {
		s := "rgb(235, 64, 52)"
		m := models.RGBValidate{
			RGB:     s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate rgb when rgb invalid", func() {
		s := "453e2b"
		m := models.RGBValidate{
			RGB:     "rgb(235, 64, 52)",
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("rgb"))
	})

	It("should fail to validate without field `RGB`", func() {
		s := "rgb(235, 64, 52)"
		m := models.RGBValidate{
			// RGB: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("RGB"))
		Expect(errs[0].Tag()).To(Equal("rgb"))
	})

	It("should fail to validater without field `Pointer`", func() {
		s := "rgb(235, 64, 52)"
		m := models.RGBValidate{
			RGB: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("rgb"))
	})

})
