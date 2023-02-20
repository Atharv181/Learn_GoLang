package main

import "fmt"

const Writer = "Atharv"  // Public variable 
const surname = "Bobade" // Private variable

func main() {
	// fmt.Println("Hello Variables");

	//string
	var username string = "Atharv";
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n", username);

	//Boolean
	var isBoolean bool = true;
	fmt.Println(isBoolean);
	fmt.Printf("Variable is of type: %T \n",isBoolean);

	// Integer
	var number uint = 1234
	// var number uint8 = 1234;
	// var number uint16 = 1234;
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n",number)

	//Float
	var floatNumber float32 = 255.4222224;
	// var floatNumber float64 = 255.4222224;
	fmt.Println(floatNumber);
	fmt.Printf("Variable is of type: %T \n",floatNumber);

	//implicit type
	var website = "Hello from Website"
	fmt.Println(website);
	fmt.Printf("Variable is of type: %T \n",website);

	//no var style       &&   only can be used inside any method and not ouside
	noVarStyle := 83473.98765
	// noVarStyle = 9328746.9
	fmt.Println(noVarStyle);


	fmt.Println(Writer);
	fmt.Printf("Variable is of type: %T \n",Writer);

	fmt.Println(surname);
	fmt.Printf("Variable is of type: %T \n",surname);
}