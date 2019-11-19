package models_test

import (
	"os"

	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dir", func() {

	var (
		s, _  = os.Getwd()
		slice = []string{s}
	)

	It("should empty validation", func() {
		m := models.DirValidate{
			Dir:     s,
			Pointer: &s,
			Slice:   slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail validate dir when dir invalid", func() {
		ss := "/home/host/workspace"
		m := models.DirValidate{
			Dir:     s,
			Pointer: &ss,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("dir"))
	})

	It("should fail validate dir when Slice dirs invalid", func() {
		m := models.DirValidate{
			Dir:     s,
			Pointer: &s,
			Slice:   []string{s, "/home/host/workspace"},
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Slice"))
		Expect(errs[0].Tag()).To(Equal("dir"))
	})

	It("should fail to validate without field `Dir`", func() {
		m := models.DirValidate{
			// Dir: s,
			Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Dir"))
		Expect(errs[0].Tag()).To(Equal("dir"))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.DirValidate{
			Dir: s,
			// Pointer: &s,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("dir"))
	})

})
