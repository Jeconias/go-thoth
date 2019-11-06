package models

// Validate TODO
func (m *MapStringToInterface) Validate() (errs ValidationErrors) {
	if IsValid(&m.A) {
		errs = append(errs, ErrEmpty("A", "required"))
	}
	if IsValid(m.B) {
		errs = append(errs, ErrEmpty("B", "required"))
	}
	return errs
}

// Validate TODO
func (m *MapStringToString) Validate() (errs ValidationErrors) {
	if IsValid(&m.A) {
		errs = append(errs, ErrEmpty("A", "required"))
	}
	if IsValid(m.B) {
		errs = append(errs, ErrEmpty("B", "required"))
	}
	return errs
}

// Validate TODO
func (m *MapIntToString) Validate() (errs ValidationErrors) {
	if IsValid(&m.A) {
		errs = append(errs, ErrEmpty("A", "required"))
	}
	if IsValid(m.B) {
		errs = append(errs, ErrEmpty("B", "required"))
	}
	return errs
}

// Validate TODO
func (m *MapIntToBool) Validate() (errs ValidationErrors) {
	if IsValid(&m.A) {
		errs = append(errs, ErrEmpty("A", "required"))
	}
	if IsValid(m.B) {
		errs = append(errs, ErrEmpty("B", "required"))
	}
	return errs
}

// Validate TODO
func (m *MapIntToStruct) Validate() (errs ValidationErrors) {
	if IsValid(&m.A) {
		errs = append(errs, ErrEmpty("A", "required"))
	}
	if IsValid(&m.APointer) {
		errs = append(errs, ErrEmpty("APointer", "required"))
	}
	if IsValid(m.B) {
		errs = append(errs, ErrEmpty("B", "required"))
	}
	if IsValid(m.BPointer) {
		errs = append(errs, ErrEmpty("BPointer", "required"))
	}
	return errs
}

// Validate TODO
func (h *Home) Validate() (errs ValidationErrors) {
	if IsInt(&h.ID) {
		errs = append(errs, ErrNumberRequired("ID", "required"))
	}
	if IsValid(&h.Address) {
		errs = append(errs, ErrEmpty("Address", "required"))
	}
	return errs
}

