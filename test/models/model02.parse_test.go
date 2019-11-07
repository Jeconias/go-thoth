package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Model02", func() {
	Describe("Required", func() {
		It("should to validate User (struct)", func() {
			m := models.User{
				ID:   8839,
				Age:  48527,
				Name: "149 Zelma Circles",
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should to validate User (struct) without id", func() {
			m := models.User{
				Age:  48527,
				Name: "149 Zelma Circles",
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate User (struct) without name", func() {
			m := models.User{
				ID:  8839,
				Age: 48527,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Name"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})
	})
})
