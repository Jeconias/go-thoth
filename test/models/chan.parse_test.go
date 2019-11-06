package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chan", func() {
	Describe("Required", func() {
		When("Channel of String", func() {
			It("should return empty validation", func() {
				d := "Chico Bento"

				s := make(chan string, 1)
				s <- d

				sPtr := make(chan *string, 1)
				sPtr <- &d

				sSlice := make(chan []string, 1)
				sSlice <- []string{d}

				sSlicePtr := make(chan []*string, 1)
				sSlicePtr <- []*string{&d}

				m := models.TypeChanString{
					ChanString:             s,
					PointerChanString:      sPtr,
					ChanSliceString:        sSlice,
					ChanSlicePointerString: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))

				close(s)
				close(sPtr)
				close(sSlice)
				close(sSlicePtr)
			})

			It("should fail to validate field `Chan`", func() {
				d := "Chico Bento!"
				sPtr := make(chan *string, 1)
				sPtr <- &d

				sSlice := make(chan []string, 1)
				sSlice <- []string{d}

				sSlicePtr := make(chan []*string, 1)
				sSlicePtr <- []*string{&d}

				m := models.TypeChanString{
					// ChanString: s,
					PointerChanString:      sPtr,
					ChanSliceString:        sSlice,
					ChanSlicePointerString: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("ChanString"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(sPtr)
				close(sSlice)
				close(sSlicePtr)
			})

			It("should fail to validate field `PointerChanString`", func() {
				d := "Chico Bento!"
				s := make(chan string, 1)
				s <- d

				sSlice := make(chan []string, 1)
				sSlice <- []string{d}

				sSlicePtr := make(chan []*string, 1)
				sSlicePtr <- []*string{&d}

				m := models.TypeChanString{
					ChanString: s,
					// PointerChanString: &s,
					ChanSliceString:        sSlice,
					ChanSlicePointerString: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerChanString"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(s)
				close(sSlice)
				close(sSlicePtr)
			})

			It("should fail to validate field `ChanSliceString`", func() {
				d := "Chico Bento!"
				s := make(chan string, 1)
				s <- d

				sPtr := make(chan *string, 1)
				sPtr <- &d

				sSlicePtr := make(chan []*string, 1)
				sSlicePtr <- []*string{&d}

				m := models.TypeChanString{
					ChanString:             s,
					PointerChanString:      sPtr,
					ChanSlicePointerString: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("ChanSliceString"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(s)
				close(sPtr)
				close(sSlicePtr)
			})

			It("should fail to validate field `ChanSlicePointerString`", func() {
				d := "Chico Bento!"
				s := make(chan string, 1)
				s <- d

				sPtr := make(chan *string, 1)
				sPtr <- &d

				sSlice := make(chan []string, 1)
				sSlice <- []string{d}

				m := models.TypeChanString{
					ChanString:        s,
					PointerChanString: sPtr,
					ChanSliceString:   sSlice,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("ChanSlicePointerString"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(s)
				close(sPtr)
				close(sSlice)
			})
		})

		When("Channel of uint", func() {
			It("should return empty validation", func() {
				var d uint = 1

				s := make(chan uint, 1)
				s <- d

				sPtr := make(chan *uint, 1)
				sPtr <- &d

				sSlice := make(chan []uint, 1)
				sSlice <- []uint{d}

				sSlicePtr := make(chan []*uint, 1)
				sSlicePtr <- []*uint{&d}

				m := models.TypeChanUint{
					ChanUint:             s,
					PointerChanUint:      sPtr,
					ChanSliceUint:        sSlice,
					ChanSlicePointerUint: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))

				close(s)
				close(sPtr)
				close(sSlice)
				close(sSlicePtr)
			})

			It("should fail to validate field `Chan`", func() {
				var d uint = 1

				sPtr := make(chan *uint, 1)
				sPtr <- &d

				sSlice := make(chan []uint, 1)
				sSlice <- []uint{d}

				sSlicePtr := make(chan []*uint, 1)
				sSlicePtr <- []*uint{&d}

				m := models.TypeChanUint{
					// ChanUint: s,
					PointerChanUint:      sPtr,
					ChanSliceUint:        sSlice,
					ChanSlicePointerUint: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("ChanUint"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(sPtr)
				close(sSlice)
				close(sSlicePtr)
			})

			It("should fail to validate field `PointerChanUint`", func() {
				var d uint = 1
				s := make(chan uint, 1)
				s <- d

				sSlice := make(chan []uint, 1)
				sSlice <- []uint{d}

				sSlicePtr := make(chan []*uint, 1)
				sSlicePtr <- []*uint{&d}

				m := models.TypeChanUint{
					ChanUint: s,
					// PointerChanUint: &s,
					ChanSliceUint:        sSlice,
					ChanSlicePointerUint: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerChanUint"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(s)
				close(sSlice)
				close(sSlicePtr)
			})

			It("should fail to validate field `ChanSliceUint`", func() {
				var d uint = 1
				s := make(chan uint, 1)
				s <- d

				sPtr := make(chan *uint, 1)
				sPtr <- &d

				sSlicePtr := make(chan []*uint, 1)
				sSlicePtr <- []*uint{&d}

				m := models.TypeChanUint{
					ChanUint:             s,
					PointerChanUint:      sPtr,
					ChanSlicePointerUint: sSlicePtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("ChanSliceUint"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(s)
				close(sPtr)
				close(sSlicePtr)
			})

			It("should fail to validate field `ChanSlicePointerUint`", func() {
				var d uint = 1
				s := make(chan uint, 1)
				s <- d

				sPtr := make(chan *uint, 1)
				sPtr <- &d

				sSlice := make(chan []uint, 1)
				sSlice <- []uint{d}

				m := models.TypeChanUint{
					ChanUint:        s,
					PointerChanUint: sPtr,
					ChanSliceUint:   sSlice,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("ChanSlicePointerUint"))
				Expect(errs[0].Tag()).To(Equal("required"))

				close(s)
				close(sPtr)
				close(sSlice)
			})
		})
	})
})
