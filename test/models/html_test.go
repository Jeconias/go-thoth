package models_test

import (
	"github.com/lab259/go-thoth/test/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTML", func() {

	When("html", func() {
		var (
			s = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    Hello World
</body>
</html>
			`
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.HTMLValidate{
				HTML:    s,
				Pointer: &s,
				Slice:   slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate html when html invalid", func() {
			ss := "978-63"
			m := models.HTMLValidate{
				HTML:    s,
				Pointer: &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("html"))
		})

		It("should fail validate html when Slice html invalid", func() {
			m := models.HTMLValidate{
				HTML:    s,
				Pointer: &s,
				Slice:   []string{s, "invalid@name"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("html"))
		})

		It("should fail to validate without field `HTML`", func() {
			m := models.HTMLValidate{
				// HTML: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("HTML"))
			Expect(errs[0].Tag()).To(Equal("html"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.HTMLValidate{
				HTML: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("html"))
		})
	})

	When("html Encoded", func() {
		var (
			s = `
&lt;!DOCTYPE html&gt;
&lt;html lang="en"&gt;
&lt;head&gt;
    &lt;meta charset="UTF-8"&gt;
    &lt;meta name="viewport" content="width=device-width, initial-scale=1.0"&gt;
    &lt;meta http-equiv="X-UA-Compatible" content="ie=edge"&gt;
    &lt;title&gt;Document&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
    Hello World
&lt;/body&gt;
&lt;/html&gt;
			`
			slice = []string{s}
		)

		It("should empty validation", func() {
			m := models.HTMLEncodedValidate{
				HTMLEncoded: s,
				Pointer:     &s,
				Slice:       slice,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(0))
		})

		It("should fail validate html when html invalid", func() {
			ss := "invalid#html"
			m := models.HTMLEncodedValidate{
				HTMLEncoded: s,
				Pointer:     &ss,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("html_encoded"))
		})

		It("should fail validate html when Slice html invalid", func() {
			m := models.HTMLEncodedValidate{
				HTMLEncoded: s,
				Pointer:     &s,
				Slice:       []string{"maximobiz@go"},
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Slice"))
			Expect(errs[0].Tag()).To(Equal("html_encoded"))
		})

		It("should fail to validate without field `HTMLEncoded`", func() {
			m := models.HTMLEncodedValidate{
				// HTMLEncoded: s,
				Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("HTMLEncoded"))
			Expect(errs[0].Tag()).To(Equal("html_encoded"))
		})

		It("should fail to validater without field `Pointer`", func() {
			m := models.HTMLEncodedValidate{
				HTMLEncoded: s,
				// Pointer: &s,
			}

			errs := m.Validate()
			Expect(errs).To(HaveLen(1))
			Expect(errs[0].Field()).To(Equal("Pointer"))
			Expect(errs[0].Tag()).To(Equal("html_encoded"))
		})
	})

})
