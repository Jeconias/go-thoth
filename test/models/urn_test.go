package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("URN", func() {

	var (
		s     = "urn:isbn:0451450523"
		slice = []string{"urn:lex:br:federal:lei:2008-06-19;11705", "urn:issn:0167-6423"}
	)

	It("should empty validation", func() {
		m := models.URNValidate{
			URN:     s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate urn when urn invalid", func() {
		ss := "http:urn:invalid-urn"
		m := models.URNValidate{
			URN:     s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("urn_rfc2141"))
	})

	It("should fail validate urn when Slice urn invalid", func() {
		m := models.URNValidate{
			URN:     s,
			Pointer: &s,
			Slice: []string{
				"urn:lex:br:federal:lei:2008-06-19;11705",
				"urn:issn:0167-6423",
				"http:urn:invalid-urn",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("urn_rfc2141"))
	})

	It("should fail to validate without field `URN`", func() {
		m := models.URNValidate{
			// URN: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("URN"))
		Expect(errs[0].Tag()).To(Equal("urn_rfc2141"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.URNValidate{
			URN: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("urn_rfc2141"))
	})

})
