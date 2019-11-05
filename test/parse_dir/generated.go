package any

// Validate TODO
func (m *MapStringInterface) Validate() (errs ValidationErrors) {
	if IsValid(&m.A) {
		errs = append(errs, ErrNumberRequired("A", "required"))
	} // TODO: To extract map
	if IsValid(m.B) {
		errs = append(errs, ErrNumberRequired("B", "required"))
	} // TODO: To extract map
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
	if Empty(len(*c.LastName)) {
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
	if IsUint(n.UintPointer) {
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
func (s *SliceString) Validate() (errs ValidationErrors) {
	if Empty(len(s.A)) {
		errs = append(errs, ErrEmpty("A", "required"))
	}
	if Empty(len(*s.B)) {
		errs = append(errs, ErrEmpty("B", "required"))
	}
	if Empty(len(s.C)) {
		errs = append(errs, ErrEmpty("C", "required"))
	}
	if Empty(len(*s.D)) {
		errs = append(errs, ErrEmpty("D", "required"))
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
func (s *Struct) Validate() (errs ValidationErrors) {
	if IsValid(&s.A) {
		errs = append(errs, ErrEmpty("A", "required"))
	}
	if IsValid(s.B) {
		errs = append(errs, ErrEmpty("B", "required"))
	}
	if Empty(len(s.C)) {
		errs = append(errs, ErrEmpty("C", "required"))
	}
	if Empty(len(s.D)) {
		errs = append(errs, ErrEmpty("D", "required"))
	}
	if IsValid(s.E) {
		errs = append(errs, ErrEmpty("E", "required"))
	}
	if IsValid(s.F) {
		errs = append(errs, ErrEmpty("F", "required"))
	}
	return errs
}
