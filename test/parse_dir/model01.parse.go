package any

// Home TODO
type Home struct {
	ID      int     `thoth:"required"`
	Address Address `thoth:"required"`
}

// Address TODO
type Address struct {
	ID     int64
	Street string
}

// Client TODO
type Client struct {
	ID       int64    `thoth:"required"`
	Name     string   `thoth:"required"`
	LastName *string  `thoth:"required"`
	Address  *Address `thoth:"required"`
}
