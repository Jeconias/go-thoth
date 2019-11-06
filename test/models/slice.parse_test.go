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

		When("Uint8", func() {
			It("should empty validation", func() {
				var el uint8 = 1
				s := []uint8{el}
				sPtr := []*uint8{&el}

				m := models.TypeSliceUint8{
					SliceUint8:               s,
					PointerSliceUint8:        &s,
					SlicePointerUint8:        sPtr,
					PointerSlicePointerUint8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceUint8`", func() {
				var el uint8 = 1
				s := []uint8{el}
				sPtr := []*uint8{&el}

				m := models.TypeSliceUint8{
					// SliceUint8:               s,
					PointerSliceUint8:        &s,
					SlicePointerUint8:        sPtr,
					PointerSlicePointerUint8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceUint8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceUint8`", func() {
				var el uint8 = 1
				s := []uint8{el}
				sPtr := []*uint8{&el}

				m := models.TypeSliceUint8{
					SliceUint8: s,
					// PointerSliceUint8:        &s,
					SlicePointerUint8:        sPtr,
					PointerSlicePointerUint8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceUint8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerUint8`", func() {
				var el uint8 = 1
				s := []uint8{el}
				sPtr := []*uint8{&el}

				m := models.TypeSliceUint8{
					SliceUint8:        s,
					PointerSliceUint8: &s,
					// SlicePointerUint8:        sPtr,
					PointerSlicePointerUint8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerUint8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerUint8`", func() {
				var el uint8 = 1
				s := []uint8{el}
				sPtr := []*uint8{&el}

				m := models.TypeSliceUint8{
					SliceUint8:        s,
					PointerSliceUint8: &s,
					SlicePointerUint8: sPtr,
					// PointerSlicePointerUint8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerUint8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Uint16", func() {
			It("should empty validation", func() {
				var el uint16 = 1
				s := []uint16{el}
				sPtr := []*uint16{&el}

				m := models.TypeSliceUint16{
					SliceUint16:               s,
					PointerSliceUint16:        &s,
					SlicePointerUint16:        sPtr,
					PointerSlicePointerUint16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceUint16`", func() {
				var el uint16 = 1
				s := []uint16{el}
				sPtr := []*uint16{&el}

				m := models.TypeSliceUint16{
					// SliceUint16:               s,
					PointerSliceUint16:        &s,
					SlicePointerUint16:        sPtr,
					PointerSlicePointerUint16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceUint16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceUint16`", func() {
				var el uint16 = 1
				s := []uint16{el}
				sPtr := []*uint16{&el}

				m := models.TypeSliceUint16{
					SliceUint16: s,
					// PointerSliceUint16:        &s,
					SlicePointerUint16:        sPtr,
					PointerSlicePointerUint16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceUint16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerUint16`", func() {
				var el uint16 = 1
				s := []uint16{el}
				sPtr := []*uint16{&el}

				m := models.TypeSliceUint16{
					SliceUint16:        s,
					PointerSliceUint16: &s,
					// SlicePointerUint16:        sPtr,
					PointerSlicePointerUint16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerUint16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerUint16`", func() {
				var el uint16 = 1
				s := []uint16{el}
				sPtr := []*uint16{&el}

				m := models.TypeSliceUint16{
					SliceUint16:        s,
					PointerSliceUint16: &s,
					SlicePointerUint16: sPtr,
					// PointerSlicePointerUint16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerUint16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Uint32", func() {
			It("should empty validation", func() {
				var el uint32 = 1
				s := []uint32{el}
				sPtr := []*uint32{&el}

				m := models.TypeSliceUint32{
					SliceUint32:               s,
					PointerSliceUint32:        &s,
					SlicePointerUint32:        sPtr,
					PointerSlicePointerUint32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceUint32`", func() {
				var el uint32 = 1
				s := []uint32{el}
				sPtr := []*uint32{&el}

				m := models.TypeSliceUint32{
					// SliceUint32:               s,
					PointerSliceUint32:        &s,
					SlicePointerUint32:        sPtr,
					PointerSlicePointerUint32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceUint32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceUint32`", func() {
				var el uint32 = 1
				s := []uint32{el}
				sPtr := []*uint32{&el}

				m := models.TypeSliceUint32{
					SliceUint32: s,
					// PointerSliceUint32:        &s,
					SlicePointerUint32:        sPtr,
					PointerSlicePointerUint32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceUint32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerUint32`", func() {
				var el uint32 = 1
				s := []uint32{el}
				sPtr := []*uint32{&el}

				m := models.TypeSliceUint32{
					SliceUint32:        s,
					PointerSliceUint32: &s,
					// SlicePointerUint32:        sPtr,
					PointerSlicePointerUint32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerUint32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerUint32`", func() {
				var el uint32 = 1
				s := []uint32{el}
				sPtr := []*uint32{&el}

				m := models.TypeSliceUint32{
					SliceUint32:        s,
					PointerSliceUint32: &s,
					SlicePointerUint32: sPtr,
					// PointerSlicePointerUint32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerUint32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Uint64", func() {
			It("should empty validation", func() {
				var el uint64 = 1
				s := []uint64{el}
				sPtr := []*uint64{&el}

				m := models.TypeSliceUint64{
					SliceUint64:               s,
					PointerSliceUint64:        &s,
					SlicePointerUint64:        sPtr,
					PointerSlicePointerUint64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceUint64`", func() {
				var el uint64 = 1
				s := []uint64{el}
				sPtr := []*uint64{&el}

				m := models.TypeSliceUint64{
					// SliceUint64:               s,
					PointerSliceUint64:        &s,
					SlicePointerUint64:        sPtr,
					PointerSlicePointerUint64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceUint64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceUint64`", func() {
				var el uint64 = 1
				s := []uint64{el}
				sPtr := []*uint64{&el}

				m := models.TypeSliceUint64{
					SliceUint64: s,
					// PointerSliceUint64:        &s,
					SlicePointerUint64:        sPtr,
					PointerSlicePointerUint64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceUint64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerUint64`", func() {
				var el uint64 = 1
				s := []uint64{el}
				sPtr := []*uint64{&el}

				m := models.TypeSliceUint64{
					SliceUint64:        s,
					PointerSliceUint64: &s,
					// SlicePointerUint64:        sPtr,
					PointerSlicePointerUint64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerUint64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerUint64`", func() {
				var el uint64 = 1
				s := []uint64{el}
				sPtr := []*uint64{&el}

				m := models.TypeSliceUint64{
					SliceUint64:        s,
					PointerSliceUint64: &s,
					SlicePointerUint64: sPtr,
					// PointerSlicePointerUint64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerUint64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Uintptr", func() {
			It("should empty validation", func() {
				var el uintptr = 1
				s := []uintptr{el}
				sPtr := []*uintptr{&el}

				m := models.TypeSliceUintptr{
					SliceUintptr:               s,
					PointerSliceUintptr:        &s,
					SlicePointerUintptr:        sPtr,
					PointerSlicePointerUintptr: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceUintptr`", func() {
				var el uintptr = 1
				s := []uintptr{el}
				sPtr := []*uintptr{&el}

				m := models.TypeSliceUintptr{
					// SliceUintptr:               s,
					PointerSliceUintptr:        &s,
					SlicePointerUintptr:        sPtr,
					PointerSlicePointerUintptr: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceUintptr"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceUintptr`", func() {
				var el uintptr = 1
				s := []uintptr{el}
				sPtr := []*uintptr{&el}

				m := models.TypeSliceUintptr{
					SliceUintptr: s,
					// PointerSliceUintptr:        &s,
					SlicePointerUintptr:        sPtr,
					PointerSlicePointerUintptr: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceUintptr"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerUintptr`", func() {
				var el uintptr = 1
				s := []uintptr{el}
				sPtr := []*uintptr{&el}

				m := models.TypeSliceUintptr{
					SliceUintptr:        s,
					PointerSliceUintptr: &s,
					// SlicePointerUintptr:        sPtr,
					PointerSlicePointerUintptr: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerUintptr"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerUintptr`", func() {
				var el uintptr = 1
				s := []uintptr{el}
				sPtr := []*uintptr{&el}

				m := models.TypeSliceUintptr{
					SliceUintptr:        s,
					PointerSliceUintptr: &s,
					SlicePointerUintptr: sPtr,
					// PointerSlicePointerUintptr: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerUintptr"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Int", func() {
			It("should empty validation", func() {
				var el int = 1
				s := []int{el}
				sPtr := []*int{&el}

				m := models.TypeSliceInt{
					SliceInt:               s,
					PointerSliceInt:        &s,
					SlicePointerInt:        sPtr,
					PointerSlicePointerInt: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceInt`", func() {
				var el int = 1
				s := []int{el}
				sPtr := []*int{&el}

				m := models.TypeSliceInt{
					// SliceInt:               s,
					PointerSliceInt:        &s,
					SlicePointerInt:        sPtr,
					PointerSlicePointerInt: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceInt"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceInt`", func() {
				var el int = 1
				s := []int{el}
				sPtr := []*int{&el}

				m := models.TypeSliceInt{
					SliceInt: s,
					// PointerSliceInt:        &s,
					SlicePointerInt:        sPtr,
					PointerSlicePointerInt: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceInt"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerInt`", func() {
				var el int = 1
				s := []int{el}
				sPtr := []*int{&el}

				m := models.TypeSliceInt{
					SliceInt:        s,
					PointerSliceInt: &s,
					// SlicePointerInt:        sPtr,
					PointerSlicePointerInt: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerInt"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerInt`", func() {
				var el int = 1
				s := []int{el}
				sPtr := []*int{&el}

				m := models.TypeSliceInt{
					SliceInt:        s,
					PointerSliceInt: &s,
					SlicePointerInt: sPtr,
					// PointerSlicePointerInt: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerInt"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Int8", func() {
			It("should empty validation", func() {
				var el int8 = 1
				s := []int8{el}
				sPtr := []*int8{&el}

				m := models.TypeSliceInt8{
					SliceInt8:               s,
					PointerSliceInt8:        &s,
					SlicePointerInt8:        sPtr,
					PointerSlicePointerInt8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceInt8`", func() {
				var el int8 = 1
				s := []int8{el}
				sPtr := []*int8{&el}

				m := models.TypeSliceInt8{
					// SliceInt8:               s,
					PointerSliceInt8:        &s,
					SlicePointerInt8:        sPtr,
					PointerSlicePointerInt8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceInt8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceInt8`", func() {
				var el int8 = 1
				s := []int8{el}
				sPtr := []*int8{&el}

				m := models.TypeSliceInt8{
					SliceInt8: s,
					// PointerSliceInt8:        &s,
					SlicePointerInt8:        sPtr,
					PointerSlicePointerInt8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceInt8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerInt8`", func() {
				var el int8 = 1
				s := []int8{el}
				sPtr := []*int8{&el}

				m := models.TypeSliceInt8{
					SliceInt8:        s,
					PointerSliceInt8: &s,
					// SlicePointerInt8:        sPtr,
					PointerSlicePointerInt8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerInt8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerInt8`", func() {
				var el int8 = 1
				s := []int8{el}
				sPtr := []*int8{&el}

				m := models.TypeSliceInt8{
					SliceInt8:        s,
					PointerSliceInt8: &s,
					SlicePointerInt8: sPtr,
					// PointerSlicePointerInt8: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerInt8"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Int16", func() {
			It("should empty validation", func() {
				var el int16 = 1
				s := []int16{el}
				sPtr := []*int16{&el}

				m := models.TypeSliceInt16{
					SliceInt16:               s,
					PointerSliceInt16:        &s,
					SlicePointerInt16:        sPtr,
					PointerSlicePointerInt16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceInt16`", func() {
				var el int16 = 1
				s := []int16{el}
				sPtr := []*int16{&el}

				m := models.TypeSliceInt16{
					// SliceInt16:               s,
					PointerSliceInt16:        &s,
					SlicePointerInt16:        sPtr,
					PointerSlicePointerInt16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceInt16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceInt16`", func() {
				var el int16 = 1
				s := []int16{el}
				sPtr := []*int16{&el}

				m := models.TypeSliceInt16{
					SliceInt16: s,
					// PointerSliceInt16:        &s,
					SlicePointerInt16:        sPtr,
					PointerSlicePointerInt16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceInt16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerInt16`", func() {
				var el int16 = 1
				s := []int16{el}
				sPtr := []*int16{&el}

				m := models.TypeSliceInt16{
					SliceInt16:        s,
					PointerSliceInt16: &s,
					// SlicePointerInt16:        sPtr,
					PointerSlicePointerInt16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerInt16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerInt16`", func() {
				var el int16 = 1
				s := []int16{el}
				sPtr := []*int16{&el}

				m := models.TypeSliceInt16{
					SliceInt16:        s,
					PointerSliceInt16: &s,
					SlicePointerInt16: sPtr,
					// PointerSlicePointerInt16: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerInt16"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Int32", func() {
			It("should empty validation", func() {
				var el int32 = 1
				s := []int32{el}
				sPtr := []*int32{&el}

				m := models.TypeSliceInt32{
					SliceInt32:               s,
					PointerSliceInt32:        &s,
					SlicePointerInt32:        sPtr,
					PointerSlicePointerInt32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceInt32`", func() {
				var el int32 = 1
				s := []int32{el}
				sPtr := []*int32{&el}

				m := models.TypeSliceInt32{
					// SliceInt32:               s,
					PointerSliceInt32:        &s,
					SlicePointerInt32:        sPtr,
					PointerSlicePointerInt32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceInt32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceInt32`", func() {
				var el int32 = 1
				s := []int32{el}
				sPtr := []*int32{&el}

				m := models.TypeSliceInt32{
					SliceInt32: s,
					// PointerSliceInt32:        &s,
					SlicePointerInt32:        sPtr,
					PointerSlicePointerInt32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceInt32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerInt32`", func() {
				var el int32 = 1
				s := []int32{el}
				sPtr := []*int32{&el}

				m := models.TypeSliceInt32{
					SliceInt32:        s,
					PointerSliceInt32: &s,
					// SlicePointerInt32:        sPtr,
					PointerSlicePointerInt32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerInt32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerInt32`", func() {
				var el int32 = 1
				s := []int32{el}
				sPtr := []*int32{&el}

				m := models.TypeSliceInt32{
					SliceInt32:        s,
					PointerSliceInt32: &s,
					SlicePointerInt32: sPtr,
					// PointerSlicePointerInt32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerInt32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Int64", func() {
			It("should empty validation", func() {
				var el int64 = 1
				s := []int64{el}
				sPtr := []*int64{&el}

				m := models.TypeSliceInt64{
					SliceInt64:               s,
					PointerSliceInt64:        &s,
					SlicePointerInt64:        sPtr,
					PointerSlicePointerInt64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceInt64`", func() {
				var el int64 = 1
				s := []int64{el}
				sPtr := []*int64{&el}

				m := models.TypeSliceInt64{
					// SliceInt64:               s,
					PointerSliceInt64:        &s,
					SlicePointerInt64:        sPtr,
					PointerSlicePointerInt64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceInt64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceInt64`", func() {
				var el int64 = 1
				s := []int64{el}
				sPtr := []*int64{&el}

				m := models.TypeSliceInt64{
					SliceInt64: s,
					// PointerSliceInt64:        &s,
					SlicePointerInt64:        sPtr,
					PointerSlicePointerInt64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceInt64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerInt64`", func() {
				var el int64 = 1
				s := []int64{el}
				sPtr := []*int64{&el}

				m := models.TypeSliceInt64{
					SliceInt64:        s,
					PointerSliceInt64: &s,
					// SlicePointerInt64:        sPtr,
					PointerSlicePointerInt64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerInt64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerInt64`", func() {
				var el int64 = 1
				s := []int64{el}
				sPtr := []*int64{&el}

				m := models.TypeSliceInt64{
					SliceInt64:        s,
					PointerSliceInt64: &s,
					SlicePointerInt64: sPtr,
					// PointerSlicePointerInt64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerInt64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Float32", func() {
			It("should empty validation", func() {
				var el float32 = 1
				s := []float32{el}
				sPtr := []*float32{&el}

				m := models.TypeSliceFloat32{
					SliceFloat32:               s,
					PointerSliceFloat32:        &s,
					SlicePointerFloat32:        sPtr,
					PointerSlicePointerFloat32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceFloat32`", func() {
				var el float32 = 1
				s := []float32{el}
				sPtr := []*float32{&el}

				m := models.TypeSliceFloat32{
					// SliceFloat32:               s,
					PointerSliceFloat32:        &s,
					SlicePointerFloat32:        sPtr,
					PointerSlicePointerFloat32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceFloat32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceFloat32`", func() {
				var el float32 = 1
				s := []float32{el}
				sPtr := []*float32{&el}

				m := models.TypeSliceFloat32{
					SliceFloat32: s,
					// PointerSliceFloat32:        &s,
					SlicePointerFloat32:        sPtr,
					PointerSlicePointerFloat32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceFloat32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerFloat32`", func() {
				var el float32 = 1
				s := []float32{el}
				sPtr := []*float32{&el}

				m := models.TypeSliceFloat32{
					SliceFloat32:        s,
					PointerSliceFloat32: &s,
					// SlicePointerFloat32:        sPtr,
					PointerSlicePointerFloat32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerFloat32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerFloat32`", func() {
				var el float32 = 1
				s := []float32{el}
				sPtr := []*float32{&el}

				m := models.TypeSliceFloat32{
					SliceFloat32:        s,
					PointerSliceFloat32: &s,
					SlicePointerFloat32: sPtr,
					// PointerSlicePointerFloat32: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerFloat32"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Float64", func() {
			It("should empty validation", func() {
				var el float64 = 1
				s := []float64{el}
				sPtr := []*float64{&el}

				m := models.TypeSliceFloat64{
					SliceFloat64:               s,
					PointerSliceFloat64:        &s,
					SlicePointerFloat64:        sPtr,
					PointerSlicePointerFloat64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceFloat64`", func() {
				var el float64 = 1
				s := []float64{el}
				sPtr := []*float64{&el}

				m := models.TypeSliceFloat64{
					// SliceFloat64:               s,
					PointerSliceFloat64:        &s,
					SlicePointerFloat64:        sPtr,
					PointerSlicePointerFloat64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceFloat64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceFloat64`", func() {
				var el float64 = 1
				s := []float64{el}
				sPtr := []*float64{&el}

				m := models.TypeSliceFloat64{
					SliceFloat64: s,
					// PointerSliceFloat64:        &s,
					SlicePointerFloat64:        sPtr,
					PointerSlicePointerFloat64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceFloat64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerFloat64`", func() {
				var el float64 = 1
				s := []float64{el}
				sPtr := []*float64{&el}

				m := models.TypeSliceFloat64{
					SliceFloat64:        s,
					PointerSliceFloat64: &s,
					// SlicePointerFloat64:        sPtr,
					PointerSlicePointerFloat64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerFloat64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerFloat64`", func() {
				var el float64 = 1
				s := []float64{el}
				sPtr := []*float64{&el}

				m := models.TypeSliceFloat64{
					SliceFloat64:        s,
					PointerSliceFloat64: &s,
					SlicePointerFloat64: sPtr,
					// PointerSlicePointerFloat64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerFloat64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Complex64", func() {
			It("should empty validation", func() {
				var el complex64 = 1
				s := []complex64{el}
				sPtr := []*complex64{&el}

				m := models.TypeSliceComplex64{
					SliceComplex64:               s,
					PointerSliceComplex64:        &s,
					SlicePointerComplex64:        sPtr,
					PointerSlicePointerComplex64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceComplex64`", func() {
				var el complex64 = 1
				s := []complex64{el}
				sPtr := []*complex64{&el}

				m := models.TypeSliceComplex64{
					// SliceComplex64:               s,
					PointerSliceComplex64:        &s,
					SlicePointerComplex64:        sPtr,
					PointerSlicePointerComplex64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceComplex64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceComplex64`", func() {
				var el complex64 = 1
				s := []complex64{el}
				sPtr := []*complex64{&el}

				m := models.TypeSliceComplex64{
					SliceComplex64: s,
					// PointerSliceComplex64:        &s,
					SlicePointerComplex64:        sPtr,
					PointerSlicePointerComplex64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceComplex64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerComplex64`", func() {
				var el complex64 = 1
				s := []complex64{el}
				sPtr := []*complex64{&el}

				m := models.TypeSliceComplex64{
					SliceComplex64:        s,
					PointerSliceComplex64: &s,
					// SlicePointerComplex64:        sPtr,
					PointerSlicePointerComplex64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerComplex64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerComplex64`", func() {
				var el complex64 = 1
				s := []complex64{el}
				sPtr := []*complex64{&el}

				m := models.TypeSliceComplex64{
					SliceComplex64:        s,
					PointerSliceComplex64: &s,
					SlicePointerComplex64: sPtr,
					// PointerSlicePointerComplex64: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerComplex64"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("Complex128", func() {
			It("should empty validation", func() {
				var el complex128 = 1
				s := []complex128{el}
				sPtr := []*complex128{&el}

				m := models.TypeSliceComplex128{
					SliceComplex128:               s,
					PointerSliceComplex128:        &s,
					SlicePointerComplex128:        sPtr,
					PointerSlicePointerComplex128: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should check if field `SliceComplex128`", func() {
				var el complex128 = 1
				s := []complex128{el}
				sPtr := []*complex128{&el}

				m := models.TypeSliceComplex128{
					// SliceComplex128:               s,
					PointerSliceComplex128:        &s,
					SlicePointerComplex128:        sPtr,
					PointerSlicePointerComplex128: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SliceComplex128"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSliceComplex128`", func() {
				var el complex128 = 1
				s := []complex128{el}
				sPtr := []*complex128{&el}

				m := models.TypeSliceComplex128{
					SliceComplex128: s,
					// PointerSliceComplex128:        &s,
					SlicePointerComplex128:        sPtr,
					PointerSlicePointerComplex128: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSliceComplex128"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `SlicePointerComplex128`", func() {
				var el complex128 = 1
				s := []complex128{el}
				sPtr := []*complex128{&el}

				m := models.TypeSliceComplex128{
					SliceComplex128:        s,
					PointerSliceComplex128: &s,
					// SlicePointerComplex128:        sPtr,
					PointerSlicePointerComplex128: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("SlicePointerComplex128"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should check if field `PointerSlicePointerComplex128`", func() {
				var el complex128 = 1
				s := []complex128{el}
				sPtr := []*complex128{&el}

				m := models.TypeSliceComplex128{
					SliceComplex128:        s,
					PointerSliceComplex128: &s,
					SlicePointerComplex128: sPtr,
					// PointerSlicePointerComplex128: &sPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerSlicePointerComplex128"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

	})
})
