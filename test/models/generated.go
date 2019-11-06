package models

// Validate TODO
func (t *TypeChanString) Validate() (errs ValidationErrors) {
	if Empty(len(t.ChanString)) {
		errs = append(errs, ErrEmpty("ChanString", "required"))
	}
	if Empty(len(t.PointerChanString)) {
		errs = append(errs, ErrEmpty("PointerChanString", "required"))
	}
	if Empty(len(t.ChanSliceString)) {
		errs = append(errs, ErrEmpty("ChanSliceString", "required"))
	}
	if Empty(len(t.ChanSlicePointerString)) {
		errs = append(errs, ErrEmpty("ChanSlicePointerString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeChanUint) Validate() (errs ValidationErrors) {
	if Empty(len(t.ChanUint)) {
		errs = append(errs, ErrEmpty("ChanUint", "required"))
	}
	if Empty(len(t.PointerChanUint)) {
		errs = append(errs, ErrEmpty("PointerChanUint", "required"))
	}
	if Empty(len(t.ChanSliceUint)) {
		errs = append(errs, ErrEmpty("ChanSliceUint", "required"))
	}
	if Empty(len(t.ChanSlicePointerUint)) {
		errs = append(errs, ErrEmpty("ChanSlicePointerUint", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeInterface) Validate() (errs ValidationErrors) {
	if IsValid(t.Interface) {
		errs = append(errs, ErrEmpty("Interface", "required"))
	}
	if t.PointerInterface == nil {
		errs = append(errs, ErrEmpty("PointerInterface", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapStringToInterface) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapStringToInterface)) {
		errs = append(errs, ErrEmpty("MapStringToInterface", "required"))
	}
	if t.PointerMapStringToInterface == nil {
		errs = append(errs, ErrEmpty("PointerMapStringToInterface", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapStringToString) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapStringToString)) {
		errs = append(errs, ErrEmpty("MapStringToString", "required"))
	}
	if t.PointerMapStringToString == nil {
		errs = append(errs, ErrEmpty("PointerMapStringToString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapIntToString) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapIntToString)) {
		errs = append(errs, ErrEmpty("MapIntToString", "required"))
	}
	if t.PointerMapIntToString == nil {
		errs = append(errs, ErrEmpty("PointerMapIntToString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapIntToBool) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapIntToBool)) {
		errs = append(errs, ErrEmpty("MapIntToBool", "required"))
	}
	if t.PointerMapIntToBool == nil {
		errs = append(errs, ErrEmpty("PointerMapIntToBool", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapIntToStruct) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapIntToStructA)) {
		errs = append(errs, ErrEmpty("MapIntToStructA", "required"))
	}
	if Empty(len(t.PointerMapIntToStructA)) {
		errs = append(errs, ErrEmpty("PointerMapIntToStructA", "required"))
	}
	if t.MapIntToStructB == nil {
		errs = append(errs, ErrEmpty("MapIntToStructB", "required"))
	}
	if t.PointerMapIntToStructPointerB == nil {
		errs = append(errs, ErrEmpty("PointerMapIntToStructPointerB", "required"))
	}
	return errs
}

// Validate TODO
func (h *Home) Validate() (errs ValidationErrors) {
	if IsInt(h.ID) {
		errs = append(errs, ErrNumberRequired("ID", "required"))
	}
	if IsValid(h.Address) {
		errs = append(errs, ErrEmpty("Address", "required"))
	}
	return errs
}

// Validate TODO
func (c *Client) Validate() (errs ValidationErrors) {
	if IsInt64(c.ID) {
		errs = append(errs, ErrNumberRequired("ID", "required"))
	}
	if Empty(len(c.Name)) {
		errs = append(errs, ErrEmpty("Name", "required"))
	}
	if c.LastName == nil {
		errs = append(errs, ErrEmpty("LastName", "required"))
	}
	if c.Address == nil {
		errs = append(errs, ErrEmpty("Address", "required"))
	}
	return errs
}

// Validate TODO
func (u *User) Validate() (errs ValidationErrors) {
	if Empty(len(u.Name)) {
		errs = append(errs, ErrEmpty("Name", "required"))
	}
	return errs
}

// Validate TODO
func (n *Number) Validate() (errs ValidationErrors) {
	if IsUint(n.Uint) {
		errs = append(errs, ErrNumberRequired("Uint", "required"))
	}
	if n.UintPointer == nil {
		errs = append(errs, ErrNumberRequired("UintPointer", "required"))
	}
	if IsUint8(n.Uint8) {
		errs = append(errs, ErrNumberRequired("Uint8", "required"))
	}
	if n.Uint8Pointer == nil {
		errs = append(errs, ErrNumberRequired("Uint8Pointer", "required"))
	}
	if IsUint16(n.Uint16) {
		errs = append(errs, ErrNumberRequired("Uint16", "required"))
	}
	if n.Uint16Pointer == nil {
		errs = append(errs, ErrNumberRequired("Uint16Pointer", "required"))
	}
	if IsUint32(n.Uint32) {
		errs = append(errs, ErrNumberRequired("Uint32", "required"))
	}
	if n.Uint32Pointer == nil {
		errs = append(errs, ErrNumberRequired("Uint32Pointer", "required"))
	}
	if IsUint64(n.Uint64) {
		errs = append(errs, ErrNumberRequired("Uint64", "required"))
	}
	if n.Uint64Pointer == nil {
		errs = append(errs, ErrNumberRequired("Uint64Pointer", "required"))
	}
	if IsUintptr(n.Uintptr) {
		errs = append(errs, ErrNumberRequired("Uintptr", "required"))
	}
	if n.UintptrPointer == nil {
		errs = append(errs, ErrNumberRequired("UintptrPointer", "required"))
	}
	if IsInt(n.Int) {
		errs = append(errs, ErrNumberRequired("Int", "required"))
	}
	if n.IntPointer == nil {
		errs = append(errs, ErrNumberRequired("IntPointer", "required"))
	}
	if IsInt8(n.Int8) {
		errs = append(errs, ErrNumberRequired("Int8", "required"))
	}
	if n.Int8Pointer == nil {
		errs = append(errs, ErrNumberRequired("Int8Pointer", "required"))
	}
	if IsInt16(n.Int16) {
		errs = append(errs, ErrNumberRequired("Int16", "required"))
	}
	if n.Int16Pointer == nil {
		errs = append(errs, ErrNumberRequired("Int16Pointer", "required"))
	}
	if IsInt32(n.Int32) {
		errs = append(errs, ErrNumberRequired("Int32", "required"))
	}
	if n.Int32Pointer == nil {
		errs = append(errs, ErrNumberRequired("Int32Pointer", "required"))
	}
	if IsInt64(n.Int64) {
		errs = append(errs, ErrNumberRequired("Int64", "required"))
	}
	if n.Int64Pointer == nil {
		errs = append(errs, ErrNumberRequired("Int64Pointer", "required"))
	}
	if IsFloat32(n.Float32) {
		errs = append(errs, ErrNumberRequired("Float32", "required"))
	}
	if n.Float32Pointer == nil {
		errs = append(errs, ErrNumberRequired("Float32Pointer", "required"))
	}
	if IsFloat64(n.Float64) {
		errs = append(errs, ErrNumberRequired("Float64", "required"))
	}
	if n.Float64Pointer == nil {
		errs = append(errs, ErrNumberRequired("Float64Pointer", "required"))
	}
	if IsComplex64(n.Complex64) {
		errs = append(errs, ErrNumberRequired("Complex64", "required"))
	}
	if n.Complex64Pointer == nil {
		errs = append(errs, ErrNumberRequired("Complex64Pointer", "required"))
	}
	if IsComplex128(n.Complex128) {
		errs = append(errs, ErrNumberRequired("Complex128", "required"))
	}
	if n.Complex128Pointer == nil {
		errs = append(errs, ErrNumberRequired("Complex128Pointer", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceString) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceString)) {
		errs = append(errs, ErrEmpty("SliceString", "required"))
	}
	if t.PointerSliceString == nil {
		errs = append(errs, ErrEmpty("PointerSliceString", "required"))
	}
	if Empty(len(t.SlicePointerString)) {
		errs = append(errs, ErrEmpty("SlicePointerString", "required"))
	}
	if t.PointerSlicePointerString == nil {
		errs = append(errs, ErrEmpty("PointerSlicePointerString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint)) {
		errs = append(errs, ErrEmpty("SliceUint", "required"))
	}
	if t.PointerSliceUint == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceUint", "required"))
	}
	if Empty(len(t.SlicePointerUint)) {
		errs = append(errs, ErrEmpty("SlicePointerUint", "required"))
	}
	if t.PointerSlicePointerUint == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerUint", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint8) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint8)) {
		errs = append(errs, ErrEmpty("SliceUint8", "required"))
	}
	if t.PointerSliceUint8 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceUint8", "required"))
	}
	if Empty(len(t.SlicePointerUint8)) {
		errs = append(errs, ErrEmpty("SlicePointerUint8", "required"))
	}
	if t.PointerSlicePointerUint8 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerUint8", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint16) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint16)) {
		errs = append(errs, ErrEmpty("SliceUint16", "required"))
	}
	if t.PointerSliceUint16 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceUint16", "required"))
	}
	if Empty(len(t.SlicePointerUint16)) {
		errs = append(errs, ErrEmpty("SlicePointerUint16", "required"))
	}
	if t.PointerSlicePointerUint16 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerUint16", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint32) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint32)) {
		errs = append(errs, ErrEmpty("SliceUint32", "required"))
	}
	if t.PointerSliceUint32 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceUint32", "required"))
	}
	if Empty(len(t.SlicePointerUint32)) {
		errs = append(errs, ErrEmpty("SlicePointerUint32", "required"))
	}
	if t.PointerSlicePointerUint32 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerUint32", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint64)) {
		errs = append(errs, ErrEmpty("SliceUint64", "required"))
	}
	if t.PointerSliceUint64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceUint64", "required"))
	}
	if Empty(len(t.SlicePointerUint64)) {
		errs = append(errs, ErrEmpty("SlicePointerUint64", "required"))
	}
	if t.PointerSlicePointerUint64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerUint64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUintptr) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUintptr)) {
		errs = append(errs, ErrEmpty("SliceUintptr", "required"))
	}
	if t.PointerSliceUintptr == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceUintptr", "required"))
	}
	if Empty(len(t.SlicePointerUintptr)) {
		errs = append(errs, ErrEmpty("SlicePointerUintptr", "required"))
	}
	if t.PointerSlicePointerUintptr == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerUintptr", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt)) {
		errs = append(errs, ErrEmpty("SliceInt", "required"))
	}
	if t.PointerSliceInt == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceInt", "required"))
	}
	if Empty(len(t.SlicePointerInt)) {
		errs = append(errs, ErrEmpty("SlicePointerInt", "required"))
	}
	if t.PointerSlicePointerInt == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerInt", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt8) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt8)) {
		errs = append(errs, ErrEmpty("SliceInt8", "required"))
	}
	if t.PointerSliceInt8 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceInt8", "required"))
	}
	if Empty(len(t.SlicePointerInt8)) {
		errs = append(errs, ErrEmpty("SlicePointerInt8", "required"))
	}
	if t.PointerSlicePointerInt8 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerInt8", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt16) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt16)) {
		errs = append(errs, ErrEmpty("SliceInt16", "required"))
	}
	if t.PointerSliceInt16 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceInt16", "required"))
	}
	if Empty(len(t.SlicePointerInt16)) {
		errs = append(errs, ErrEmpty("SlicePointerInt16", "required"))
	}
	if t.PointerSlicePointerInt16 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerInt16", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt32) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt32)) {
		errs = append(errs, ErrEmpty("SliceInt32", "required"))
	}
	if t.PointerSliceInt32 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceInt32", "required"))
	}
	if Empty(len(t.SlicePointerInt32)) {
		errs = append(errs, ErrEmpty("SlicePointerInt32", "required"))
	}
	if t.PointerSlicePointerInt32 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerInt32", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt64)) {
		errs = append(errs, ErrEmpty("SliceInt64", "required"))
	}
	if t.PointerSliceInt64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceInt64", "required"))
	}
	if Empty(len(t.SlicePointerInt64)) {
		errs = append(errs, ErrEmpty("SlicePointerInt64", "required"))
	}
	if t.PointerSlicePointerInt64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerInt64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceFloat32) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceFloat32)) {
		errs = append(errs, ErrEmpty("SliceFloat32", "required"))
	}
	if t.PointerSliceFloat32 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceFloat32", "required"))
	}
	if Empty(len(t.SlicePointerFloat32)) {
		errs = append(errs, ErrEmpty("SlicePointerFloat32", "required"))
	}
	if t.PointerSlicePointerFloat32 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerFloat32", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceFloat64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceFloat64)) {
		errs = append(errs, ErrEmpty("SliceFloat64", "required"))
	}
	if t.PointerSliceFloat64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceFloat64", "required"))
	}
	if Empty(len(t.SlicePointerFloat64)) {
		errs = append(errs, ErrEmpty("SlicePointerFloat64", "required"))
	}
	if t.PointerSlicePointerFloat64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerFloat64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceComplex64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceComplex64)) {
		errs = append(errs, ErrEmpty("SliceComplex64", "required"))
	}
	if t.PointerSliceComplex64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceComplex64", "required"))
	}
	if Empty(len(t.SlicePointerComplex64)) {
		errs = append(errs, ErrEmpty("SlicePointerComplex64", "required"))
	}
	if t.PointerSlicePointerComplex64 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerComplex64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceComplex128) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceComplex128)) {
		errs = append(errs, ErrEmpty("SliceComplex128", "required"))
	}
	if t.PointerSliceComplex128 == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceComplex128", "required"))
	}
	if Empty(len(t.SlicePointerComplex128)) {
		errs = append(errs, ErrEmpty("SlicePointerComplex128", "required"))
	}
	if t.PointerSlicePointerComplex128 == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerComplex128", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeString) Validate() (errs ValidationErrors) {
	if Empty(len(t.String)) {
		errs = append(errs, ErrEmpty("String", "required"))
	}
	if t.Pointer == nil {
		errs = append(errs, ErrEmpty("Pointer", "required"))
	}
	return errs
}

// Validate TODO
func (s *Struct) Validate() (errs ValidationErrors) {
	if IsValid(s.Struct) {
		errs = append(errs, ErrEmpty("Struct", "required"))
	}
	if s.StructPointer == nil {
		errs = append(errs, ErrEmpty("StructPointer", "required"))
	}
	if Empty(len(s.SliceStruct)) {
		errs = append(errs, ErrEmpty("SliceStruct", "required"))
	}
	if Empty(len(s.SliceStructPointer)) {
		errs = append(errs, ErrEmpty("SliceStructPointer", "required"))
	}
	if s.SlicePointerStruct == nil {
		errs = append(errs, ErrEmpty("SlicePointerStruct", "required"))
	}
	if s.SlicePointerStructPointer == nil {
		errs = append(errs, ErrEmpty("SlicePointerStructPointer", "required"))
	}
	return errs
}
