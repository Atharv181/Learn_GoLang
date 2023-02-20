package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"

	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);

	fmt.Print("Enter your name ");
	
	// error syntax || comma ok syntax
	name, _ := reader.ReadString('\n');   //Type of name is always string
	fmt.Println("Your name is",name);
}