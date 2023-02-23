package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"

	fmt.Println("List of languages", languages)
	// List of languages map[JS:JavaScript RB:Ruby]
	fmt.Println(languages["JS"])
	//JavaScript

	//Delete some values
	delete(languages,"JS")
	fmt.Println("List of languages", languages)
	// List of languages map[RB:Ruby]
}