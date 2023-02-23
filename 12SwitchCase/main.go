package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Start with Switch case in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value on dice is",diceNumber)
	
	switch diceNumber {
	case 1:
		fmt.Println("Value appeared on dice is 1")
	case 2:
		fmt.Println("Value appeared on dice is 2")
	case 3:
		fmt.Println("Value appeared on dice is 3")
	case 4:
		fmt.Println("Value appeared on dice is 4")
	case 5:
		fmt.Println("Value appeared on dice is 5")
		fallthrough //if case 5 hits then case 6 will also execute
	case 6:
		fmt.Println("Value appeared on dice is 6")
	default:
		fmt.Println("If not any case then this will execute")
	}
}
