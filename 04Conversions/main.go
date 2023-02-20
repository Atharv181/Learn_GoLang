package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our App")
	fmt.Println("Add the rating to our app")

	reader := bufio.NewReader(os.Stdin);

	rating, _ := reader.ReadString('\n');

	fmt.Println("Thank you for rating ", rating)

	
	//the input we get in rating is "4\n" hence we cannot convert directly into integer
	// So first we need to trim and then convert into int or float
	numrating, err := strconv.ParseFloat(strings.TrimSpace(rating),64) ;

	if err != nil {
		fmt.Println(err);
	}else{
		fmt.Println("Rating after added 1 is", numrating+1);
	}

	
}