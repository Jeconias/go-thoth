package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ISBN", func() {

	When("isbn", func() {
		var (
			s10   = "0-590-76484-5"
			s13   = "978-85-7254-036-0"
			slice = []string{s10, s13}
		)

		It("should empty validation", func() {
			m := models.ISBNValidate{
				ISBN:    s10,
				Pointer: &s13,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate isbn when isbn invalid", func() {
			ss := "978-610"
			m := models.ISBNValidate{
				ISBN:    s13,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("isbn"))
		})

		It("should fail validate isbn when Slice isbn invalid", func() {
			m := models.ISBNValidate{
				ISBN:    s13,
				Pointer: &s10,
				Slice:   []string{s10, s13, "978-610"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("isbn"))
		})

		It("should fail to validate without field `ISBN`", func() {
			m := models.ISBNValidate{
				// ISBN: s,
				Pointer: &s13,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ISBN"))
			Expect(errs[0].Tag()).To(Equal("isbn"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.ISBNValidate{
				ISBN: s10,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("isbn"))
		})
	})

	When("isbn10", func() {
		var (
			s     = " 0-590-76484-5"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.ISBN10Validate{
				ISBN10:  s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate isbn10 when isbn10 invalid", func() {
			ss := "978-617-7754-09-0"
			m := models.ISBN10Validate{
				ISBN10:  s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("isbn10"))
		})

		It("should fail validate isbn10 when Slice isbn10 invalid", func() {
			m := models.ISBN10Validate{
				ISBN10:  s,
				Pointer: &s,
				Slice:   []string{"978-617-7754-09-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("isbn10"))
		})

		It("should fail to validate without field `ISBN`", func() {
			m := models.ISBN10Validate{
				// ISBN10: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ISBN10"))
			Expect(errs[0].Tag()).To(Equal("isbn10"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.ISBN10Validate{
				ISBN10: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("isbn10"))
		})
	})

	When("isbn13", func() {
		var (
			s     = "978-85-7254-036-0"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.ISBN13Validate{
				ISBN13:  s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate isbn13 when isbn13 invalid", func() {
			ss := "978-617-77-0"
			m := models.ISBN13Validate{
				ISBN13:  s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("isbn13"))
		})

		It("should fail validate isbn13 when slice isbn13 invalid", func() {
			m := models.ISBN13Validate{
				ISBN13:  s,
				Pointer: &s,
				Slice:   []string{"978-85-7254-036-0", "978-617-77-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("isbn13"))
		})

		It("should fail to validate without field `ISBN13`", func() {
			m := models.ISBN13Validate{
				// ISBN13: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("ISBN13"))
			Expect(errs[0].Tag()).To(Equal("isbn13"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.ISBN13Validate{
				ISBN13: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("isbn13"))
		})
	})

})
