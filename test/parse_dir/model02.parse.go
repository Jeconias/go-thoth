package any

// User TODO
type User struct {
	ID   int64
	Name string `thoth:"required"`
	Age  int64
}
