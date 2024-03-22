package main

import "fmt"

// when
// 1. when we need to update state
// pointer 8bytes

// 2. when we want to optimize the memory for Large objects that are getting called a lot.
type User struct {
	email    string
	username string
	age      int
	file     []byte // ?? Large or small
}

func getUser() (User, error) {
	return User{}, fmt.Errorf("foo")
}

// 8bytes
func (u *User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

// 1gb user size
// x amount of bytes => sizeOf(user)
func Email(user User) string {
	return user.email
}

func main() {
	user := User{
		email:    "test@test.com",
		username: "Henry",
		age:      22,
	}
	user.updateEmail("fool@g.aol")
	fmt.Println(user.Email())
}
