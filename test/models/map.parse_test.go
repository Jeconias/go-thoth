package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map", func() {
	Describe("Required", func() {

		When("MapStringToInterface", func() {
			It("should empty validation", func() {
				d := map[string]interface{}{
					"name": "Chico Bento",
				}

				m := models.TypeMapStringToInterface{
					MapStringToInterface:        d,
					PointerMapStringToInterface: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.TypeMapStringToInterface{}

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

				m := models.TypeMapStringToInterface{
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

				m := models.TypeMapStringToInterface{
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

				m := models.TypeMapStringToString{
					MapStringToString:        d,
					PointerMapStringToString: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.TypeMapStringToString{}

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

				m := models.TypeMapStringToString{
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

				m := models.TypeMapStringToString{
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

				m := models.TypeMapIntToString{
					MapIntToString:        d,
					PointerMapIntToString: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.TypeMapIntToString{}

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

				m := models.TypeMapIntToString{
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

				m := models.TypeMapIntToString{
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

				m := models.TypeMapIntToBool{
					MapIntToBool:        d,
					PointerMapIntToBool: &d,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.TypeMapIntToBool{}

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

				m := models.TypeMapIntToBool{
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

				m := models.TypeMapIntToBool{
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

				m := models.TypeMapIntToStruct{
					MapIntToStructA:               a,
					PointerMapIntToStructA:        aPtr,
					MapIntToStructB:               &b,
					PointerMapIntToStructPointerB: &bPtr,
				}

				errs := m.Validate()
				Expect(errs).To(HaveLen(0))
			})

			It("should fail to validate empty struct", func() {
				m := models.TypeMapIntToStruct{}

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

				m := models.TypeMapIntToStruct{
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

				m := models.TypeMapIntToStruct{
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

				m := models.TypeMapIntToStruct{
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

				m := models.TypeMapIntToStruct{
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
