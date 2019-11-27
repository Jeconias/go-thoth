package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dive", func() {

	var (
		strSlice    = []string{"Chico"}
		intSlice    = []int{1, 2, 3}
		mapStrToStr = map[string]string{
			"k": "v",
		}
	)

	It("should empty validation", func() {
		m := models.DiveValidate{
			SliceString:       strSlice,
			SliceInt:          intSlice,
			MapStringToString: mapStrToStr,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail to validation when empty string", func() {
		m := models.DiveValidate{
			SliceString:       []string{"Chico", "", "Bento"},
			SliceInt:          intSlice,
			MapStringToString: mapStrToStr,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("SliceString"))
		Expect(errs[0].Tag()).To(Equal("dive"))
	})

	It("should fail to validation when number is zero", func() {
		m := models.DiveValidate{
			SliceString:       strSlice,
			SliceInt:          []int{1, 0, 2, 3},
			MapStringToString: mapStrToStr,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("SliceInt"))
		Expect(errs[0].Tag()).To(Equal("dive"))
	})

	It("should fail to validation when map is zero", func() {
		m := models.DiveValidate{
			SliceString: strSlice,
			SliceInt:    intSlice,
			MapStringToString: map[string]string{
				"k":  "v",
				"":   "",
				"k2": "v2",
			},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("MapStringToString"))
		Expect(errs[0].Tag()).To(Equal("dive"))
	})

})
