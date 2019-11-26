package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("URLEncoded", func() {

	var (
		s     = "lab259.rocks%3Fq%3DChico%20Bento"
		slice = []string{s}
	)

	It("should empty validation", func() {
		m := models.URLEncodedValidate{
			URLEncoded: s,
			Pointer:    &s,
			Slice:      slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate url when url invalid", func() {
		ss := "invalid-url"
		m := models.URLEncodedValidate{
			URLEncoded: s,
			Pointer:    &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("url_encoded"))
	})

	It("should fail validate url when Slice url is invalid", func() {
		m := models.URLEncodedValidate{
			URLEncoded: s,
			Pointer:    &s,
			Slice:      []string{"invalid-url", "https%3A%2F%2Ffernanda.com.br%3Fq%3DChico%20bento"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("url_encoded"))
	})

	It("should fail to validate without field `URLEncoded`", func() {
		m := models.URLEncodedValidate{
			// URLEncoded: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("URLEncoded"))
		Expect(errs[0].Tag()).To(Equal("url_encoded"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.URLEncodedValidate{
			URLEncoded: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("url_encoded"))
	})

})
