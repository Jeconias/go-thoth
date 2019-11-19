package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("URI", func() {

	var (
		s     = "ftp://example.org/resource.txt"
		slice = []string{"http://example.org/absolute/URI/with/absolute/path/to/resource.txt", "ftp://example.org/resource.txt"}
	)

	It("should empty validation", func() {
		m := models.URIValidate{
			URI:     s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate uri when uri invalid", func() {
		ss := "invalid-uri"
		m := models.URIValidate{
			URI:     s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("uri"))
	})

	It("should fail to validate without field `URI`", func() {
		m := models.URIValidate{
			// URI: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("URI"))
		Expect(errs[0].Tag()).To(Equal("uri"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.URIValidate{
			URI: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("uri"))
	})

})