package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world") //puts it at the end of code, stores stuff into stack and works like last in first out then at the end it empties te stack
	fmt.Println("Hello")
}
