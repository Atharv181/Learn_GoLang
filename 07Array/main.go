package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in GO")
	// Not used much in language instead slices used

	var names [3]string

	names[0] = "Atharv"
	names[1] = "King"

	fmt.Println("Value of array is",names)
	//OP 
	//Value of array is [Atharv King ]

	fmt.Println("Lenght of array is",len(names))
	//Lenght of array is 3

	var surname = [4]string{"Bobade" , "Hello" , "Queen", "OK"}
	fmt.Println(surname)
	fmt.Println("Lenght of array is",len(surname))
	//[Bobade Hello Queen OK]
	 //Lenght of array is 4

}