// Validate TODO
func (c *Client) Validate() (errs ValidationErrors) {
	if IsInt64(&c.ID) {
		errs = append(errs, ErrNumberRequired("ID", "required"))
	}
	if Empty(len(c.Name)) {
		errs = append(errs, ErrEmpty("Name", "required"))
	}
	if c.LastName == nil {
		errs = append(errs, ErrEmpty("LastName", "required"))
	}
	if IsValid(c.Address) {
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
	if IsUint(&n.Uint) {
		errs = append(errs, ErrNumberRequired("Uint", "required"))
	}
	if n.UintPointer == nil {
		errs = append(errs, ErrNumberRequired("UintPointer", "required"))
	}
	if IsUint8(&n.Uint8) {
		errs = append(errs, ErrNumberRequired("Uint8", "required"))
	}
	if IsUint8(n.Uint8Pointer) {
		errs = append(errs, ErrNumberRequired("Uint8Pointer", "required"))
	}
	if IsUint16(&n.Uint16) {
		errs = append(errs, ErrNumberRequired("Uint16", "required"))
	}
	if IsUint16(n.Uint16Pointer) {
		errs = append(errs, ErrNumberRequired("Uint16Pointer", "required"))
	}
	if IsUint32(&n.Uint32) {
		errs = append(errs, ErrNumberRequired("Uint32", "required"))
	}
	if IsUint32(n.Uint32Pointer) {
		errs = append(errs, ErrNumberRequired("Uint32Pointer", "required"))
	}
	if IsUint64(&n.Uint64) {
		errs = append(errs, ErrNumberRequired("Uint64", "required"))
	}
	if IsUint64(n.Uint64Pointer) {
		errs = append(errs, ErrNumberRequired("Uint64Pointer", "required"))
	}
	if IsUintptr(&n.Uintptr) {
		errs = append(errs, ErrNumberRequired("Uintptr", "required"))
	}
	if IsUintptr(n.UintptrPointer) {
		errs = append(errs, ErrNumberRequired("UintptrPointer", "required"))
	}
	if IsInt(&n.Int) {
		errs = append(errs, ErrNumberRequired("Int", "required"))
	}
	if IsInt(n.IntPointer) {
		errs = append(errs, ErrNumberRequired("IntPointer", "required"))
	}
	if IsInt8(&n.Int8) {
		errs = append(errs, ErrNumberRequired("Int8", "required"))
	}
	if IsInt8(n.Int8Pointer) {
		errs = append(errs, ErrNumberRequired("Int8Pointer", "required"))
	}
	if IsInt16(&n.Int16) {
		errs = append(errs, ErrNumberRequired("Int16", "required"))
	}
	if IsInt16(n.Int16Pointer) {
		errs = append(errs, ErrNumberRequired("Int16Pointer", "required"))
	}
	if IsInt32(&n.Int32) {
		errs = append(errs, ErrNumberRequired("Int32", "required"))
	}
	if IsInt32(n.Int32Pointer) {
		errs = append(errs, ErrNumberRequired("Int32Pointer", "required"))
	}
	if IsInt64(&n.Int64) {
		errs = append(errs, ErrNumberRequired("Int64", "required"))
	}
	if IsInt64(n.Int64Pointer) {
		errs = append(errs, ErrNumberRequired("Int64Pointer", "required"))
	}
	if IsFloat32(&n.Float32) {
		errs = append(errs, ErrNumberRequired("Float32", "required"))
	}
	if IsFloat32(n.Float32Pointer) {
		errs = append(errs, ErrNumberRequired("Float32Pointer", "required"))
	}
	if IsFloat64(&n.Float64) {
		errs = append(errs, ErrNumberRequired("Float64", "required"))
	}
	if IsFloat64(n.Float64Pointer) {
		errs = append(errs, ErrNumberRequired("Float64Pointer", "required"))
	}
	if IsComplex64(&n.Complex64) {
		errs = append(errs, ErrNumberRequired("Complex64", "required"))
	}
	if IsComplex64(n.Complex64Pointer) {
		errs = append(errs, ErrNumberRequired("Complex64Pointer", "required"))
	}
	if IsComplex128(&n.Complex128) {
		errs = append(errs, ErrNumberRequired("Complex128", "required"))
	}
	if IsComplex128(n.Complex128Pointer) {
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
		errs = append(errs, ErrNumberRequired("SliceUint", "required"))
	}
	if t.PointerSliceUint == nil {
		errs = append(errs, ErrNumberRequired("PointerSliceUint", "required"))
	}
	if Empty(len(t.SlicePointerUint)) {
		errs = append(errs, ErrNumberRequired("SlicePointerUint", "required"))
	}
	if t.PointerSlicePointerUint == nil {
		errs = append(errs, ErrNumberRequired("PointerSlicePointerUint", "required"))
	}
	return errs
}

// Validate TODO
func (s *SliceInt64) Validate() (errs ValidationErrors) {
	if Empty(len(s.A)) {
		errs = append(errs, ErrNumberRequired("A", "required"))
	}
	if Empty(len(s.C)) {
		errs = append(errs, ErrNumberRequired("C", "required"))
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
	if IsValid(&s.Struct) {
		errs = append(errs, ErrEmpty("Struct", "required"))
	}
	if IsValid(s.StructPointer) {
		errs = append(errs, ErrEmpty("StructPointer", "required"))
	}
	if Empty(len(s.SliceStruct)) {
		errs = append(errs, ErrEmpty("SliceStruct", "required"))
	}
	if Empty(len(s.SliceStructPointer)) {
		errs = append(errs, ErrEmpty("SliceStructPointer", "required"))
	}
	if IsValid(s.SlicePointerStruct) {
		errs = append(errs, ErrEmpty("SlicePointerStruct", "required"))
	}
	if IsValid(s.SlicePointerStructPointer) {
		errs = append(errs, ErrEmpty("SlicePointerStructPointer", "required"))
	}
	return errs
}
