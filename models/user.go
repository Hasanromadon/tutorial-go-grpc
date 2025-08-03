package models

type User struct {
	Name string
	Age  int
}

// Method Receiver
func (u User) Greet() string {
	return "Halo, saya " + u.Name
}

// Interface
type Greeter interface {
	Greet() string
}
