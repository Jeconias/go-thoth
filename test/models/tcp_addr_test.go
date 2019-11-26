package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TCPAddr", func() {

	When("tcp_addr", func() {
		var (
			s     = "127.0.0.1:8080"
			slice = []string{s, "[dc17:9083:bed1:e7a7:41f3:f7be:a216:c8d7]:8081"}
		)

		It("should empty validation", func() {
			m := models.TCPAddrValidate{
				TCPAddr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate tcp_addr when tcp_addr invalid", func() {
			ss := "978-63"
			m := models.TCPAddrValidate{
				TCPAddr: s,
				Pointer: &ss,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("tcp_addr"))
		})

		It("should fail validate tcp_addr when Slice tcp_addr invalid", func() {
			m := models.TCPAddrValidate{
				TCPAddr: s,
				Pointer: &s,
				Slice:   []string{s, "978-6"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("tcp_addr"))
		})

		It("should fail to validate without field `TCPAddr`", func() {
			m := models.TCPAddrValidate{
				// TCPAddr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("TCPAddr"))
			Expect(errs[0].Tag()).To(Equal("tcp_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.TCPAddrValidate{
				TCPAddr: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("tcp_addr"))
		})
	})

	When("tcp4_addr", func() {
		var (
			s     = "127.0.0.1:8080"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.TCP4AddrValidate{
				TCP4Addr: s,
				Pointer:  &s,
				Slice:    slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate tcp4_addr when tcp4_addr invalid", func() {
			ss := "3054:e41a:f507:1da1:2ea7:0383:3a62:729a"
			m := models.TCP4AddrValidate{
				TCP4Addr: s,
				Pointer:  &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("tcp4_addr"))
		})

		It("should fail validate tcp4_addr when Slice tcp4_addr invalid", func() {
			m := models.TCP4AddrValidate{
				TCP4Addr: s,
				Pointer:  &s,
				Slice:    []string{"1397:37a8:fffc:dc0c:0fee:e42a:7ab6:93fd"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("tcp4_addr"))
		})

		It("should fail to validate without field `IP`", func() {
			m := models.TCP4AddrValidate{
				// TCP4Addr: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("TCP4Addr"))
			Expect(errs[0].Tag()).To(Equal("tcp4_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.TCP4AddrValidate{
				TCP4Addr: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("tcp4_addr"))
		})
	})

	When("tcp6_addr", func() {
		var (
			s     = "[1fff:0:a88:85a3::ac1f]:8001"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.TCP6AddrValidate{
				TCP6Addr: s,
				Pointer:  &s,
				Slice:    slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate tcp6_addr when tcp6_addr invalid", func() {
			ss := "78.227.208.1"
			m := models.TCP6AddrValidate{
				TCP6Addr: s,
				Pointer:  &ss,
				Slice:    slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("tcp6_addr"))
		})

		It("should fail validate tcp6_addr when slice tcp6_addr invalid", func() {
			m := models.TCP6AddrValidate{
				TCP6Addr: s,
				Pointer:  &s,
				Slice:    []string{"[dc17:9083:bed1:e7a7:41f3:f7be:a216:c8d7]:8081", "8.191.228.80"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("tcp6_addr"))
		})

		It("should fail to validate without field `TCP6Addr`", func() {
			m := models.TCP6AddrValidate{
				// TCP6Addr: s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("TCP6Addr"))
			Expect(errs[0].Tag()).To(Equal("tcp6_addr"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.TCP6AddrValidate{
				TCP6Addr: s,
				// Pointer: &s,
				Slice: slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("tcp6_addr"))
		})
	})

})
