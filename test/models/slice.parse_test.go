package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Slice", func() {
	Describe("Required", func() {
		When("String", func() {
			It("should empty validation", func() {
				el := "Chico Bento!"
				s := []string{el}
				sPtr := []*string{&el}

				m := models.TypeSliceString{
					SliceString:               s,
					PointerSliceString:        &s,
					SlicePointerString:        sPtr,
					PointerSlicePointerString: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceString`", func() {
				el := "Chico Bento!"
				s := []string{el}
				sPtr := []*string{&el}

				m := models.TypeSliceString{
					// SliceString:               s,
					PointerSliceString:        &s,
					SlicePointerString:        sPtr,
					PointerSlicePointerString: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceString`", func() {
				el := "Chico Bento!"
				s := []string{el}
				sPtr := []*string{&el}

				m := models.TypeSliceString{
					SliceString: s,
					// PointerSliceString:        &s,
					SlicePointerString:        sPtr,
					PointerSlicePointerString: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerString`", func() {
				el := "Chico Bento!"
				s := []string{el}
				sPtr := []*string{&el}

				m := models.TypeSliceString{
					SliceString:        s,
					PointerSliceString: &s,
					// SlicePointerString:        sPtr,
					PointerSlicePointerString: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerString`", func() {
				el := "Chico Bento!"
				s := []string{el}
				sPtr := []*string{&el}

				m := models.TypeSliceString{
					SliceString:        s,
					PointerSliceString: &s,
					SlicePointerString: sPtr,
					// PointerSlicePointerString: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Uint", func() {
			It("should empty validation", func() {
				var el uint = 1
				s := []uint{el}
				sPtr := []*uint{&el}

				m := models.TypeSliceUint{
					SliceUint:               s,
					PointerSliceUint:        &s,
					SlicePointerUint:        sPtr,
					PointerSlicePointerUint: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceUint`", func() {
				var el uint = 1
				s := []uint{el}
				sPtr := []*uint{&el}

				m := models.TypeSliceUint{
					// SliceUint:               s,
					PointerSliceUint:        &s,
					SlicePointerUint:        sPtr,
					PointerSlicePointerUint: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceUint"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceUint`", func() {
				var el uint = 1
				s := []uint{el}
				sPtr := []*uint{&el}

				m := models.TypeSliceUint{
					SliceUint: s,
					// PointerSliceUint:        &s,
					SlicePointerUint:        sPtr,
					PointerSlicePointerUint: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceUint"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerUint`", func() {
				var el uint = 1
				s := []uint{el}
				sPtr := []*uint{&el}

				m := models.TypeSliceUint{
					SliceUint:        s,
					PointerSliceUint: &s,
					// SlicePointerUint:        sPtr,
					PointerSlicePointerUint: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerUint"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerUint`", func() {
				var el uint = 1
				s := []uint{el}
				sPtr := []*uint{&el}

				m := models.TypeSliceUint{
					SliceUint:        s,
					PointerSliceUint: &s,
					SlicePointerUint: sPtr,
					// PointerSlicePointerUint: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerUint"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

	})
})
