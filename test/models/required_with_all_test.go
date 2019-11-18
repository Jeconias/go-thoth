package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Required With All", func() {
	When("String", func() {
		It("should to validate required field", func() {
			m := models.RequiredWithAllField{
				Status: true,
				Name:   "Chico Bento",
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate required with field (Name)", func() {
			m := models.RequiredWithAllField{
				Status: true,
				// Name:   "Chico Bento",
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Name"))
			Expect(errs[0].Tag()).To(Equal("required_with_all"))
		})

		It("should fail to validate required with field (Status)", func() {
			m := models.RequiredWithAllField{
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
			m := models.RequiredWithAllFieldStrPointer{
				Status: true,
				Name:   &name,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate required with field (Name)", func() {
			m := models.RequiredWithAllFieldStrPointer{
				Status: true,
				// Name:   &name,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Name"))
			Expect(errs[0].Tag()).To(Equal("required_with_all"))
		})

		It("should fail to validate required with field (Status)", func() {
			m := models.RequiredWithAllFieldStrPointer{
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
			m := models.RequiredWithAllFields{
				ID:     1,
				Status: true,
				Name:   &name,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate required with field (Name)", func() {
			m := models.RequiredWithAllFields{
				ID:     1,
				Status: true,
				// Name:   &name,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Name"))
			Expect(errs[0].Tag()).To(Equal("required_with_all"))
		})

		It("should fail to validate field not equal one (ID)", func() {
			m := models.RequiredWithAllFields{
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
			m := models.RequiredWithAllFields{
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
})
