package main

import "fmt"

func main() {
	fmt.Println("arries")

	var fruitList [4]string // important to explicity give size

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fmt.Println(fruitList)

	fmt.Println(fruitList[0])
	fmt.Println(len(fruitList))
	var veg = [3]string{"p", "b"}
	fmt.Printf("type of %T\n", veg)
	fmt.Println(veg)
}
