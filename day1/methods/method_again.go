package main 

import "fmt"

// methods in go
type User struct {
	name string
	age int
}

// (u User) is the receiver of the method
func (user *User) updateName(newName string) {
	user.name = newName
}

func main() {
	user := User{name: "John", age: 25}
	fmt.Println(user)
	user.updateName("Doe")
	fmt.Println(user)
}