package any

type User struct {
	ID   int64
	Name string `thoth:"name"`
	Age  int64
}
