package models

// Person TODO
type Person struct {
	Name string `validate:"required"`
	Age  int    `validate:"-"`
}

func (p *Person) Validate() error {

	return nil
}
