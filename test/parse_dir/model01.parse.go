package any

type Home struct {
	ID      int64  `thoth:"required"`
	Address string `thoth:"required"`
}

type Address struct {
	ID     int64
	Street string
}

type Client struct {
	ID   int64  `thoth:"required"`
	Name string `thoth:"required"`
}
