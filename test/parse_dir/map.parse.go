package any

// MapStringInterface TODO
type MapStringInterface struct {
	A map[string]interface{}  `thoth:"required"`
	B *map[string]interface{} `thoth:"required"`
}
