package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FQDN", func() {

	var (
		s     = "somecollege.edu"
		slice = []string{"www.indiana.edu", "mymail.somecollege.edu"}
	)

	It("should empty validation", func() {
		m := models.FQDNValidate{
			FQDN:    s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate fqdn when fqdn invalid", func() {
		ss := "invalid-fqdn"
		m := models.FQDNValidate{
			FQDN:    s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("fqdn"))
	})

	It("should fail validate fqdn when Slice fqdn is invalid", func() {
		m := models.FQDNValidate{
			FQDN:    s,
			Pointer: &s,
			Slice:   []string{"somecollege.edu", "A123"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("fqdn"))
	})

	It("should fail to validate without field `FQDN`", func() {
		m := models.FQDNValidate{
			// FQDN: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("FQDN"))
		Expect(errs[0].Tag()).To(Equal("fqdn"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.FQDNValidate{
			FQDN: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("fqdn"))
	})

})
