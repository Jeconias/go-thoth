package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interface", func() {
	Describe("Required", func() {
		It("should return empty validation (struct)", func() {
			var s interface{}
			s = models.MapTypeA{
				Int:  123,
				Bool: true,
			}

			m := models.TypeInterface{
				Interface:        s,
				PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should return empty validation", func() {
			var s interface{}
			s = "Chico Bento!"

			m := models.TypeInterface{
				Interface:        s,
				PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate without field `Interface`", func() {
			var s interface{}
			s = false
			m := models.TypeInterface{
				// Interface: s,
				PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Interface"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should fail to validate without field `PointerInterface`", func() {
			var s interface{}
			s = 222
			m := models.TypeInterface{
				Interface: s,
				// PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerInterface"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})
	})
})
