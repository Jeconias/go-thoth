package models

// UUIDValidate TODO
type UUIDValidate struct {
	UUID    string   `thoth:"uuid"`
	Pointer *string  `thoth:"uuid"`
	Slice   []string `thoth:"uuid"`
}

// UUID3Validate TODO
type UUID3Validate struct {
	UUID3   string   `thoth:"uuid3"`
	Pointer *string  `thoth:"uuid3"`
	Slice   []string `thoth:"uuid3"`
}

// UUID4Validate TODO
type UUID4Validate struct {
	UUID4   string   `thoth:"uuid4"`
	Pointer *string  `thoth:"uuid4"`
	Slice   []string `thoth:"uuid4"`
}

// UUID5Validate TODO
type UUID5Validate struct {
	UUID5   string   `thoth:"uuid5"`
	Pointer *string  `thoth:"uuid5"`
	Slice   []string `thoth:"uuid5"`
}

// UUIDrfc4122Validate TODO
type UUIDrfc4122Validate struct {
	UUID    string   `thoth:"uuid_rfc4122"`
	Pointer *string  `thoth:"uuid_rfc4122"`
	Slice   []string `thoth:"uuid_rfc4122"`
}

// UUID3rfc4122Validate TODO
type UUID3rfc4122Validate struct {
	UUID3rfc4122 string   `thoth:"uuid3_rfc4122"`
	Pointer      *string  `thoth:"uuid3_rfc4122"`
	Slice        []string `thoth:"uuid3_rfc4122"`
}

// UUID4rfc4122Validate TODO
type UUID4rfc4122Validate struct {
	UUID4rfc4122 string   `thoth:"uuid4_rfc4122"`
	Pointer      *string  `thoth:"uuid4_rfc4122"`
	Slice        []string `thoth:"uuid4_rfc4122"`
}

// UUID5rfc4122Validate TODO
type UUID5rfc4122Validate struct {
	UUID5rfc4122 string   `thoth:"uuid5_rfc4122"`
	Pointer      *string  `thoth:"uuid5_rfc4122"`
	Slice        []string `thoth:"uuid5_rfc4122"`
}
