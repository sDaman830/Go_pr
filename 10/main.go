package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("rand")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", diceNumberx)

	switch diceNumber {
	case 1:
		fmt.Println(diceNumber)
	default:
		fmt.Println("tell wgat it is")
	}
}
