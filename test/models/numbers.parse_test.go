package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Numbers", func() {
	Describe("Required", func() {
		It("should fail to validate all numbers", func() {
			m := models.Number{}

			errs := m.Validate()
			Expect(errs).To(HaveLen(30))
		})
	})
})
