package models

// TypeInterface TODO
type TypeInterface struct {
	Interface        interface{}  `thoth:"required"`
	PointerInterface *interface{} `thoth:"required"`
}
