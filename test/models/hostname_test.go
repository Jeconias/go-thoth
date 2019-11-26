package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hostname", func() {

	When("hostname RFC952", func() {
		var (
			s     = "garth.net"
			slice = []string{s, "charley.name"}
		)

		It("should empty validation", func() {
			m := models.HostnameRFC952Validate{
				HostnameRFC952: s,
				Pointer:        &s,
				Slice:          slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate hostname when hostname invalid", func() {
			ss := "978-63"
			m := models.HostnameRFC952Validate{
				HostnameRFC952: s,
				Pointer:        &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("hostname"))
		})

		It("should fail validate hostname when Slice hostname invalid", func() {
			m := models.HostnameRFC952Validate{
				HostnameRFC952: s,
				Pointer:        &s,
				Slice:          []string{s, "invalid@name"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("hostname"))
		})

		It("should fail to validate without field `HostnameRFC952`", func() {
			m := models.HostnameRFC952Validate{
				// HostnameRFC952: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("HostnameRFC952"))
			Expect(errs[0].Tag()).To(Equal("hostname"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.HostnameRFC952Validate{
				HostnameRFC952: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("hostname"))
		})
	})

	When("hostname RFC1123", func() {
		var (
			s     = "august.name.org"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.HostnameRFC1123Validate{
				HostnameRFC1123: s,
				Pointer:         &s,
				Slice:           slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate hostname when hostname invalid", func() {
			ss := "invalid#hostname"
			m := models.HostnameRFC1123Validate{
				HostnameRFC1123: s,
				Pointer:         &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("hostname_rfc1123"))
		})

		It("should fail validate hostname when Slice hostname invalid", func() {
			m := models.HostnameRFC1123Validate{
				HostnameRFC1123: s,
				Pointer:         &s,
				Slice:           []string{"maximobiz@go"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("hostname_rfc1123"))
		})

		It("should fail to validate without field `HostnameRFC1123`", func() {
			m := models.HostnameRFC1123Validate{
				// HostnameRFC1123: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("HostnameRFC1123"))
			Expect(errs[0].Tag()).To(Equal("hostname_rfc1123"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.HostnameRFC1123Validate{
				HostnameRFC1123: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("hostname_rfc1123"))
		})
	})

})
