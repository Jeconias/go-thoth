package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Len", func() {

	When("string", func() {
		It("should empty validation", func() {
			s := "chico"
			m := models.LenStringValidate{
				String:  s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should  fail `len` validation when text is longer than five characters", func() {
			s := "Chico Bento"
			m := models.LenStringValidate{
				String:  "chico",
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field `String`", func() {
			s := "chico"
			m := models.LenStringValidate{
				// String:        s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("String"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail to validater without field `Pointer`", func() {
			s := "chico"
			m := models.LenStringValidate{
				String: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})
	})

	When("chan", func() {
		It("should empty validation", func() {
			s1 := "chico"
			s2 := "bento"
			var u1 uint = 1
			var u2 uint = 2

			chanSource := make(chan string, 2)
			chanSource <- s1
			chanSource <- s2

			chanSourcePtr := make(chan *string, 2)
			chanSourcePtr <- &s1
			chanSourcePtr <- &s2

			chanSourceUint := make(chan uint, 2)
			chanSourceUint <- u1
			chanSourceUint <- u2

			chanSourceUintPtr := make(chan *uint, 2)
			chanSourceUintPtr <- &u1
			chanSourceUintPtr <- &u2

			m := models.LenChanValidate{
				ChanString:        chanSource,
				ChanStringPointer: chanSourcePtr,
				ChanUint:          chanSourceUint,
				ChanUintPointer:   chanSourceUintPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))

			close(chanSource)
			close(chanSourcePtr)
			close(chanSourceUint)
			close(chanSourceUintPtr)
		})

		It("should  fail `len` validation when text is longer than five characters", func() {
			s1 := "chico"
			s2 := "bento"
			s3 := "junior"
			var u1 uint = 1
			var u2 uint = 2

			chanSource := make(chan string, 2)
			chanSource <- s1
			chanSource <- s2

			chanSourcePtr := make(chan *string, 3)
			chanSourcePtr <- &s1
			chanSourcePtr <- &s2
			chanSourcePtr <- &s3

			chanSourceUint := make(chan uint, 2)
			chanSourceUint <- u1
			chanSourceUint <- u2

			chanSourceUintPtr := make(chan *uint, 2)
			chanSourceUintPtr <- &u1
			chanSourceUintPtr <- &u2

			m := models.LenChanValidate{
				ChanString:        chanSource,
				ChanStringPointer: chanSourcePtr,
				ChanUint:          chanSourceUint,
				ChanUintPointer:   chanSourceUintPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ChanStringPointer"))
			Expect(errs[0].Tag()).To(Equal("len"))

			close(chanSource)
			close(chanSourcePtr)
			close(chanSourceUint)
			close(chanSourceUintPtr)
		})

		It("should fail `len` validation without field `ChanString`", func() {
			s1 := "chico"
			s2 := "bento"
			var u1 uint = 1
			var u2 uint = 2

			chanSource := make(chan string, 2)
			chanSource <- s1
			chanSource <- s2

			chanSourcePtr := make(chan *string, 2)
			s1 = "Chico Bento Junior"
			chanSourcePtr <- &s1
			chanSourcePtr <- &s2

			chanSourceUint := make(chan uint, 2)
			chanSourceUint <- u1
			chanSourceUint <- u2

			chanSourceUintPtr := make(chan *uint, 2)
			chanSourceUintPtr <- &u1
			chanSourceUintPtr <- &u2

			m := models.LenChanValidate{
				// ChanString:        chanSource,
				ChanStringPointer: chanSourcePtr,
				ChanUint:          chanSourceUint,
				ChanUintPointer:   chanSourceUintPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ChanString"))
			Expect(errs[0].Tag()).To(Equal("len"))

			close(chanSource)
			close(chanSourcePtr)
			close(chanSourceUint)
			close(chanSourceUintPtr)
		})

		It("should fail `len` validation without field `ChanStringPointer`", func() {
			s1 := "chico"
			s2 := "bento"
			var u1 uint = 1
			var u2 uint = 2

			chanSource := make(chan string, 2)
			chanSource <- s1
			chanSource <- s2

			chanSourcePtr := make(chan *string, 2)
			chanSourcePtr <- &s1
			chanSourcePtr <- &s2

			chanSourceUint := make(chan uint, 2)
			chanSourceUint <- u1
			chanSourceUint <- u2

			chanSourceUintPtr := make(chan *uint, 2)
			chanSourceUintPtr <- &u1
			chanSourceUintPtr <- &u2

			m := models.LenChanValidate{
				ChanString: chanSource,
				// ChanStringPointer: chanSourcePtr,
				ChanUint:        chanSourceUint,
				ChanUintPointer: chanSourceUintPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ChanStringPointer"))
			Expect(errs[0].Tag()).To(Equal("len"))

			close(chanSource)
			close(chanSourcePtr)
			close(chanSourceUint)
			close(chanSourceUintPtr)
		})

		It("should fail `len` validation without field `ChanUint`", func() {
			s1 := "chico"
			s2 := "bento"
			var u1 uint = 1
			var u2 uint = 2

			chanSource := make(chan string, 2)
			chanSource <- s1
			chanSource <- s2

			chanSourcePtr := make(chan *string, 2)
			chanSourcePtr <- &s1
			chanSourcePtr <- &s2

			chanSourceUint := make(chan uint, 2)
			chanSourceUint <- u1
			chanSourceUint <- u2

			chanSourceUintPtr := make(chan *uint, 2)
			chanSourceUintPtr <- &u1
			chanSourceUintPtr <- &u2

			m := models.LenChanValidate{
				ChanString:        chanSource,
				ChanStringPointer: chanSourcePtr,
				// ChanUint:        chanSourceUint,
				ChanUintPointer: chanSourceUintPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ChanUint"))
			Expect(errs[0].Tag()).To(Equal("len"))

			close(chanSource)
			close(chanSourcePtr)
			close(chanSourceUint)
			close(chanSourceUintPtr)
		})

		It("should fail `len` validation without field `ChanUintPointer`", func() {
			s1 := "chico"
			s2 := "bento"
			var u1 uint = 1
			var u2 uint = 2

			chanSource := make(chan string, 2)
			chanSource <- s1
			chanSource <- s2

			chanSourcePtr := make(chan *string, 2)
			chanSourcePtr <- &s1
			chanSourcePtr <- &s2

			chanSourceUint := make(chan uint, 2)
			chanSourceUint <- u1
			chanSourceUint <- u2

			chanSourceUintPtr := make(chan *uint, 2)
			chanSourceUintPtr <- &u1
			chanSourceUintPtr <- &u2

			m := models.LenChanValidate{
				ChanString:        chanSource,
				ChanStringPointer: chanSourcePtr,
				ChanUint:          chanSourceUint,
				// ChanUintPointer: chanSourceUintPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ChanUintPointer"))
			Expect(errs[0].Tag()).To(Equal("len"))

			close(chanSource)
			close(chanSourcePtr)
			close(chanSourceUint)
			close(chanSourceUintPtr)
		})
	})

	When("slice string", func() {
		It("should `len` validation successfully", func() {
			s1 := "chico"
			s2 := "bento"
			s3 := "jose"
			s4 := "snake"
			s5 := "sallet"

			s := []string{s1, s2, s3, s4, s5}
			sPtr := []*string{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceStringValidate{
				String:              s,
				SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail `len` validation without field `String`", func() {
			s1 := "chico"
			s2 := "bento"
			s3 := "jose"
			s4 := "snake"
			s5 := "sallet"

			s := []string{s1, s2, s3, s4, s5}
			sPtr := []*string{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceStringValidate{
				// String:              s,
				SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("String"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field `SlicePointer`", func() {
			s1 := "chico"
			s2 := "bento"
			s3 := "jose"
			s4 := "snake"
			s5 := "sallet"

			s := []string{s1, s2, s3, s4, s5}
			sPtr := []*string{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceStringValidate{
				String: s,
				// SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("SlicePointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field `PointerSlice`", func() {
			s1 := "chico"
			s2 := "bento"
			s3 := "jose"
			s4 := "snake"
			s5 := "sallet"

			s := []string{s1, s2, s3, s4, s5}
			sPtr := []*string{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceStringValidate{
				String:       s,
				SlicePointer: sPtr,
				// PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerSlice"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field `PointerSlicePointer`", func() {
			s1 := "chico"
			s2 := "bento"
			s3 := "jose"
			s4 := "snake"
			s5 := "sallet"

			s := []string{s1, s2, s3, s4, s5}
			sPtr := []*string{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceStringValidate{
				String:       s,
				SlicePointer: sPtr,
				PointerSlice: &s,
				// PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerSlicePointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})
	})

	When("slice int", func() {
		It("should `len` validation successfully", func() {
			s1 := 1
			s2 := 2
			s3 := 3
			s4 := 4
			s5 := 5

			s := []int{s1, s2, s3, s4, s5}
			sPtr := []*int{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceIntValidate{
				Int:                 s,
				SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail `len` validation without field int", func() {
			s1 := 1
			s2 := 2
			s3 := 3
			s4 := 4
			s5 := 5

			s := []int{s1, s2, s3, s4, s5}
			sPtr := []*int{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceIntValidate{
				// Int:              s,
				SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Int"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field SlicePointer", func() {
			s1 := 1
			s2 := 2
			s3 := 3
			s4 := 4
			s5 := 5

			s := []int{s1, s2, s3, s4, s5}
			sPtr := []*int{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceIntValidate{
				Int: s,
				// SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("SlicePointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field PointerSlice", func() {
			s1 := 1
			s2 := 2
			s3 := 3
			s4 := 4
			s5 := 5

			s := []int{s1, s2, s3, s4, s5}
			sPtr := []*int{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceIntValidate{
				Int:          s,
				SlicePointer: sPtr,
				// PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerSlice"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field PointerSlicePointer", func() {
			s1 := 1
			s2 := 2
			s3 := 3
			s4 := 4
			s5 := 5

			s := []int{s1, s2, s3, s4, s5}
			sPtr := []*int{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceIntValidate{
				Int:          s,
				SlicePointer: sPtr,
				PointerSlice: &s,
				// PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerSlicePointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})
	})

	When("slice float64", func() {
		It("should `len` validation successfully", func() {
			var s1 float64 = 1
			var s2 float64 = 2
			var s3 float64 = 3
			var s4 float64 = 4
			var s5 float64 = 5

			s := []float64{s1, s2, s3, s4, s5}
			sPtr := []*float64{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceFloat64Validate{
				Float64:             s,
				SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail `len` validation without field int", func() {
			var s1 float64 = 1
			var s2 float64 = 2
			var s3 float64 = 3
			var s4 float64 = 4
			var s5 float64 = 5

			s := []float64{s1, s2, s3, s4, s5}
			sPtr := []*float64{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceFloat64Validate{
				// Float64:              s,
				SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Float64"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field SlicePointer", func() {
			var s1 float64 = 1
			var s2 float64 = 2
			var s3 float64 = 3
			var s4 float64 = 4
			var s5 float64 = 5

			s := []float64{s1, s2, s3, s4, s5}
			sPtr := []*float64{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceFloat64Validate{
				Float64: s,
				// SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("SlicePointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field PointerSlice", func() {
			var s1 float64 = 1
			var s2 float64 = 2
			var s3 float64 = 3
			var s4 float64 = 4
			var s5 float64 = 5

			s := []float64{s1, s2, s3, s4, s5}
			sPtr := []*float64{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceFloat64Validate{
				Float64:      s,
				SlicePointer: sPtr,
				// PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerSlice"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should fail `len` validation without field PointerSlicePointer", func() {
			var s1 float64 = 1
			var s2 float64 = 2
			var s3 float64 = 3
			var s4 float64 = 4
			var s5 float64 = 5

			s := []float64{s1, s2, s3, s4, s5}
			sPtr := []*float64{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceFloat64Validate{
				Float64:      s,
				SlicePointer: sPtr,
				PointerSlice: &s,
				// PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerSlicePointer"))
			Expect(errs[0].Tag()).To(Equal("len"))
		})

		It("should  fail `len` validation when text is longer than six elements", func() {
			var s1 float64 = 1
			var s2 float64 = 2
			var s3 float64 = 3
			var s4 float64 = 4
			var s5 float64 = 5
			var s6 float64 = 5

			s := []float64{s1, s2, s3, s4, s5, s6}
			sPtr := []*float64{&s1, &s2, &s3, &s4, &s5}

			m := models.LenSliceFloat64Validate{
				Float64:             s,
				SlicePointer:        sPtr,
				PointerSlice:        &s,
				PointerSlicePointer: &sPtr,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(2))
			Expect(errs[0].Field()).To(Equal("Float64"))
			Expect(errs[0].Tag()).To(Equal("len"))
			Expect(errs[1].Field()).To(Equal("PointerSlice"))
			Expect(errs[1].Tag()).To(Equal("len"))
		})
	})
})
