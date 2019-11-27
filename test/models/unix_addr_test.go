package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UNIXAddr", func() {

	var (
		s     = "golang.org:80"
		slice = []string{s}
	)

	It("should empty validation", func() {
		m := models.UNIXAddrValidate{
			UNIXAddr: s,
			Pointer:  &s,
			Slice:    slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(0))
	})

	It("should fail to validater without field `Pointer`", func() {
		m := models.UNIXAddrValidate{
			UNIXAddr: s,
			// Pointer: &s,
			Slice: slice,
		}

		errs := m.Validate()
		Expect(errs).To(HaveLen(1))
		Expect(errs[0].Field()).To(Equal("Pointer"))
		Expect(errs[0].Tag()).To(Equal("unix_addr"))
	})

})
