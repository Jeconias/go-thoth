package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UUID", func() {

	When("uuid", func() {
		var (
			s     = "f4ebe677-045f-4369-9484-910e1860baa5"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUIDValidate{
				UUID:    s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid when uuid invalid", func() {
			ss := "978-63"
			m := models.UUIDValidate{
				UUID:    s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid"))
		})

		It("should fail validate uuid when Slice uuid invalid", func() {
			m := models.UUIDValidate{
				UUID:    s,
				Pointer: &s,
				Slice:   []string{s, s, "978-6"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid"))
		})

		It("should fail to validate without field `UUID`", func() {
			m := models.UUIDValidate{
				// UUID: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID"))
			Expect(errs[0].Tag()).To(Equal("uuid"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUIDValidate{
				UUID: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid"))
		})
	})

	When("uuid3", func() {
		var (
			s     = "e13a319e-16d9-3ff5-a83c-96564007998e"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUID3Validate{
				UUID3:   s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid3 when uuid3 invalid", func() {
			ss := "978-617-7754-09-0"
			m := models.UUID3Validate{
				UUID3:   s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid3"))
		})

		It("should fail validate uuid3 when Slice uuid3 invalid", func() {
			m := models.UUID3Validate{
				UUID3:   s,
				Pointer: &s,
				Slice:   []string{"978-617-7754-09-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid3"))
		})

		It("should fail to validate without field `UUID`", func() {
			m := models.UUID3Validate{
				// UUID3: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID3"))
			Expect(errs[0].Tag()).To(Equal("uuid3"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUID3Validate{
				UUID3: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid3"))
		})
	})

	When("uuid4", func() {
		var (
			s     = "ebfd1c49-3118-4a7d-8c7b-8cfa7706b7c2"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUID4Validate{
				UUID4:   s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid4 when uuid4 invalid", func() {
			ss := "978-617-77-0"
			m := models.UUID4Validate{
				UUID4:   s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid4"))
		})

		It("should fail validate uuid4 when slice uuid4 invalid", func() {
			m := models.UUID4Validate{
				UUID4:   s,
				Pointer: &s,
				Slice:   []string{"63d2415a-a072-4177-9c31-e7b96a3d3198", "978-617-77-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid4"))
		})

		It("should fail to validate without field `UUID4`", func() {
			m := models.UUID4Validate{
				// UUID4: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID4"))
			Expect(errs[0].Tag()).To(Equal("uuid4"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUID4Validate{
				UUID4: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid4"))
		})
	})

	When("uuid5", func() {
		var (
			s     = "c74a196f-f19d-5ea9-bffd-a2742432fc9c"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUID5Validate{
				UUID5:   s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid5 when uuid5 invalid", func() {
			ss := "978-617-77-0"
			m := models.UUID5Validate{
				UUID5:   s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid5"))
		})

		It("should fail validate uuid5 when slice uuid5 invalid", func() {
			m := models.UUID5Validate{
				UUID5:   s,
				Pointer: &s,
				Slice:   []string{"c74a196f-f19d-5ea9-bffd-a2742432fc9c", "978-617-77-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid5"))
		})

		It("should fail to validate without field `UUID5`", func() {
			m := models.UUID5Validate{
				// UUID5: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID5"))
			Expect(errs[0].Tag()).To(Equal("uuid5"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUID5Validate{
				UUID5: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid5"))
		})
	})

	When("uuid_rfc4122", func() {
		var (
			s     = "f4ebe677-045f-4369-9484-910e1860baa5"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUIDrfc4122Validate{
				UUID:    s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid when uuid invalid", func() {
			ss := "978-63"
			m := models.UUIDrfc4122Validate{
				UUID:    s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid_rfc4122"))
		})

		It("should fail validate uuid when Slice uuid invalid", func() {
			m := models.UUIDrfc4122Validate{
				UUID:    s,
				Pointer: &s,
				Slice:   []string{s, s, "978-6"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid_rfc4122"))
		})

		It("should fail to validate without field `UUID`", func() {
			m := models.UUIDrfc4122Validate{
				// UUID: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID"))
			Expect(errs[0].Tag()).To(Equal("uuid_rfc4122"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUIDrfc4122Validate{
				UUID: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid_rfc4122"))
		})
	})

	When("uuid3_rfc4122", func() {
		var (
			s     = "e13a319e-16d9-3ff5-a83c-96564007998e"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUID3rfc4122Validate{
				UUID3rfc4122: s,
				Pointer:      &s,
				Slice:        slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid3 when uuid3 invalid", func() {
			ss := "978-617-7754-09-0"
			m := models.UUID3rfc4122Validate{
				UUID3rfc4122: s,
				Pointer:      &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid3_rfc4122"))
		})

		It("should fail validate uuid3 when Slice uuid3 invalid", func() {
			m := models.UUID3rfc4122Validate{
				UUID3rfc4122: s,
				Pointer:      &s,
				Slice:        []string{"978-617-7754-09-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid3_rfc4122"))
		})

		It("should fail to validate without field `UUID`", func() {
			m := models.UUID3rfc4122Validate{
				// UUID3rfc4122: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID3rfc4122"))
			Expect(errs[0].Tag()).To(Equal("uuid3_rfc4122"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUID3rfc4122Validate{
				UUID3rfc4122: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid3_rfc4122"))
		})
	})

	When("uuid4_rfc4122", func() {
		var (
			s     = "ebfd1c49-3118-4a7d-8c7b-8cfa7706b7c2"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUID4rfc4122Validate{
				UUID4rfc4122: s,
				Pointer:      &s,
				Slice:        slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid4 when uuid4 invalid", func() {
			ss := "978-617-77-0"
			m := models.UUID4rfc4122Validate{
				UUID4rfc4122: s,
				Pointer:      &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid4_rfc4122"))
		})

		It("should fail validate uuid4 when slice uuid4 invalid", func() {
			m := models.UUID4rfc4122Validate{
				UUID4rfc4122: s,
				Pointer:      &s,
				Slice:        []string{"63d2415a-a072-4177-9c31-e7b96a3d3198", "978-617-77-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid4_rfc4122"))
		})

		It("should fail to validate without field `UUID4`", func() {
			m := models.UUID4rfc4122Validate{
				// UUID4rfc4122: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID4rfc4122"))
			Expect(errs[0].Tag()).To(Equal("uuid4_rfc4122"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUID4rfc4122Validate{
				UUID4rfc4122: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid4_rfc4122"))
		})
	})

	When("uuid5_rfc4122", func() {
		var (
			s     = "c74a196f-f19d-5ea9-bffd-a2742432fc9c"
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.UUID5rfc4122Validate{
				UUID5rfc4122: s,
				Pointer:      &s,
				Slice:        slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate uuid5 when uuid5 invalid", func() {
			ss := "978-617-77-0"
			m := models.UUID5rfc4122Validate{
				UUID5rfc4122: s,
				Pointer:      &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid5_rfc4122"))
		})

		It("should fail validate uuid5 when slice uuid5 invalid", func() {
			m := models.UUID5rfc4122Validate{
				UUID5rfc4122: s,
				Pointer:      &s,
				Slice:        []string{"c74a196f-f19d-5ea9-bffd-a2742432fc9c", "978-617-77-0"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("uuid5_rfc4122"))
		})

		It("should fail to validate without field `UUID5rfc4122`", func() {
			m := models.UUID5rfc4122Validate{
				// UUID5rfc4122: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("UUID5rfc4122"))
			Expect(errs[0].Tag()).To(Equal("uuid5_rfc4122"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.UUID5rfc4122Validate{
				UUID5rfc4122: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("uuid5_rfc4122"))
		})
	})

})
