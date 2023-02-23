package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else statements")

	Number := 10
	var result string

	if Number > 10 {
		result = "Number is greater than 10"
	} else if Number < 10{
		result = "Number is less than 10"
	} else {
		result = "Value is exactly 10"
	}

	fmt.Println(result)
	// Value is exactly 10



	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else {
		fmt.Println("Number is odd")
	}
	// Number is odd



	if num := 1; num <3 {
		fmt.Println("Num is less than 3")
	}else {
		fmt.Println("Num is greater than 3")
	}
	// Num is less than 3
}