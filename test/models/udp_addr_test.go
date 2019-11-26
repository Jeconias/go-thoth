package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UDPAddr", func() {

	When("udp_addr", func() {
		var (
			s     = "127.0.0.1:8080"
			slice = []string{s, "[dc17:9083:bed1:e7a7:41f3:f7be:a216:c8d7]:8081"}
		)

		It("should empty validation", func() {
			m := models.UDPAddrValidate{
				UDPAddr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate udp_addr when udp_addr invalid", func() {
			ss := "978-63"
			m := models.UDPAddrValidate{
				UDPAddr: s,
				Pointer: &ss,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("udp_addr"))
		})

		It("should fail validate udp_addr when Slice udp_addr invalid", func() {
			m := models.UDPAddrValidate{
				UDPAddr: s,
				Pointer: &s,
				Slice:   []string{s, "978-6"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("udp_addr"))
		})

		It("should fail to validate without field `UDPAddr`", func() {
			m := models.UDPAddrValidate{
				// UDPAddr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UDPAddr"))
			Expect(errs[0].Tag()).To(Equal("udp_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UDPAddrValidate{
				UDPAddr: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("udp_addr"))
		})
	})

	When("udp4_addr", func() {
		var (
			s     = "127.0.0.1:8080"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UDP4AddrValidate{
				UDP4Addr: s,
				Pointer:  &s,
				Slice:    slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate udp4_addr when udp4_addr invalid", func() {
			ss := "3054:e41a:f507:1da1:2ea7:0383:3a62:729a"
			m := models.UDP4AddrValidate{
				UDP4Addr: s,
				Pointer:  &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("udp4_addr"))
		})

		It("should fail validate udp4_addr when Slice udp4_addr invalid", func() {
			m := models.UDP4AddrValidate{
				UDP4Addr: s,
				Pointer:  &s,
				Slice:    []string{"1397:37a8:fffc:dc0c:0fee:e42a:7ab6:93fd"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("udp4_addr"))
		})

		It("should fail to validate without field `IP`", func() {
			m := models.UDP4AddrValidate{
				// UDP4Addr: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UDP4Addr"))
			Expect(errs[0].Tag()).To(Equal("udp4_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UDP4AddrValidate{
				UDP4Addr: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("udp4_addr"))
		})
	})

	When("udp6_addr", func() {
		var (
			s     = "[1fff:0:a88:85a3::ac1f]:8001"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UDP6AddrValidate{
				UDP6Addr: s,
				Pointer:  &s,
				Slice:    slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate udp6_addr when udp6_addr invalid", func() {
			ss := "78.227.208.1"
			m := models.UDP6AddrValidate{
				UDP6Addr: s,
				Pointer:  &ss,
				Slice:    slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("udp6_addr"))
		})

		It("should fail validate udp6_addr when slice udp6_addr invalid", func() {
			m := models.UDP6AddrValidate{
				UDP6Addr: s,
				Pointer:  &s,
				Slice:    []string{"[dc17:9083:bed1:e7a7:41f3:f7be:a216:c8d7]:8081", "8.191.228.80"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("udp6_addr"))
		})

		It("should fail to validate without field `UDP6Addr`", func() {
			m := models.UDP6AddrValidate{
				// UDP6Addr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UDP6Addr"))
			Expect(errs[0].Tag()).To(Equal("udp6_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UDP6AddrValidate{
				UDP6Addr: s,
				// Pointer: &s,
				Slice: slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("udp6_addr"))
		})
	})

})
