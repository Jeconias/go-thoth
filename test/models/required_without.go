package models

// RequiredWithoutField TODO
type RequiredWithoutField struct {
	Status bool   `thoth:"required"`
	Name   string `thoth:"required_without=Status"`
}

// RequiredWithoutFieldStrPointer TODO
type RequiredWithoutFieldStrPointer struct {
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_without=Status"`
}

// RequiredWithoutFields TODO
type RequiredWithoutFields struct {
	ID     int     `thoth:"eq=1"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_without=ID Status"`
}

// RequiredWithoutConfirmation TODO
type RequiredWithoutConfirmation struct {
	URL              string `thoth:"url"`
	NeedConfirmation bool   `thoth:"required_without=URL"`
}
