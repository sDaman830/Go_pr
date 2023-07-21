package main

import "fmt"

func main() {
	fmt.Println("sllices")
	var fruitList = []string{}

	fmt.Printf("type of fruitlist %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
}
