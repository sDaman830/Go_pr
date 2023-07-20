package main

import "fmt"

func main() {
	var username string = "daman"
	fmt.Println(username)
	fmt.Printf("variable is of: %T\n", username)

	var another int
	fmt.Println(another)
	fmt.Printf("it is of type: %T\n", another)

	daman := 3000
	fmt.Println(daman)
}
