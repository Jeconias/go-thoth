package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hexadecimal", func() {

	var (
		s     = "ABCDEF"
		slice = []string{"123", "1495"}
	)

	It("should empty validation", func() {
		m := models.HexadecimalValidate{
			Hexadecimal: s,
			Pointer:     &s,
			Slice:       slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate hexadecimal when hexadecimal invalid", func() {
		ss := "invalid-hexadecimal"
		m := models.HexadecimalValidate{
			Hexadecimal: s,
			Pointer:     &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("hexadecimal"))
	})

	It("should fail validate hexadecimal when Slice hexadecimal is invalid", func() {
		m := models.HexadecimalValidate{
			Hexadecimal: s,
			Pointer:     &s,
			Slice:       []string{"GHI", "A123"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("hexadecimal"))
	})

	It("should fail to validate without field `Hexadecimal`", func() {
		m := models.HexadecimalValidate{
			// Hexadecimal: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Hexadecimal"))
		Expect(errs[0].Tag()).To(Equal("hexadecimal"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.HexadecimalValidate{
			Hexadecimal: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("hexadecimal"))
	})

})
