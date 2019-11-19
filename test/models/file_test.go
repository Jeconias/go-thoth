package models_test

import (
	"os"
	"path"

	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {

	var (
		dir, _ = os.Getwd()
		s      = path.Join(dir, "generated.go")
		slice  = []string{s}
	)

	It("should empty validation", func() {
		m := models.FileValidate{
			File:    s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate file when file invalid", func() {
		ss := "invalid-file.txt"
		m := models.FileValidate{
			File:    s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("file"))
	})

	It("should fail validate file when files contais file invalid", func() {
		ss := []string{s, "invalid-file.txt"}
		m := models.FileValidate{
			File:    s,
			Pointer: &s,
			Slice:   ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("file"))
	})

	It("should fail to validate without field `File`", func() {
		m := models.FileValidate{
			// File: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("File"))
		Expect(errs[0].Tag()).To(Equal("file"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.FileValidate{
			File: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("file"))
	})

})
