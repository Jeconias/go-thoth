package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Email", func() {

	It("should empty validation", func() {
		s := "chicobento@lab259.rocks"
		m := models.EmailValidation{
			Email:   s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate email when email invalid", func() {
		s := "invalid-email"
		m := models.EmailValidation{
			Email:   "chicobento@lab259.rocks",
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("email"))
	})

	It("should fail to validate without field `Email`", func() {
		s := "chicobento@lab259.rocks"
		m := models.EmailValidation{
			// Email: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Email"))
		Expect(errs[0].Tag()).To(Equal("email"))
	})

	It("should fail to validater without field `Pointer`", func() {
		s := "chicobento@lab259.rocks"
		m := models.EmailValidation{
			Email: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("email"))
	})
})
