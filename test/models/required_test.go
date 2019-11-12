package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Required", func() {
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

			m := models.RequiredChanString{
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

			m := models.RequiredChanString{
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

			m := models.RequiredChanString{
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

			m := models.RequiredChanString{
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

			m := models.RequiredChanString{
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

			m := models.RequiredChanUint{
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

			m := models.RequiredChanUint{
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

			m := models.RequiredChanUint{
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

			m := models.RequiredChanUint{
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

			m := models.RequiredChanUint{
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

	When("String", func() {
		It("should empty validation", func() {
			s := "Chico Bento!"
			m := models.RequiredString{
				String:  s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should check if field `String`", func() {
			s := "Chico Bento!"
			m := models.RequiredString{
				// String: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("String"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should check if field `String Pointer`", func() {
			s := "Chico Bento!"
			m := models.RequiredString{
				String: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("required"))
			Expect(errs[0].(error).Error()).To(Equal("Error: Validation of field 'Pointer' failed on tag 'required'"))
		})
	})

	When("Interface", func() {
		It("should return empty validation (struct)", func() {
			var s interface{}
			s = models.MapTypeA{
				Int:  123,
				Bool: true,
			}

			m := models.RequiredInterface{
				Interface:        s,
				PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should return empty validation", func() {
			var s interface{}
			s = "Chico Bento!"

			m := models.RequiredInterface{
				Interface:        s,
				PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail to validate without field `Interface`", func() {
			var s interface{}
			s = false
			m := models.RequiredInterface{
				// Interface: s,
				PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Interface"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})

		It("should fail to validate without field `PointerInterface`", func() {
			var s interface{}
			s = 222
			m := models.RequiredInterface{
				Interface: s,
				// PointerInterface: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("PointerInterface"))
			Expect(errs[0].Tag()).To(Equal("required"))
		})
	})

	When("Struct", func() {
		It("should fail to validate field `Struct`", func() {
			structA := models.RequiredStructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.RequiredStructA{structA}
			sliceStructAPtr := []*models.RequiredStructA{&structA}
			t := models.RequiredStruct{
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
			structA := models.RequiredStructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.RequiredStructA{structA}
			sliceStructAPtr := []*models.RequiredStructA{&structA}
			t := models.RequiredStruct{
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
			structA := models.RequiredStructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.RequiredStructA{structA}
			sliceStructAPtr := []*models.RequiredStructA{&structA}
			t := models.RequiredStruct{
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
			structA := models.RequiredStructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.RequiredStructA{structA}
			sliceStructAPtr := []*models.RequiredStructA{&structA}
			t := models.RequiredStruct{
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
			structA := models.RequiredStructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.RequiredStructA{structA}
			sliceStructAPtr := []*models.RequiredStructA{&structA}
			t := models.RequiredStruct{
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
			structA := models.RequiredStructA{
				A: "Chico Bento!",
				B: 22,
			}

			sliceStructA := []models.RequiredStructA{structA}
			sliceStructAPtr := []*models.RequiredStructA{&structA}
			t := models.RequiredStruct{
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

	When("Slice", func() {
		When("String", func() {
			It("should empty validation", func() {
				el := "Chico Bento!"
				s := []string{el}
				sPtr := []*string{&el}

				m := models.RequiredSliceString{
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

				m := models.RequiredSliceString{
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

				m := models.RequiredSliceString{
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

				m := models.RequiredSliceString{
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

				m := models.RequiredSliceString{
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

				m := models.RequiredSliceUint{
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

				m := models.RequiredSliceUint{
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

				m := models.RequiredSliceUint{
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

				m := models.RequiredSliceUint{
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

				m := models.RequiredSliceUint{
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

				m := models.RequiredSliceUint8{
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

				m := models.RequiredSliceUint8{
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

				m := models.RequiredSliceUint8{
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

				m := models.RequiredSliceUint8{
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

				m := models.RequiredSliceUint8{
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

				m := models.RequiredSliceUint16{
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

				m := models.RequiredSliceUint16{
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

				m := models.RequiredSliceUint16{
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

				m := models.RequiredSliceUint16{
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

				m := models.RequiredSliceUint16{
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

				m := models.RequiredSliceUint32{
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

				m := models.RequiredSliceUint32{
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

				m := models.RequiredSliceUint32{
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

				m := models.RequiredSliceUint32{
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

				m := models.RequiredSliceUint32{
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

				m := models.RequiredSliceUint64{
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

				m := models.RequiredSliceUint64{
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

				m := models.RequiredSliceUint64{
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

				m := models.RequiredSliceUint64{
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

				m := models.RequiredSliceUint64{
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

				m := models.RequiredSliceUintptr{
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

				m := models.RequiredSliceUintptr{
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

				m := models.RequiredSliceUintptr{
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

				m := models.RequiredSliceUintptr{
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

				m := models.RequiredSliceUintptr{
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

				m := models.RequiredSliceInt{
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

				m := models.RequiredSliceInt{
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

				m := models.RequiredSliceInt{
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

				m := models.RequiredSliceInt{
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

				m := models.RequiredSliceInt{
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

				m := models.RequiredSliceInt8{
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

				m := models.RequiredSliceInt8{
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

				m := models.RequiredSliceInt8{
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

				m := models.RequiredSliceInt8{
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

				m := models.RequiredSliceInt8{
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

				m := models.RequiredSliceInt16{
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

				m := models.RequiredSliceInt16{
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

				m := models.RequiredSliceInt16{
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

				m := models.RequiredSliceInt16{
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

				m := models.RequiredSliceInt16{
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

				m := models.RequiredSliceInt32{
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

				m := models.RequiredSliceInt32{
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

				m := models.RequiredSliceInt32{
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

				m := models.RequiredSliceInt32{
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

				m := models.RequiredSliceInt32{
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

				m := models.RequiredSliceInt64{
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

				m := models.RequiredSliceInt64{
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

				m := models.RequiredSliceInt64{
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

				m := models.RequiredSliceInt64{
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

				m := models.RequiredSliceInt64{
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

				m := models.RequiredSliceFloat32{
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

				m := models.RequiredSliceFloat32{
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

				m := models.RequiredSliceFloat32{
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

				m := models.RequiredSliceFloat32{
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

				m := models.RequiredSliceFloat32{
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

				m := models.RequiredSliceFloat64{
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

				m := models.RequiredSliceFloat64{
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

				m := models.RequiredSliceFloat64{
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

				m := models.RequiredSliceFloat64{
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

				m := models.RequiredSliceFloat64{
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

				m := models.RequiredSliceComplex64{
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

				m := models.RequiredSliceComplex64{
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

				m := models.RequiredSliceComplex64{
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

				m := models.RequiredSliceComplex64{
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

				m := models.RequiredSliceComplex64{
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

				m := models.RequiredSliceComplex128{
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

				m := models.RequiredSliceComplex128{
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

				m := models.RequiredSliceComplex128{
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

				m := models.RequiredSliceComplex128{
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

				m := models.RequiredSliceComplex128{
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

	When("Numbers", func() {
		It("should fail to validate all numbers", func() {
			m := models.RequiredNumber{}

			errs := m.Validate()
			Expect(errs).To(HaveLen(30))
		})
	})

	When("Map", func() {

		When("MapStringToInterface", func() {
			It("should empty validation", func() {
				d := map[string]interface{}{
					"name": "Chico Bento",
				}

				m := models.RequiredMapStringToInterface{
					MapStringToInterface:        d,
					PointerMapStringToInterface: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.RequiredMapStringToInterface{}

				errs := m.Validate()
				Expect(errs).To(HaveLen(2))
				Expect(errs[0].Field()).To(Equal("MapStringToInterface"))
				Expect(errs[0].Tag()).To(Equal("required"))
				Expect(errs[1].Field()).To(Equal("PointerMapStringToInterface"))
				Expect(errs[1].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `MapStringToInterface`", func() {
				d := map[string]interface{}{
					"name": "Chico Bento",
				}

				m := models.RequiredMapStringToInterface{
					// MapStringToInterface:        d,
					PointerMapStringToInterface: &d,
				}
				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("MapStringToInterface"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `PointerMapStringToInterface`", func() {
				d := map[string]interface{}{
					"name": "Chico Bento",
				}

				m := models.RequiredMapStringToInterface{
					MapStringToInterface: d,
					// PointerMapStringToInterface: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerMapStringToInterface"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("MapStringToString", func() {
			It("should empty validation", func() {
				d := map[string]string{
					"name": "Chico Bento",
				}

				m := models.RequiredMapStringToString{
					MapStringToString:        d,
					PointerMapStringToString: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.RequiredMapStringToString{}

				errs := m.Validate()
				Expect(errs).To(HaveLen(2))
				Expect(errs[0].Field()).To(Equal("MapStringToString"))
				Expect(errs[0].Tag()).To(Equal("required"))
				Expect(errs[1].Field()).To(Equal("PointerMapStringToString"))
				Expect(errs[1].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `MapStringToString`", func() {
				d := map[string]string{
					"name": "Chico Bento",
				}

				m := models.RequiredMapStringToString{
					// MapStringToString:        d,
					PointerMapStringToString: &d,
				}
				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("MapStringToString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `PointerMapStringToString`", func() {
				d := map[string]string{
					"name": "Chico Bento",
				}

				m := models.RequiredMapStringToString{
					MapStringToString: d,
					// PointerMapStringToString: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerMapStringToString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("MapIntToString", func() {
			It("should empty validation", func() {
				d := map[int]string{
					0: "Chico Bento",
				}

				m := models.RequiredMapIntToString{
					MapIntToString:        d,
					PointerMapIntToString: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.RequiredMapIntToString{}

				errs := m.Validate()
				Expect(errs).To(HaveLen(2))
				Expect(errs[0].Field()).To(Equal("MapIntToString"))
				Expect(errs[0].Tag()).To(Equal("required"))
				Expect(errs[1].Field()).To(Equal("PointerMapIntToString"))
				Expect(errs[1].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `MapIntToString`", func() {
				d := map[int]string{
					0: "Chico Bento",
				}

				m := models.RequiredMapIntToString{
					// MapIntToString:        d,
					PointerMapIntToString: &d,
				}
				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("MapIntToString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `PointerMapIntToString`", func() {
				d := map[int]string{
					0: "Chico Bento",
				}

				m := models.RequiredMapIntToString{
					MapIntToString: d,
					// PointerMapIntToString: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerMapIntToString"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("MapIntToBool", func() {
			It("should empty validation", func() {
				d := map[int]bool{
					0: false,
				}

				m := models.RequiredMapIntToBool{
					MapIntToBool:        d,
					PointerMapIntToBool: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.RequiredMapIntToBool{}

				errs := m.Validate()
				Expect(errs).To(HaveLen(2))
				Expect(errs[0].Field()).To(Equal("MapIntToBool"))
				Expect(errs[0].Tag()).To(Equal("required"))
				Expect(errs[1].Field()).To(Equal("PointerMapIntToBool"))
				Expect(errs[1].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `MapIntToBool`", func() {
				d := map[int]bool{
					0: true,
				}

				m := models.RequiredMapIntToBool{
					// MapIntToBool:        d,
					PointerMapIntToBool: &d,
				}
				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("MapIntToBool"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `PointerMapIntToBool`", func() {
				d := map[int]bool{
					0: false,
				}

				m := models.RequiredMapIntToBool{
					MapIntToBool: d,
					// PointerMapIntToBool: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerMapIntToBool"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})

		When("MapIntToStruct", func() {
			It("should empty validation", func() {
				dataA := models.MapTypeA{
					Int:  22,
					Bool: false,
				}

				dataB := models.MapTypeB{
					Bool:   true,
					String: "Chico Bento",
					Float:  1.0,
				}

				a := map[int]models.MapTypeA{
					0: dataA,
				}

				aPtr := map[int]*models.MapTypeA{
					0: &dataA,
				}

				b := map[int]models.MapTypeB{
					0: dataB,
				}

				bPtr := map[int]*models.MapTypeB{
					0: &dataB,
				}

				m := models.RequiredMapIntToStruct{
					MapIntToStructA:               a,
					PointerMapIntToStructA:        aPtr,
					MapIntToStructB:               &b,
					PointerMapIntToStructPointerB: &bPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.RequiredMapIntToStruct{}

				errs := m.Validate()
				Expect(errs).To(HaveLen(4))
				Expect(errs[0].Field()).To(Equal("MapIntToStructA"))
				Expect(errs[0].Tag()).To(Equal("required"))
				Expect(errs[1].Field()).To(Equal("PointerMapIntToStructA"))
				Expect(errs[1].Tag()).To(Equal("required"))
				Expect(errs[2].Field()).To(Equal("MapIntToStructB"))
				Expect(errs[2].Tag()).To(Equal("required"))
				Expect(errs[3].Field()).To(Equal("PointerMapIntToStructPointerB"))
				Expect(errs[3].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `MapIntToStructA`", func() {
				dataA := models.MapTypeA{
					Int:  22,
					Bool: false,
				}

				dataB := models.MapTypeB{
					Bool:   true,
					String: "Chico Bento",
					Float:  1.0,
				}

				aPtr := map[int]*models.MapTypeA{
					0: &dataA,
				}

				b := map[int]models.MapTypeB{
					0: dataB,
				}

				bPtr := map[int]*models.MapTypeB{
					0: &dataB,
				}

				m := models.RequiredMapIntToStruct{
					// MapIntToStructA:               a,
					PointerMapIntToStructA:        aPtr,
					MapIntToStructB:               &b,
					PointerMapIntToStructPointerB: &bPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("MapIntToStructA"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `PointerMapIntToStructA`", func() {
				dataA := models.MapTypeA{
					Int:  22,
					Bool: false,
				}
				dataB := models.MapTypeB{
					Bool:   true,
					String: "Chico Bento",
					Float:  1.0,
				}

				a := map[int]models.MapTypeA{
					0: dataA,
				}

				b := map[int]models.MapTypeB{
					0: dataB,
				}

				bPtr := map[int]*models.MapTypeB{
					0: &dataB,
				}

				m := models.RequiredMapIntToStruct{
					MapIntToStructA: a,
					// PointerMapIntToStructA:        aPtr,
					MapIntToStructB:               &b,
					PointerMapIntToStructPointerB: &bPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerMapIntToStructA"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `MapIntToStructB`", func() {
				dataA := models.MapTypeA{
					Int:  22,
					Bool: false,
				}
				dataB := models.MapTypeB{
					Bool:   true,
					String: "Chico Bento",
					Float:  1.0,
				}

				a := map[int]models.MapTypeA{
					0: dataA,
				}

				aPtr := map[int]*models.MapTypeA{
					0: &dataA,
				}

				bPtr := map[int]*models.MapTypeB{
					0: &dataB,
				}

				m := models.RequiredMapIntToStruct{
					MapIntToStructA:        a,
					PointerMapIntToStructA: aPtr,
					// MapIntToStructB:               &b,
					PointerMapIntToStructPointerB: &bPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("MapIntToStructB"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})

			It("should fail to validate without field `PointerMapIntToStructPointerB`", func() {
				dataA := models.MapTypeA{
					Int:  22,
					Bool: false,
				}
				dataB := models.MapTypeB{
					Bool:   true,
					String: "Chico Bento",
					Float:  1.0,
				}

				a := map[int]models.MapTypeA{
					0: dataA,
				}

				aPtr := map[int]*models.MapTypeA{
					0: &dataA,
				}

				b := map[int]models.MapTypeB{
					0: dataB,
				}

				m := models.RequiredMapIntToStruct{
					MapIntToStructA:        a,
					PointerMapIntToStructA: aPtr,
					MapIntToStructB:        &b,
					// PointerMapIntToStructPointerB: &bPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(1))
				Expect(errs[0].Field()).To(Equal("PointerMapIntToStructPointerB"))
				Expect(errs[0].Tag()).To(Equal("required"))
			})
		})
	})
})
