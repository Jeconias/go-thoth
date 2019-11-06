package models

// TypeMapStringToInterface TODO
type TypeMapStringToInterface struct {
	MapStringToInterface        map[string]interface{}  `thoth:"required"`
	PointerMapStringToInterface *map[string]interface{} `thoth:"required"`
}

// TypeMapStringToString TODO
type TypeMapStringToString struct {
	MapStringToString        map[string]string  `thoth:"required"`
	PointerMapStringToString *map[string]string `thoth:"required"`
}

// TypeMapIntToString TODO
type TypeMapIntToString struct {
	MapIntToString        map[int]string  `thoth:"required"`
	PointerMapIntToString *map[int]string `thoth:"required"`
}

// TypeMapIntToBool TODO
type TypeMapIntToBool struct {
	MapIntToBool        map[int]bool  `thoth:"required"`
	PointerMapIntToBool *map[int]bool `thoth:"required"`
}

// MapTypeA TODO
type MapTypeA struct {
	Int  int
	Bool bool
}

// MapTypeB TODO
type MapTypeB struct {
	Bool   bool
	String string
	Float  float64
}

// TypeMapIntToStruct TODO
type TypeMapIntToStruct struct {
	MapIntToStructA               map[int]MapTypeA   `thoth:"required"`
	PointerMapIntToStructA        map[int]*MapTypeA  `thoth:"required"`
	MapIntToStructB               *map[int]MapTypeB  `thoth:"required"`
	PointerMapIntToStructPointerB *map[int]*MapTypeB `thoth:"required"`
}
