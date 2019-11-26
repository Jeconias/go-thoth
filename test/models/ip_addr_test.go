package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IP Addr", func() {

	When("ip_addr", func() {
		var (
			s     = "138.21.71.114"
			slice = []string{s, "2da1:9e75:a1c0:1dc3:3707:18f1:34d7:9480"}
		)

		It("should empty validation", func() {
			m := models.IPAddrValidate{
				IPAddr:  s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate ip_addr when ip_addr invalid", func() {
			ss := "978-63"
			m := models.IPAddrValidate{
				IPAddr:  s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip_addr"))
		})

		It("should fail validate ip_addr when Slice ip_addr invalid", func() {
			m := models.IPAddrValidate{
				IPAddr:  s,
				Pointer: &s,
				Slice:   []string{s, s, "978-6"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("ip_addr"))
		})

		It("should fail to validate without field `IP`", func() {
			m := models.IPAddrValidate{
				// IPAddr: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("IPAddr"))
			Expect(errs[0].Tag()).To(Equal("ip_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.IPAddrValidate{
				IPAddr: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip_addr"))
		})
	})

	When("ip4 addr", func() {
		var (
			s     = "155.56.202.241"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.IP4AddrValidate{
				IP4Addr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate ip4 addr when ip4 addr invalid", func() {
			ss := "3054:e41a:f507:1da1:2ea7:0383:3a62:729a"
			m := models.IP4AddrValidate{
				IP4Addr: s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip4_addr"))
		})

		It("should fail validate ip4 addr when Slice ip4 addr invalid", func() {
			m := models.IP4AddrValidate{
				IP4Addr: s,
				Pointer: &s,
				Slice:   []string{"1397:37a8:fffc:dc0c:0fee:e42a:7ab6:93fd"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("ip4_addr"))
		})

		It("should fail to validate without field `IP`", func() {
			m := models.IP4AddrValidate{
				// IP4Addr: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("IP4Addr"))
			Expect(errs[0].Tag()).To(Equal("ip4_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.IP4AddrValidate{
				IP4Addr: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip4_addr"))
		})
	})

	When("ip6_addr", func() {
		var (
			s     = "e567:f79c:95c7:acdf:bfe4:f9cf:6194:fec9"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.IP6AddrValidate{
				IP6Addr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate ip6_addr when ip6_addr invalid", func() {
			ss := "78.227.208.1"
			m := models.IP6AddrValidate{
				IP6Addr: s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip6_addr"))
		})

		It("should fail validate ip6_addr when slice ip6_addr invalid", func() {
			m := models.IP6AddrValidate{
				IP6Addr: s,
				Pointer: &s,
				Slice:   []string{"dc17:9083:bed1:e7a7:41f3:f7be:a216:c8d7", "8.191.228.80"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("ip6_addr"))
		})

		It("should fail to validate without field `IPv6`", func() {
			m := models.IP6AddrValidate{
				// IP6Addr: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("IP6Addr"))
			Expect(errs[0].Tag()).To(Equal("ip6_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.IP6AddrValidate{
				IP6Addr: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("ip6_addr"))
		})
	})

})
