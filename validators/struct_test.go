package validators_test

import (
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/validators"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/go-playground/validator.v9"
)

var _ = Describe("Validator", func() {
	It("should to check if struct implements validator `StructLevel` interface", func() {
		var t validator.StructLevel = new(validators.Level)
		var i interface{} = t

		_, ok := i.(validator.StructLevel)
		Expect(ok).To(BeTrue())
	})

	It("should to verify struct validation", func() {
		// 1. create new an environment
		env, err := myasthurts.NewEnvironment()
		Expect(err).ToNot(HaveOccurred())

		// 2. loading the file
		pkg, err := env.ParseDir("../test/validators")
		Expect(err).ToNot(HaveOccurred())
		s, _ := pkg.StructByName("Person")

		Expect(s.Fields).To(HaveLen(2))

		tagParam1 := s.Fields[0].Tag.TagParamByName("validate")

		Expect(tagParam1.Name).To(Equal("validate"))
		Expect(tagParam1.Value).To(Equal("required"))
		Expect(tagParam1.Options).To(Equal([]string{"required"}))

		tagParam2 := s.Fields[1].Tag.TagParamByName("validate")
		Expect(tagParam2.Value).To(Equal("-"))

		// err = env.parseFile(newDataPackageContext(env), "test/validators/person.sample.go")

		// // 3. loading packages
		// pkg, ok := env.PackageByImportPath("models")
		// Expect(ok).To(BeTrue())
		// Expect(pkg).NotTo(BeNil())

		// Expect(pkg.Structs).To(HaveLen(1))
		// Expect(pkg.Structs[0].Fields).To(HaveLen(2))
		// Expect(pkg.Structs[0].Fields[0].Tag.Params[0].Value).To(Equal("required"))

		// s := models.Person{
		// 	Name: "Chico Bento",
		// 	Age:  22,
		// }

		// var _ = myasthurts.ErrBuiltInNotFound

		// Expect(t.Validator().Struct(s)).To(Succeed())

		// var t validator.StructLevel = new(validators.Level)
		// var i interface{} = t

		// _, ok := i.(validator.StructLevel)
		// Expect(ok).To(BeTrue())
	})
})
