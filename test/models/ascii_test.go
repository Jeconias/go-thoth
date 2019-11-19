package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ASCII", func() {

	var (
		s     = "Chico Bento :)"
		slice = []string{"Ch1qu1nh0"}
	)

	It("should empty validation", func() {
		m := models.ASCIIValidate{
			ASCII:   s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate ascii when ascii invalid", func() {
		ss := "http:ascii:invalid-ascii'\xc3'"
		m := models.ASCIIValidate{
			ASCII:   s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("ascii"))
	})

	It("should fail validate ascii when slice field is invalid", func() {
		m := models.ASCIIValidate{
			ASCII:   s,
			Pointer: &s,
			Slice:   []string{"localhost", "SGVsbG8gd29ybGQ=", "'\xc3'"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("ascii"))
	})

	It("should fail to validate field `ASCII`", func() {
		m := models.ASCIIValidate{
			ASCII:   "'\xc3",
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("ASCII"))
		Expect(errs[0].Tag()).To(Equal("ascii"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.ASCIIValidate{
			ASCII: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("ascii"))
	})

})
