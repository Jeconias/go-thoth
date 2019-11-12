package models

// Validate TODO
func (t *TypeChanString) Validate() (errs ValidationErrors) {
	if Empty(len(t.ChanString)) {
		errs = append(errs, NewError("ChanString", "required"))
	}

	if Empty(len(t.PointerChanString)) {
		errs = append(errs, NewError("PointerChanString", "required"))
	}

	if Empty(len(t.ChanSliceString)) {
		errs = append(errs, NewError("ChanSliceString", "required"))
	}

	if Empty(len(t.ChanSlicePointerString)) {
		errs = append(errs, NewError("ChanSlicePointerString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeChanUint) Validate() (errs ValidationErrors) {
	if Empty(len(t.ChanUint)) {
		errs = append(errs, NewError("ChanUint", "required"))
	}

	if Empty(len(t.PointerChanUint)) {
		errs = append(errs, NewError("PointerChanUint", "required"))
	}

	if Empty(len(t.ChanSliceUint)) {
		errs = append(errs, NewError("ChanSliceUint", "required"))
	}

	if Empty(len(t.ChanSlicePointerUint)) {
		errs = append(errs, NewError("ChanSlicePointerUint", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeInterface) Validate() (errs ValidationErrors) {
	if IsValid(t.Interface) {
		errs = append(errs, NewError("Interface", "required"))
	}

	if t.PointerInterface == nil {
		errs = append(errs, NewError("PointerInterface", "required"))
	}
	return errs
}

// Validate TODO
func (l *LenField) Validate() (errs ValidationErrors) {
	if !l.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if len(l.Name) == 10 {
		errs = append(errs, NewError("Name", "len"))
	}
	return errs
}

// Validate TODO
func (l *LenFieldStrPointer) Validate() (errs ValidationErrors) {
	if !l.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if l.Name == nil {
		errs = append(errs, NewError("Name", "len"))
	}
	return errs
}

// Validate TODO
func (l *LenFields) Validate() (errs ValidationErrors) {
	if l.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !l.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if len(l.Names) == 2 {
		errs = append(errs, NewError("Names", "len"))
	}

	if l.Subjects == nil {
		errs = append(errs, NewError("Subjects", "len"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapStringToInterface) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapStringToInterface)) {
		errs = append(errs, NewError("MapStringToInterface", "required"))
	}

	if t.PointerMapStringToInterface == nil {
		errs = append(errs, NewError("PointerMapStringToInterface", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapStringToString) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapStringToString)) {
		errs = append(errs, NewError("MapStringToString", "required"))
	}

	if t.PointerMapStringToString == nil {
		errs = append(errs, NewError("PointerMapStringToString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapIntToString) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapIntToString)) {
		errs = append(errs, NewError("MapIntToString", "required"))
	}

	if t.PointerMapIntToString == nil {
		errs = append(errs, NewError("PointerMapIntToString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapIntToBool) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapIntToBool)) {
		errs = append(errs, NewError("MapIntToBool", "required"))
	}

	if t.PointerMapIntToBool == nil {
		errs = append(errs, NewError("PointerMapIntToBool", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeMapIntToStruct) Validate() (errs ValidationErrors) {
	if Empty(len(t.MapIntToStructA)) {
		errs = append(errs, NewError("MapIntToStructA", "required"))
	}

	if Empty(len(t.PointerMapIntToStructA)) {
		errs = append(errs, NewError("PointerMapIntToStructA", "required"))
	}

	if t.MapIntToStructB == nil {
		errs = append(errs, NewError("MapIntToStructB", "required"))
	}

	if t.PointerMapIntToStructPointerB == nil {
		errs = append(errs, NewError("PointerMapIntToStructPointerB", "required"))
	}
	return errs
}

// Validate TODO
func (h *Home) Validate() (errs ValidationErrors) {
	if IsInt(h.ID) {
		errs = append(errs, NewError("ID", "required"))
	}
	if (h.Address == Address{}) {
		errs = append(errs, NewError("Address", "required"))
	}
	return errs
}

// Validate TODO
func (c *Client) Validate() (errs ValidationErrors) {
	if IsInt64(c.ID) {
		errs = append(errs, NewError("ID", "required"))
	}

	if Empty(len(c.Name)) {
		errs = append(errs, NewError("Name", "required"))
	}

	if c.LastName == nil {
		errs = append(errs, NewError("LastName", "required"))
	}

	if c.Address == nil {
		errs = append(errs, NewError("Address", "required"))
	}
	return errs
}

// Validate TODO
func (u *User) Validate() (errs ValidationErrors) {
	if Empty(len(u.Name)) {
		errs = append(errs, NewError("Name", "required"))
	}

	return errs
}

// Validate TODO
func (t *TypeEqNumber) Validate() (errs ValidationErrors) {
	if t.Uint != 1 {
		errs = append(errs, NewError("Uint", "eq"))
	}
	if t.UintPointer == nil || *t.UintPointer != 2 {
		errs = append(errs, NewError("UintPointer", "eq"))
	}
	if t.Uint8 != 3 {
		errs = append(errs, NewError("Uint8", "eq"))
	}
	if t.Uint8Pointer == nil || *t.Uint8Pointer != 4 {
		errs = append(errs, NewError("Uint8Pointer", "eq"))
	}
	if t.Uint16 != 5 {
		errs = append(errs, NewError("Uint16", "eq"))
	}
	if t.Uint16Pointer == nil || *t.Uint16Pointer != 6 {
		errs = append(errs, NewError("Uint16Pointer", "eq"))
	}
	if t.Uint32 != 7 {
		errs = append(errs, NewError("Uint32", "eq"))
	}
	if t.Uint32Pointer == nil || *t.Uint32Pointer != 8 {
		errs = append(errs, NewError("Uint32Pointer", "eq"))
	}
	if t.Uint64 != 9 {
		errs = append(errs, NewError("Uint64", "eq"))
	}
	if t.Uint64Pointer == nil || *t.Uint64Pointer != 10 {
		errs = append(errs, NewError("Uint64Pointer", "eq"))
	}
	if t.Uintptr != 11 {
		errs = append(errs, NewError("Uintptr", "eq"))
	}
	if t.UintptrPointer == nil || *t.UintptrPointer != 12 {
		errs = append(errs, NewError("UintptrPointer", "eq"))
	}
	if t.Int != 13 {
		errs = append(errs, NewError("Int", "eq"))
	}
	if t.IntPointer == nil || *t.IntPointer != 14 {
		errs = append(errs, NewError("IntPointer", "eq"))
	}
	if t.Int8 != 15 {
		errs = append(errs, NewError("Int8", "eq"))
	}
	if t.Int8Pointer == nil || *t.Int8Pointer != 16 {
		errs = append(errs, NewError("Int8Pointer", "eq"))
	}
	if t.Int16 != 17 {
		errs = append(errs, NewError("Int16", "eq"))
	}
	if t.Int16Pointer == nil || *t.Int16Pointer != 18 {
		errs = append(errs, NewError("Int16Pointer", "eq"))
	}
	if t.Int32 != 19 {
		errs = append(errs, NewError("Int32", "eq"))
	}
	if t.Int32Pointer == nil || *t.Int32Pointer != 20 {
		errs = append(errs, NewError("Int32Pointer", "eq"))
	}
	if t.Int64 != 21 {
		errs = append(errs, NewError("Int64", "eq"))
	}
	if t.Int64Pointer == nil || *t.Int64Pointer != 22 {
		errs = append(errs, NewError("Int64Pointer", "eq"))
	}
	if t.Float32 != 23 {
		errs = append(errs, NewError("Float32", "eq"))
	}
	if t.Float32Pointer == nil || *t.Float32Pointer != 24 {
		errs = append(errs, NewError("Float32Pointer", "eq"))
	}
	if t.Float64 != 25 {
		errs = append(errs, NewError("Float64", "eq"))
	}
	if t.Float64Pointer == nil || *t.Float64Pointer != 26 {
		errs = append(errs, NewError("Float64Pointer", "eq"))
	}
	if t.Complex64 != 27 {
		errs = append(errs, NewError("Complex64", "eq"))
	}
	if t.Complex64Pointer == nil || *t.Complex64Pointer != 28 {
		errs = append(errs, NewError("Complex64Pointer", "eq"))
	}
	if t.Complex128 != 29 {
		errs = append(errs, NewError("Complex128", "eq"))
	}
	if t.Complex128Pointer == nil || *t.Complex128Pointer != 30 {
		errs = append(errs, NewError("Complex128Pointer", "eq"))
	}
	return errs
}

// Validate TODO
func (n *Number) Validate() (errs ValidationErrors) {
	if IsUint(n.Uint) {
		errs = append(errs, NewError("Uint", "required"))
	}

	if n.UintPointer == nil {
		errs = append(errs, NewError("UintPointer", "required"))
	}

	if IsUint8(n.Uint8) {
		errs = append(errs, NewError("Uint8", "required"))
	}

	if n.Uint8Pointer == nil {
		errs = append(errs, NewError("Uint8Pointer", "required"))
	}

	if IsUint16(n.Uint16) {
		errs = append(errs, NewError("Uint16", "required"))
	}

	if n.Uint16Pointer == nil {
		errs = append(errs, NewError("Uint16Pointer", "required"))
	}

	if IsUint32(n.Uint32) {
		errs = append(errs, NewError("Uint32", "required"))
	}

	if n.Uint32Pointer == nil {
		errs = append(errs, NewError("Uint32Pointer", "required"))
	}

	if IsUint64(n.Uint64) {
		errs = append(errs, NewError("Uint64", "required"))
	}

	if n.Uint64Pointer == nil {
		errs = append(errs, NewError("Uint64Pointer", "required"))
	}

	if IsUintptr(n.Uintptr) {
		errs = append(errs, NewError("Uintptr", "required"))
	}

	if n.UintptrPointer == nil {
		errs = append(errs, NewError("UintptrPointer", "required"))
	}

	if IsInt(n.Int) {
		errs = append(errs, NewError("Int", "required"))
	}

	if n.IntPointer == nil {
		errs = append(errs, NewError("IntPointer", "required"))
	}

	if IsInt8(n.Int8) {
		errs = append(errs, NewError("Int8", "required"))
	}

	if n.Int8Pointer == nil {
		errs = append(errs, NewError("Int8Pointer", "required"))
	}

	if IsInt16(n.Int16) {
		errs = append(errs, NewError("Int16", "required"))
	}

	if n.Int16Pointer == nil {
		errs = append(errs, NewError("Int16Pointer", "required"))
	}

	if IsInt32(n.Int32) {
		errs = append(errs, NewError("Int32", "required"))
	}

	if n.Int32Pointer == nil {
		errs = append(errs, NewError("Int32Pointer", "required"))
	}

	if IsInt64(n.Int64) {
		errs = append(errs, NewError("Int64", "required"))
	}

	if n.Int64Pointer == nil {
		errs = append(errs, NewError("Int64Pointer", "required"))
	}

	if IsFloat32(n.Float32) {
		errs = append(errs, NewError("Float32", "required"))
	}

	if n.Float32Pointer == nil {
		errs = append(errs, NewError("Float32Pointer", "required"))
	}

	if IsFloat64(n.Float64) {
		errs = append(errs, NewError("Float64", "required"))
	}

	if n.Float64Pointer == nil {
		errs = append(errs, NewError("Float64Pointer", "required"))
	}

	if IsComplex64(n.Complex64) {
		errs = append(errs, NewError("Complex64", "required"))
	}

	if n.Complex64Pointer == nil {
		errs = append(errs, NewError("Complex64Pointer", "required"))
	}

	if IsComplex128(n.Complex128) {
		errs = append(errs, NewError("Complex128", "required"))
	}

	if n.Complex128Pointer == nil {
		errs = append(errs, NewError("Complex128Pointer", "required"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if Empty(len(r.Name)) && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_with"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_with"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithFields) Validate() (errs ValidationErrors) {
	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.ID == 1 || r.Status != false) {
		errs = append(errs, NewError("Name", "required_with"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithAllField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if Empty(len(r.Name)) && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_with_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithAllFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_with_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithAllFields) Validate() (errs ValidationErrors) {
	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.ID == 1 && r.Status != false) {
		errs = append(errs, NewError("Name", "required_with_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if Empty(len(r.Name)) && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_without"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_without"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutFields) Validate() (errs ValidationErrors) {
	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.ID == 1 || r.Status != false) {
		errs = append(errs, NewError("Name", "required_without"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutAllField) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if Empty(len(r.Name)) && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_without_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutAllFieldStrPointer) Validate() (errs ValidationErrors) {
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.Status != false) {
		errs = append(errs, NewError("Name", "required_without_all"))
	}
	return errs
}

// Validate TODO
func (r *RequiredWithoutAllFields) Validate() (errs ValidationErrors) {
	if r.ID != 1 {
		errs = append(errs, NewError("ID", "eq"))
	}
	if !r.Status {
		errs = append(errs, NewError("Status", "required"))
	}

	if r.Name == nil && (r.ID == 1 && r.Status != false) {
		errs = append(errs, NewError("Name", "required_without_all"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceString) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceString)) {
		errs = append(errs, NewError("SliceString", "required"))
	}

	if t.PointerSliceString == nil {
		errs = append(errs, NewError("PointerSliceString", "required"))
	}

	if Empty(len(t.SlicePointerString)) {
		errs = append(errs, NewError("SlicePointerString", "required"))
	}

	if t.PointerSlicePointerString == nil {
		errs = append(errs, NewError("PointerSlicePointerString", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint)) {
		errs = append(errs, NewError("SliceUint", "required"))
	}

	if t.PointerSliceUint == nil {
		errs = append(errs, NewError("PointerSliceUint", "required"))
	}

	if Empty(len(t.SlicePointerUint)) {
		errs = append(errs, NewError("SlicePointerUint", "required"))
	}

	if t.PointerSlicePointerUint == nil {
		errs = append(errs, NewError("PointerSlicePointerUint", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint8) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint8)) {
		errs = append(errs, NewError("SliceUint8", "required"))
	}

	if t.PointerSliceUint8 == nil {
		errs = append(errs, NewError("PointerSliceUint8", "required"))
	}

	if Empty(len(t.SlicePointerUint8)) {
		errs = append(errs, NewError("SlicePointerUint8", "required"))
	}

	if t.PointerSlicePointerUint8 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint8", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint16) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint16)) {
		errs = append(errs, NewError("SliceUint16", "required"))
	}

	if t.PointerSliceUint16 == nil {
		errs = append(errs, NewError("PointerSliceUint16", "required"))
	}

	if Empty(len(t.SlicePointerUint16)) {
		errs = append(errs, NewError("SlicePointerUint16", "required"))
	}

	if t.PointerSlicePointerUint16 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint16", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint32) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint32)) {
		errs = append(errs, NewError("SliceUint32", "required"))
	}

	if t.PointerSliceUint32 == nil {
		errs = append(errs, NewError("PointerSliceUint32", "required"))
	}

	if Empty(len(t.SlicePointerUint32)) {
		errs = append(errs, NewError("SlicePointerUint32", "required"))
	}

	if t.PointerSlicePointerUint32 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint32", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUint64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUint64)) {
		errs = append(errs, NewError("SliceUint64", "required"))
	}

	if t.PointerSliceUint64 == nil {
		errs = append(errs, NewError("PointerSliceUint64", "required"))
	}

	if Empty(len(t.SlicePointerUint64)) {
		errs = append(errs, NewError("SlicePointerUint64", "required"))
	}

	if t.PointerSlicePointerUint64 == nil {
		errs = append(errs, NewError("PointerSlicePointerUint64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceUintptr) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceUintptr)) {
		errs = append(errs, NewError("SliceUintptr", "required"))
	}

	if t.PointerSliceUintptr == nil {
		errs = append(errs, NewError("PointerSliceUintptr", "required"))
	}

	if Empty(len(t.SlicePointerUintptr)) {
		errs = append(errs, NewError("SlicePointerUintptr", "required"))
	}

	if t.PointerSlicePointerUintptr == nil {
		errs = append(errs, NewError("PointerSlicePointerUintptr", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt)) {
		errs = append(errs, NewError("SliceInt", "required"))
	}

	if t.PointerSliceInt == nil {
		errs = append(errs, NewError("PointerSliceInt", "required"))
	}

	if Empty(len(t.SlicePointerInt)) {
		errs = append(errs, NewError("SlicePointerInt", "required"))
	}

	if t.PointerSlicePointerInt == nil {
		errs = append(errs, NewError("PointerSlicePointerInt", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt8) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt8)) {
		errs = append(errs, NewError("SliceInt8", "required"))
	}

	if t.PointerSliceInt8 == nil {
		errs = append(errs, NewError("PointerSliceInt8", "required"))
	}

	if Empty(len(t.SlicePointerInt8)) {
		errs = append(errs, NewError("SlicePointerInt8", "required"))
	}

	if t.PointerSlicePointerInt8 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt8", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt16) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt16)) {
		errs = append(errs, NewError("SliceInt16", "required"))
	}

	if t.PointerSliceInt16 == nil {
		errs = append(errs, NewError("PointerSliceInt16", "required"))
	}

	if Empty(len(t.SlicePointerInt16)) {
		errs = append(errs, NewError("SlicePointerInt16", "required"))
	}

	if t.PointerSlicePointerInt16 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt16", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt32) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt32)) {
		errs = append(errs, NewError("SliceInt32", "required"))
	}

	if t.PointerSliceInt32 == nil {
		errs = append(errs, NewError("PointerSliceInt32", "required"))
	}

	if Empty(len(t.SlicePointerInt32)) {
		errs = append(errs, NewError("SlicePointerInt32", "required"))
	}

	if t.PointerSlicePointerInt32 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt32", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceInt64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceInt64)) {
		errs = append(errs, NewError("SliceInt64", "required"))
	}

	if t.PointerSliceInt64 == nil {
		errs = append(errs, NewError("PointerSliceInt64", "required"))
	}

	if Empty(len(t.SlicePointerInt64)) {
		errs = append(errs, NewError("SlicePointerInt64", "required"))
	}

	if t.PointerSlicePointerInt64 == nil {
		errs = append(errs, NewError("PointerSlicePointerInt64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceFloat32) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceFloat32)) {
		errs = append(errs, NewError("SliceFloat32", "required"))
	}

	if t.PointerSliceFloat32 == nil {
		errs = append(errs, NewError("PointerSliceFloat32", "required"))
	}

	if Empty(len(t.SlicePointerFloat32)) {
		errs = append(errs, NewError("SlicePointerFloat32", "required"))
	}

	if t.PointerSlicePointerFloat32 == nil {
		errs = append(errs, NewError("PointerSlicePointerFloat32", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceFloat64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceFloat64)) {
		errs = append(errs, NewError("SliceFloat64", "required"))
	}

	if t.PointerSliceFloat64 == nil {
		errs = append(errs, NewError("PointerSliceFloat64", "required"))
	}

	if Empty(len(t.SlicePointerFloat64)) {
		errs = append(errs, NewError("SlicePointerFloat64", "required"))
	}

	if t.PointerSlicePointerFloat64 == nil {
		errs = append(errs, NewError("PointerSlicePointerFloat64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceComplex64) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceComplex64)) {
		errs = append(errs, NewError("SliceComplex64", "required"))
	}

	if t.PointerSliceComplex64 == nil {
		errs = append(errs, NewError("PointerSliceComplex64", "required"))
	}

	if Empty(len(t.SlicePointerComplex64)) {
		errs = append(errs, NewError("SlicePointerComplex64", "required"))
	}

	if t.PointerSlicePointerComplex64 == nil {
		errs = append(errs, NewError("PointerSlicePointerComplex64", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeSliceComplex128) Validate() (errs ValidationErrors) {
	if Empty(len(t.SliceComplex128)) {
		errs = append(errs, NewError("SliceComplex128", "required"))
	}

	if t.PointerSliceComplex128 == nil {
		errs = append(errs, NewError("PointerSliceComplex128", "required"))
	}

	if Empty(len(t.SlicePointerComplex128)) {
		errs = append(errs, NewError("SlicePointerComplex128", "required"))
	}

	if t.PointerSlicePointerComplex128 == nil {
		errs = append(errs, NewError("PointerSlicePointerComplex128", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeString) Validate() (errs ValidationErrors) {
	if Empty(len(t.String)) {
		errs = append(errs, NewError("String", "required"))
	}

	if t.Pointer == nil {
		errs = append(errs, NewError("Pointer", "required"))
	}
	return errs
}

// Validate TODO
func (t *TypeEqString) Validate() (errs ValidationErrors) {
	if t.String != "chico" {
		errs = append(errs, NewError("String", "eq"))
	}
	if t.Pointer == nil || *t.Pointer != "bento" {
		errs = append(errs, NewError("Pointer", "eq"))
	}
	return errs
}

// Validate TODO
func (s *Struct) Validate() (errs ValidationErrors) {
	if (s.Struct == StructA{}) {
		errs = append(errs, NewError("Struct", "required"))
	}

	if s.StructPointer == nil {
		errs = append(errs, NewError("StructPointer", "required"))
	}

	if Empty(len(s.SliceStruct)) {
		errs = append(errs, NewError("SliceStruct", "required"))
	}

	if Empty(len(s.SliceStructPointer)) {
		errs = append(errs, NewError("SliceStructPointer", "required"))
	}

	if s.SlicePointerStruct == nil {
		errs = append(errs, NewError("SlicePointerStruct", "required"))
	}

	if s.SlicePointerStructPointer == nil {
		errs = append(errs, NewError("SlicePointerStructPointer", "required"))
	}
	return errs
}
