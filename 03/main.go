package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	wel := " welcome"
	fmt.Println(wel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter an integer")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numRating, err := strconv.ParseFloat(input, 64)
	fmt.Println(input)
	fmt.Printf("type of input %T", numRating)
	if err != nil {

		fmt.Println(err)
	}

}
