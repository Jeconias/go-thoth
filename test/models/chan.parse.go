package models

// TypeChanString TODO
type TypeChanString struct {
	ChanString             chan string    `thoth:"required"`
	PointerChanString      chan *string   `thoth:"required"`
	ChanSliceString        chan []string  `thoth:"required"`
	ChanSlicePointerString chan []*string `thoth:"required"`
}

// TypeChanUint TODO
type TypeChanUint struct {
	ChanUint             chan uint    `thoth:"required"`
	PointerChanUint      chan *uint   `thoth:"required"`
	ChanSliceUint        chan []uint  `thoth:"required"`
	ChanSlicePointerUint chan []*uint `thoth:"required"`
}
