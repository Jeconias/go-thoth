package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSN", func() {

	var (
		s     = "043-16-9406"
		slice = []string{s}
	)

	It("should empty validation", func() {
		m := models.SSNValidate{
			SSN:     s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate ssn when ssn invalid", func() {
		ss := "011234567"
		m := models.SSNValidate{
			SSN:     s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("ssn"))
	})

	It("should fail validate ssn when Slice ssn is invalid", func() {
		m := models.SSNValidate{
			SSN:     s,
			Pointer: &s,
			Slice:   []string{"011234567", s},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("ssn"))
	})

	It("should fail to validate without field `SSN`", func() {
		m := models.SSNValidate{
			// SSN: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("SSN"))
		Expect(errs[0].Tag()).To(Equal("ssn"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.SSNValidate{
			SSN: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("ssn"))
	})

})
