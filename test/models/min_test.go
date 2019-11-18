package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Min", func() {

	It("should empty validation", func() {
		pwd := "password"
		m := models.MinValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"machine", "learning"},
			Attributes: map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail to validate field `Name`", func() {
		pwd := "password"
		m := models.MinValidate{
			Name:     "Ana",
			Age:      23,
			Password: &pwd,
			Contents: []string{"machine", "learning"},
			Attributes: map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Name"))
		Expect(errs[0].Tag()).To(Equal("min"))
	})

	It("should fail to validate without field `Age`", func() {
		pwd := "password"
		m := models.MinValidate{
			Name:     "Chico Bento",
			Age:      22,
			Password: &pwd,
			Contents: []string{"machine", "learning"},
			Attributes: map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Age"))
		Expect(errs[0].Tag()).To(Equal("min"))
	})

	It("should fail to validate without field `Password`", func() {
		// pwd := "password"
		m := models.MinValidate{
			Name: "Chico Bento",
			Age:  23,
			// Password: &pwd,
			Contents: []string{"machine", "learning"},
			Attributes: map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Password"))
		Expect(errs[0].Tag()).To(Equal("min"))
	})

	It("should fail to validate less than specified `Password`", func() {
		pwd := "tn"
		m := models.MinValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"machine", "learning"},
			Attributes: map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Password"))
		Expect(errs[0].Tag()).To(Equal("min"))
	})

	It("should fail to validate without field `Contents`", func() {
		pwd := "password"
		m := models.MinValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"machine"},
			Attributes: map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Contents"))
		Expect(errs[0].Tag()).To(Equal("min"))
	})

	It("should fail to validate less than specified `Contents`", func() {
		pwd := "password"
		m := models.MinValidate{
			Name:     "Chico Bento",
			Age:      23,
			Password: &pwd,
			Contents: []string{"machine", "learning"},
			Attributes: map[string]interface{}{
				"k1": "v1",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Attributes"))
		Expect(errs[0].Tag()).To(Equal("min"))
	})
})
