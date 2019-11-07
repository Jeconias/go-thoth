package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Model01", func() {
	Describe("Required", func() {
		It("should to validate Home (struct)", func() {
			m := models.Home{
				ID: 8839,
				Address: models.Address{
					ID:     48527,
					Street: "149 Zelma Circles",
				},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate Home (struct) without id", func() {
			m := models.Home{
				Address: models.Address{
					ID:     48527,
					Street: "149 Zelma Circles",
				},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ID"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should fail to validate Home (struct) without address", func() {
			m := models.Home{
				ID: 48527,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Address"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should to validate Client (struct)", func() {
			lastName := "Hirthe"

			m := models.Client{
				ID:       8839,
				Name:     "Kendall",
				LastName: &lastName,
				Address: &models.Address{
					ID:     48527,
					Street: "149 Zelma Circles",
				},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate Client (struct) without id", func() {
			lastName := "Kendall"

			m := models.Client{
				Name:     "Kendall",
				LastName: &lastName,
				Address: &models.Address{
					ID:     48527,
					Street: "149 Zelma Circles",
				},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ID"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should fail to validate Client (struct) without name", func() {
			lastName := "Kendall"

			m := models.Client{
				ID:       8839,
				LastName: &lastName,
				Address: &models.Address{
					ID:     48527,
					Street: "149 Zelma Circles",
				},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Name"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should fail to validate Client (struct) without lastname", func() {
			m := models.Client{
				ID:   8839,
				Name: "Kendall",
				Address: &models.Address{
					ID:     48527,
					Street: "149 Zelma Circles",
				},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("LastName"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should fail to validate Client (struct) without address", func() {
			lastName := "Kendall"

			m := models.Client{
				ID:       8839,
				Name:     "Kendall",
				LastName: &lastName,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Address"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})
	})
})
