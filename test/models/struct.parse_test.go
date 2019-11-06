package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Struct", func() {
	Describe("Required", func() {
		It("should fail to validate field `Struct`", func() {
			structA := models.StructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.StructA{structA}
			sliceStructAPtr := []*models.StructA{&structA}
			t := models.Struct{
				// Struct: structA,
				StructPointer:             &structA,
				SliceStruct:               sliceStructA,
				SliceStructPointer:        sliceStructAPtr,
				SlicePointerStruct:        &sliceStructA,
				SlicePointerStructPointer: &sliceStructAPtr,
			}

			errs := t.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Struct"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should check if field `StructPointer`", func() {
			structA := models.StructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.StructA{structA}
			sliceStructAPtr := []*models.StructA{&structA}
			t := models.Struct{
				Struct: structA,
				// StructPointer: &structA,
				SliceStruct:               sliceStructA,
				SliceStructPointer:        sliceStructAPtr,
				SlicePointerStruct:        &sliceStructA,
				SlicePointerStructPointer: &sliceStructAPtr,
			}

			errs := t.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("StructPointer"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should check if field `SliceStruct`", func() {
			structA := models.StructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.StructA{structA}
			sliceStructAPtr := []*models.StructA{&structA}
			t := models.Struct{
				Struct:        structA,
				StructPointer: &structA,
				// SliceStruct:               sliceStructA,
				SliceStructPointer:        sliceStructAPtr,
				SlicePointerStruct:        &sliceStructA,
				SlicePointerStructPointer: &sliceStructAPtr,
			}

			errs := t.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("SliceStruct"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should check if field `SliceStructPointer`", func() {
			structA := models.StructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.StructA{structA}
			sliceStructAPtr := []*models.StructA{&structA}
			t := models.Struct{
				Struct:        structA,
				StructPointer: &structA,
				SliceStruct:   sliceStructA,
				// SliceStructPointer:        sliceStructAPtr,
				SlicePointerStruct:        &sliceStructA,
				SlicePointerStructPointer: &sliceStructAPtr,
			}

			errs := t.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("SliceStructPointer"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should check if field `SlicePointerStruct`", func() {
			structA := models.StructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.StructA{structA}
			sliceStructAPtr := []*models.StructA{&structA}
			t := models.Struct{
				Struct:             structA,
				StructPointer:      &structA,
				SliceStruct:        sliceStructA,
				SliceStructPointer: sliceStructAPtr,
				// SlicePointerStruct:        &sliceStructA,
				SlicePointerStructPointer: &sliceStructAPtr,
			}

			errs := t.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("SlicePointerStruct"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should check if field `SlicePointerStructPointer`", func() {
			structA := models.StructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.StructA{structA}
			sliceStructAPtr := []*models.StructA{&structA}
			t := models.Struct{
				Struct:             structA,
				StructPointer:      &structA,
				SliceStruct:        sliceStructA,
				SliceStructPointer: sliceStructAPtr,
				SlicePointerStruct: &sliceStructA,
				// SlicePointerStructPointer: &sliceStructAPtr,
			}

			errs := t.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("SlicePointerStructPointer"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})
	})
})
