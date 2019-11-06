package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String", func() {
	Describe("Required", func() {
		It("should empty validation", func() {
			s := "Chico Bento!"
			m := models.TypeString{
				String:  s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should check if field `String`", func() {
			s := "Chico Bento!"
			m := models.TypeString{
				// String: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("String"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should check if field `String Pointer`", func() {
			s := "Chico Bento!"
			m := models.TypeString{
				String: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})
	})
})
