package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is",ptr)
	//OP
	//Value of pointer is <nil>

	var myNumber = 45
	var ptr = &myNumber

	fmt.Println("Value of ptr is", ptr)
	fmt.Println("Value inside ptr is", *ptr)
	//OP
	// Value of ptr is 0xc00000e098
	// Value inside ptr is 45

	*ptr = *ptr + 2
	fmt.Println("New value of myNumber is",myNumber)
	//OP
	//New value of myNumber is 47
}