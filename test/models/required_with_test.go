package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Required With", func() {
	Describe("Field", func() {
		When("String", func() {
			It("should to validate required field", func() {
				m := models.RequiredWithField{
					ID:     37669,
					Status: true,
					Name:   "Chico Bento",
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate required with field (Name)", func() {
				m := models.RequiredWithField{
					ID:     37669,
					Status: true,
					// Name:   "Chico Bento",
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("Name"))
				Expect(errs[0].Tag()).To(Equal("required_with"))
			})

			It("should fail to validate required with field (Status)", func() {
				m := models.RequiredWithField{
					ID:     37669,
					Status: false,
					Name:   "Chico Bento",
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("Status"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("String Pointer", func() {
			var name = "Chico Bento"

			It("should to validate required field", func() {
				m := models.RequiredWithFieldStrPointer{
					ID:     37669,
					Status: true,
					Name:   &name,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate required with field (Name)", func() {
				m := models.RequiredWithFieldStrPointer{
					ID:     37669,
					Status: true,
					// Name:   &name,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("Name"))
				Expect(errs[0].Tag()).To(Equal("required_with"))
			})

			It("should fail to validate required with field (Status)", func() {
				m := models.RequiredWithFieldStrPointer{
					ID:     37669,
					Status: false,
					Name:   &name,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("Status"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Two Fields", func() {
			var name = "Chico Bento"

			It("should to validate required field", func() {
				m := models.RequiredWithFields{
					ID:     1,
					Status: true,
					Name:   &name,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate required with field (Name)", func() {
				m := models.RequiredWithFields{
					ID:     1,
					Status: true,
					// Name:   &name,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("Name"))
				Expect(errs[0].Tag()).To(Equal("required_with"))
			})

			It("should fail to validate field not equal one (ID)", func() {
				m := models.RequiredWithFields{
					// ID:     1,
					Status: true,
					Name:   &name,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("ID"))
				Expect(errs[0].Tag()).To(Equal("eq"))
			})

			It("should fail to validate required with field (Status)", func() {
				m := models.RequiredWithFields{
					ID:     1,
					Status: false,
					Name:   &name,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("Status"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Required with url confirmation", func() {
			It("should to validate required field", func() {
				m := models.RequiredWithConfirmation{
					URL:              "https://gerson.name",
					NeedConfirmation: true,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate required with field (URL)", func() {
				m := models.RequiredWithConfirmation{
					// URL:              "https://gerson.name",
					NeedConfirmation: true,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("URL"))
				Expect(errs[0].Tag()).To(Equal("url"))
			})

			It("should required with field validation (NeedConfirmation)", func() {
				m := models.RequiredWithConfirmation{
					URL: "https://gerson.name",
					// NeedConfirmation: true,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("NeedConfirmation"))
				Expect(errs[0].Tag()).To(Equal("required_with"))
			})

			It("should fail to validate required with field (NeedConfirmation)", func() {
				m := models.RequiredWithConfirmation{
					URL: "invalid-url",
					// NeedConfirmation: true,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("URL"))
				Expect(errs[0].Tag()).To(Equal("url"))
			})
		})

	})
})