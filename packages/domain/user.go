package domain

type User struct {
	Name string
	Age  int
}

var (
	UserMatt User
	UserJohn User
)

func init() {
	UserMatt = User{Name: "Matt", Age: 33}
	UserJohn = User{"John", 42}
}
