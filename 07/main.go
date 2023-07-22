package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("sllices")
	var fruitList = []string{}

	fmt.Printf("type of fruitlist %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 224
	highScores[2] = 334
	highScores[3] = 134

	highScores = append(highScores, 555, 666, 848)
	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"r", "d", "wew", "rr"}
	fmt.Println(courses)
}
