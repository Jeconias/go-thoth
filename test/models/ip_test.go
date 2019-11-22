package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IP", func() {

	When("ip", func() {
		var (
			s     = "138.21.71.114"
			slice = []string{s, "2da1:9e75:a1c0:1dc3:3707:18f1:34d7:9480"}
		)

		It("should empty validation", func() {
			m := models.IPValidate{
				IP:      s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate ip when ip invalid", func() {
			ss := "978-63"
			m := models.IPValidate{
				IP:      s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip"))
		})

		It("should fail validate ip when Slice ip invalid", func() {
			m := models.IPValidate{
				IP:      s,
				Pointer: &s,
				Slice:   []string{s, s, "978-6"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("ip"))
		})

		It("should fail to validate without field `IP`", func() {
			m := models.IPValidate{
				// IP: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("IP"))
			Expect(errs[0].Tag()).To(Equal("ip"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.IPValidate{
				IP: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip"))
		})
	})

	When("ipv4", func() {
		var (
			s     = "155.56.202.241"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.IPv4Validate{
				IPv4:    s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate ipv4 when ipv4 invalid", func() {
			ss := "3054:e41a:f507:1da1:2ea7:0383:3a62:729a"
			m := models.IPv4Validate{
				IPv4:    s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ipv4"))
		})

		It("should fail validate ipv4 when Slice ipv4 invalid", func() {
			m := models.IPv4Validate{
				IPv4:    s,
				Pointer: &s,
				Slice:   []string{"1397:37a8:fffc:dc0c:0fee:e42a:7ab6:93fd"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("ipv4"))
		})

		It("should fail to validate without field `IP`", func() {
			m := models.IPv4Validate{
				// IPv4: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("IPv4"))
			Expect(errs[0].Tag()).To(Equal("ipv4"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.IPv4Validate{
				IPv4: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ipv4"))
		})
	})

	When("ipv6", func() {
		var (
			s     = "e567:f79c:95c7:acdf:bfe4:f9cf:6194:fec9"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.IPv6Validate{
				IPv6:    s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate ipv6 when ipv6 invalid", func() {
			ss := "78.227.208.1"
			m := models.IPv6Validate{
				IPv6:    s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ipv6"))
		})

		It("should fail validate ipv6 when slice ipv6 invalid", func() {
			m := models.IPv6Validate{
				IPv6:    s,
				Pointer: &s,
				Slice:   []string{"dc17:9083:bed1:e7a7:41f3:f7be:a216:c8d7", "8.191.228.80"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("ipv6"))
		})

		It("should fail to validate without field `IPv6`", func() {
			m := models.IPv6Validate{
				// IPv6: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("IPv6"))
			Expect(errs[0].Tag()).To(Equal("ipv6"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.IPv6Validate{
				IPv6: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ipv6"))
		})
	})

})
