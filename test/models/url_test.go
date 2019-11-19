package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("URL", func() {

	var (
		s     = "http://lonny.org"
		slice = []string{"https://willis.com", "http://roberta.biz"}
	)

	It("should empty validation", func() {
		m := models.URLValidate{
			URL:     s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate url when url invalid", func() {
		ss := "invalid-url"
		m := models.URLValidate{
			URL:     s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("url"))
	})

	It("should fail validate url when Slice url is invalid", func() {
		m := models.URLValidate{
			URL:     s,
			Pointer: &s,
			Slice:   []string{"invalid-url", "http://roberta.biz"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("url"))
	})

	It("should fail to validate without field `URL`", func() {
		m := models.URLValidate{
			// URL: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("URL"))
		Expect(errs[0].Tag()).To(Equal("url"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.URLValidate{
			URL: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("url"))
	})

})
