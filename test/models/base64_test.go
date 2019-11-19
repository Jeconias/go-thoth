package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Base64", func() {

	var (
		s     = "Q2hpY28gQmVudG8="
		slice = []string{"R3JhbmRlIGUgbm9icmUgY2hpcXVpbmhv", "SGVsbG8gd29ybGQ="}
	)

	It("should empty validation", func() {
		m := models.Base64Validate{
			Base64:  s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate base64 when base64 invalid", func() {
		ss := "http:base64:invalid-base64"
		m := models.Base64Validate{
			Base64:  s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("base64"))
	})

	It("should fail to validate without field `Base64`", func() {
		m := models.Base64Validate{
			// Base64: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Base64"))
		Expect(errs[0].Tag()).To(Equal("base64"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.Base64Validate{
			Base64: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("base64"))
	})

})
