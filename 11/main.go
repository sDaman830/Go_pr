package main

import "fmt"

func main() {
	// no inheritance in go

	daman := User{"daman", "sdaman830@gmail.com ", true, 20}
	fmt.Println(daman)
	daman.GetStatus()
	daman.GetEmail()
	fmt.Println("%v", daman.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active", u.Status)
}

func (u User) GetEmail() {
	u.Email = "daman@mail"
	fmt.Println("%v", u.Email)
}
