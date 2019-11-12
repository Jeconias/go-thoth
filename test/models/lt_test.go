package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lt", func() {

	It("should empty validation", func() {
		pwd := "tn"
		m := models.LtValidate{
			Name:     "Chico Bento",
			Age:      22,
			Password: &pwd,
			Contents: []string{"a", "b"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail to validate without field `Name`", func() {
		pwd := "tn"
		m := models.LtValidate{
			Name: "Chico Bento da Silva Junior",
			// Age:      22,
			Password: &pwd,
			Contents: []string{"a", "b"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Name"))
		Expect(errs[0].Tag()).To(Equal("lt"))
	})

	It("should fail to validate without field `Age`", func() {
		pwd := "tn"
		m := models.LtValidate{
			Name:     "Chico Bento",
			Age:      25,
			Password: &pwd,
			Contents: []string{"a", "b"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Age"))
		Expect(errs[0].Tag()).To(Equal("lt"))
	})

	It("should fail to validate without field `Password`", func() {
		m := models.LtValidate{
			Name: "Chico Bento",
			Age:  22,
			// Password: &pwd,
			Contents: []string{"a", "b"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Password"))
		Expect(errs[0].Tag()).To(Equal("lt"))
	})

	It("should fail to validate less than specified `Password`", func() {
		pwd := "abcde"
		m := models.LtValidate{
			Name:     "Chico Bento",
			Age:      22,
			Password: &pwd,
			Contents: []string{"a", "b"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Password"))
		Expect(errs[0].Tag()).To(Equal("lt"))
	})

	It("should to validate without field `Contents`", func() {
		pwd := "tn"
		m := models.LtValidate{
			Name:     "Chico Bento",
			Age:      22,
			Password: &pwd,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail to validate less than specified `Contents`", func() {
		pwd := "tn"
		m := models.LtValidate{
			Name:     "Chico Bento",
			Age:      22,
			Password: &pwd,
			Contents: []string{"a", "b", "c"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Contents"))
		Expect(errs[0].Tag()).To(Equal("lt"))
	})
})
