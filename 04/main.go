package main

import (
	"fmt"
	"time"
)

func main() {

	presentime := time.Now()
	fmt.Println(presentime)

	fmt.Println(presentime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentime)
}
