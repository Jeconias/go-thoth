package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CIDR", func() {

	When("cidr", func() {
		var (
			s     = "138.21.71.114/24"
			slice = []string{s, "2da1:9e75:a1c0:1dc3:3707:18f1:34d7:9480/48"}
		)

		It("should empty validation", func() {
			m := models.CIDRValidate{
				CIDR:    s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate cidr when cidr invalid", func() {
			ss := "978-63"
			m := models.CIDRValidate{
				CIDR:    s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("cidr"))
		})

		It("should fail validate cidr when Slice cidr invalid", func() {
			m := models.CIDRValidate{
				CIDR:    s,
				Pointer: &s,
				Slice:   []string{s, s, "978-6"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("cidr"))
		})

		It("should fail to validate without field `CIDR`", func() {
			m := models.CIDRValidate{
				// CIDR: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("CIDR"))
			Expect(errs[0].Tag()).To(Equal("cidr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.CIDRValidate{
				CIDR: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("cidr"))
		})
	})

	When("cidrv4", func() {
		var (
			s     = "155.56.202.241/12"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.CIDRv4Validate{
				CIDRv4:  s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate cidrv4 when cidrv4 invalid", func() {
			ss := "3054:e41a:f507:1da1:2ea7:0383:3a62:729a"
			m := models.CIDRv4Validate{
				CIDRv4:  s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("cidrv4"))
		})

		It("should fail validate cidrv4 when Slice cidrv4 invalid", func() {
			m := models.CIDRv4Validate{
				CIDRv4:  s,
				Pointer: &s,
				Slice:   []string{"1397:37a8:fffc:dc0c:0fee:e42a:7ab6:93fd"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("cidrv4"))
		})

		It("should fail to validate without field `CIDR`", func() {
			m := models.CIDRv4Validate{
				// CIDRv4: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("CIDRv4"))
			Expect(errs[0].Tag()).To(Equal("cidrv4"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.CIDRv4Validate{
				CIDRv4: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("cidrv4"))
		})
	})

	When("cidrv6", func() {
		var (
			s     = "e567:f79c:95c7:acdf:bfe4:f9cf:6194:fec9/35"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.CIDRv6Validate{
				CIDRv6:  s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate cidrv6 when cidrv6 invalid", func() {
			ss := "78.227.208.1"
			m := models.CIDRv6Validate{
				CIDRv6:  s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("cidrv6"))
		})

		It("should fail validate cidrv6 when slice cidrv6 invalid", func() {
			m := models.CIDRv6Validate{
				CIDRv6:  s,
				Pointer: &s,
				Slice:   []string{"dc17:9083:bed1:e7a7:41f3:f7be:a216:c8d7/12", "8.191.228.80/21"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("cidrv6"))
		})

		It("should fail to validate without field `CIDRv6`", func() {
			m := models.CIDRv6Validate{
				// CIDRv6: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("CIDRv6"))
			Expect(errs[0].Tag()).To(Equal("cidrv6"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.CIDRv6Validate{
				CIDRv6: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("cidrv6"))
		})
	})

})
