package main

import "fmt"

func main() {
	fmt.Println("maps")
	languages := make(map[string]string)

	languages["JS"] = "Js"
	languages["p"] = "py"
	languages["Ja"] = "Java"

	fmt.Println("List of all languages", languages)

	delete(languages, "JS")
}
