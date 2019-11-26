package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MAC", func() {

	var (
		s     = "9c:75:ce:a0:d8:dd"
		slice = []string{"18:3a:77:ae:92:e1", "a8:cf:c3:37:5a:e4"}
	)

	It("should empty validation", func() {
		m := models.MACValidate{
			MAC:     s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate mac when mac invalid", func() {
		ss := "invalid-mac"
		m := models.MACValidate{
			MAC:     s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("mac"))
	})

	It("should fail validate mac when Slice mac is invalid", func() {
		m := models.MACValidate{
			MAC:     s,
			Pointer: &s,
			Slice:   []string{"invalid-mac", "6f:2f:24:f9:df:d0"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("mac"))
	})

	It("should fail to validate without field `MAC`", func() {
		m := models.MACValidate{
			// MAC: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("MAC"))
		Expect(errs[0].Tag()).To(Equal("mac"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.MACValidate{
			MAC: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("mac"))
	})

})
