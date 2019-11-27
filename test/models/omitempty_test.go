package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Omitempty", func() {

	var (
		s        = "Chico"
		strSlice = []string{s}
	)

	It("should empty validation", func() {
		m := models.OmitemptyValidate{
			OmitEmptyString:      s,
			OmitEmptyInt:         5,
			OmitEmptySliceString: strSlice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should empty validation without `OmitEmptyString`", func() {
		m := models.OmitemptyValidate{
			// OmitEmptyString:      s,
			OmitEmptyInt:         5,
			OmitEmptySliceString: strSlice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should empty validation without `OmitEmptyInt`", func() {
		m := models.OmitemptyValidate{
			OmitEmptyString: s,
			// OmitEmptyInt:         5,
			OmitEmptySliceString: strSlice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should empty validation without `OmitEmptySliceString`", func() {
		m := models.OmitemptyValidate{
			OmitEmptyString: s,
			OmitEmptyInt:    5,
			// OmitEmptySliceString: strSlice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail to validate when field `OmitEmptyString` failing max", func() {
		m := models.OmitemptyValidate{
			OmitEmptyString:      "Chico Bento",
			OmitEmptyInt:         5,
			OmitEmptySliceString: strSlice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Tag()).To(Equal("[min = 5] ~ [max = 5]"))
	})

	It("should fail to validate when field `OmitEmptyString` failing min", func() {
		m := models.OmitemptyValidate{
			OmitEmptyString:      s,
			OmitEmptyInt:         15,
			OmitEmptySliceString: strSlice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Tag()).To(Equal("[min = 5] ~ [max = 5]"))
	})

	It("should fail to validate when field `OmitEmptyString` failing min", func() {
		m := models.OmitemptyValidate{
			OmitEmptyString:      s,
			OmitEmptyInt:         5,
			OmitEmptySliceString: []string{s, "Bento Junior"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Tag()).To(Equal("[min = 1]"))
	})

	It("should fail to validate when field `OmitEmptySliceInt` not is empty", func() {
		m := models.OmitemptyValidate{
			OmitEmptyString:      s,
			OmitEmptyInt:         5,
			OmitEmptySliceString: strSlice,
			OmitEmptySliceInt:    []int{1, 23, 4},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Tag()).To(Equal("omitempty"))
	})

})
