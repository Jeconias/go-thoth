package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Base64url", func() {

	var (
		s     = "VGhpcyBpcyBzaW1wbGUgQVNDSUkgQmFzZTY0IGZvciBTdGFja092ZXJmbG93IGV4YW1wbGUu"
		slice = []string{"VGhpcyBpcyBzaW1wbGUgQVNDSUkgQmFzZTY0IGZvciBTdGFja092ZXJmbG93IGV4YW1wbGUu"}
	)

	It("should empty validation", func() {
		m := models.Base64urlValidate{
			Base64url: s,
			Pointer:   &s,
			Slice:     slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate base64url when base64url invalid", func() {
		ss := "http:base64url:invalid-base64url"
		m := models.Base64urlValidate{
			Base64url: s,
			Pointer:   &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("base64url"))
	})

	It("should fail to validate without field `Base64url`", func() {
		m := models.Base64urlValidate{
			// Base64url: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Base64url"))
		Expect(errs[0].Tag()).To(Equal("base64url"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.Base64urlValidate{
			Base64url: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("base64url"))
	})

})
