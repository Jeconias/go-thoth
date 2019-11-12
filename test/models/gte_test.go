package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gte", func() {

	It("should empty validation", func() {
		pwd := "tnBAZXZ3CnT9EJy"
		m := models.GteValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"a", "b", "c"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail to validate without field `Name`", func() {
		pwd := "tnBAZXZ3CnT9EJy"
		m := models.GteValidate{
			// Name: "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"a", "b", "c"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Name"))
		Expect(errs[0].Tag()).To(Equal("gte"))
	})

	It("should fail to validate without field `Age`", func() {
		pwd := "tnBAZXZ3CnT9EJy"
		m := models.GteValidate{
			Name: "Chico Bento",
			// Age:      23,
			Password: &pwd,
			Contents: []string{"a", "b", "c"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Age"))
		Expect(errs[0].Tag()).To(Equal("gte"))
	})

	It("should fail to validate without field `Password`", func() {
		m := models.GteValidate{
			Name: "Chico Bento",
			Age:  23,
			// Password: &pwd,
			Contents: []string{"a", "b", "c"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Password"))
		Expect(errs[0].Tag()).To(Equal("gte"))
	})

	It("should fail to validate less than specified `Password`", func() {
		pwd := "ab"
		m := models.GteValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"a", "b", "c"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Password"))
		Expect(errs[0].Tag()).To(Equal("gte"))
	})

	It("should fail to validate without field `Contents`", func() {
		pwd := "tnBAZXZ3CnT9EJy"
		m := models.GteValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			// Contents: []string{"a", "b", "c"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Contents"))
		Expect(errs[0].Tag()).To(Equal("gte"))
	})

	It("should fail to validate less than specified `Contents`", func() {
		pwd := "tnBAZXZ3CnT9EJy"
		m := models.GteValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"a"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Contents"))
		Expect(errs[0].Tag()).To(Equal("gte"))
	})
})
