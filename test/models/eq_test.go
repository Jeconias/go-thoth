package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Equal", func() {
	When("Numbers", func() {

		var (
			uintPointer       uint       = 2
			uint8Pointer      uint8      = 4
			uint16Pointer     uint16     = 6
			uint32Pointer     uint32     = 8
			uint64Pointer     uint64     = 10
			uintptrPointer    uintptr    = 12
			intPointer        int        = 14
			int8Pointer       int8       = 16
			int16Pointer      int16      = 18
			int32Pointer      int32      = 20
			int64Pointer      int64      = 22
			float32Pointer    float32    = 24
			float64Pointer    float64    = 26
			complex64Pointer  complex64  = 28
			complex128Pointer complex128 = 30
		)

		It("should fail to validate all numbers", func() {
			m := models.EqNumber{}

			errs := m.Validate()
			Expect(errs).To(HaveLen(30))
			Expect(errs[0].Field()).To(Equal("Uint"))
			Expect(errs[0].Tag()).To(Equal("eq"))
		})

		It("should fail to validate field (Uint)", func() {
			m := models.EqNumber{
				// Uint:           1,
				UintPointer:    &uintPointer,
				Uint8:          3,
				Uint8Pointer:   &uint8Pointer,
				Uint16:         5,
				Uint16Pointer:  &uint16Pointer,
				Uint32:         7,
				Uint32Pointer:  &uint32Pointer,
				Uint64:         9,
				Uint64Pointer:  &uint64Pointer,
				Uintptr:        11,
				UintptrPointer: &uintptrPointer,

				Int:          13,
				IntPointer:   &intPointer,
				Int8:         15,
				Int8Pointer:  &int8Pointer,
				Int16:        17,
				Int16Pointer: &int16Pointer,
				Int32:        19,
				Int32Pointer: &int32Pointer,
				Int64:        21,
				Int64Pointer: &int64Pointer,

				Float32:        23,
				Float32Pointer: &float32Pointer,
				Float64:        25,
				Float64Pointer: &float64Pointer,

				Complex64:         27,
				Complex64Pointer:  &complex64Pointer,
				Complex128:        29,
				Complex128Pointer: &complex128Pointer,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Uint"))
			Expect(errs[0].Tag()).To(Equal("eq"))
		})

		It("should fail to validate field (UintPointer)", func() {
			m := models.EqNumber{
				Uint: 1,
				// UintPointer:    &uintPointer,
				Uint8:          3,
				Uint8Pointer:   &uint8Pointer,
				Uint16:         5,
				Uint16Pointer:  &uint16Pointer,
				Uint32:         7,
				Uint32Pointer:  &uint32Pointer,
				Uint64:         9,
				Uint64Pointer:  &uint64Pointer,
				Uintptr:        11,
				UintptrPointer: &uintptrPointer,

				Int:          13,
				IntPointer:   &intPointer,
				Int8:         15,
				Int8Pointer:  &int8Pointer,
				Int16:        17,
				Int16Pointer: &int16Pointer,
				Int32:        19,
				Int32Pointer: &int32Pointer,
				Int64:        21,
				Int64Pointer: &int64Pointer,

				Float32:        23,
				Float32Pointer: &float32Pointer,
				Float64:        25,
				Float64Pointer: &float64Pointer,

				Complex64:         27,
				Complex64Pointer:  &complex64Pointer,
				Complex128:        29,
				Complex128Pointer: &complex128Pointer,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UintPointer"))
			Expect(errs[0].Tag()).To(Equal("eq"))
		})
	})

	When("String", func() {
		It("should empty validation", func() {
			s := "bento"
			m := models.TypeEqString{
				String:  "chico",
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should to validate field `String`", func() {
			s := "bento"
			m := models.TypeEqString{
				// String: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("String"))
			Expect(errs[0].Tag()).To(Equal("eq"))
		})

		It("should check if field `String Pointer`", func() {
			s := "chico"
			m := models.TypeEqString{
				String: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("eq"))
		})
	})
})
