package main

import "fmt"

func main() {
	// no inheritance in go

	daman := User{"daman", "sdaman830@gmail.com ", true, 20}
	fmt.Println(daman)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
