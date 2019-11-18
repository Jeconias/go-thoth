package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hexcolor", func() {

	It("should empty validation", func() {
		s := "#fcb603"
		m := models.HexcolorValidate{
			Hexcolor: s,
			Pointer:  &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate hexcolor when hexcolor invalid", func() {
		s := "453e2b"
		m := models.HexcolorValidate{
			Hexcolor: "#fcb603",
			Pointer:  &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("hexcolor"))
	})

	It("should fail to validate without field `Hexcolor`", func() {
		s := "#fcb603"
		m := models.HexcolorValidate{
			// Hexcolor: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Hexcolor"))
		Expect(errs[0].Tag()).To(Equal("hexcolor"))
	})

	It("should fail to validater without field `Pointer`", func() {
		s := "#fcb603"
		m := models.HexcolorValidate{
			Hexcolor: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("hexcolor"))
	})

})
