package models

// MapStringToInterface TODO
type MapStringToInterface struct {
	A map[string]interface{}  `thoth:"required"`
	B *map[string]interface{} `thoth:"required"`
}

// MapStringToString TODO
type MapStringToString struct {
	A map[string]string  `thoth:"required"`
	B *map[string]string `thoth:"required"`
}

// MapIntToString TODO
type MapIntToString struct {
	A map[int]string  `thoth:"required"`
	B *map[int]string `thoth:"required"`
}

// MapIntToBool TODO
type MapIntToBool struct {
	A map[int]bool  `thoth:"required"`
	B *map[int]bool `thoth:"required"`
}

// MapTypeA TODO
type MapTypeA struct {
	B int
	C bool
}

// MapTypeB TODO
type MapTypeB struct {
	C bool
	D string
	E float64
}

// MapIntToStruct TODO
type MapIntToStruct struct {
	A        map[int]MapTypeA   `thoth:"required"`
	APointer map[int]*MapTypeA  `thoth:"required"`
	B        *map[int]MapTypeB  `thoth:"required"`
	BPointer *map[int]*MapTypeB `thoth:"required"`
}
