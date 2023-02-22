package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var FruitList = []string{"Apple","Mango"}
	fmt.Println(FruitList)
	fmt.Printf("Type of FruitList is %T\n",FruitList)
	// [Apple Mango]
	// Type of FruitList is []string

	FruitList = append(FruitList, "Banana","Peach")
	fmt.Println(FruitList)
	// [Apple Mango Banana Peach]

	FruitList = append(FruitList[:3])
	fmt.Println(FruitList)
	// [Apple Mango Banana]



	Scores := make([]int ,4)
	Scores[0] = 10
	Scores[1] = 20
	Scores[2] = 90
	Scores[3] = 40
//	Scores[4] = 50 //Not able to do it but

	Scores = append(Scores, 50,60,70)
	fmt.Println(Scores);
	// [10 20 90 40 50 60 70]

	fmt.Println(sort.IntsAreSorted(Scores))
	//false
	sort.Ints(Scores)
	fmt.Println(Scores)
	// [10 20 40 50 60 70 90]
	fmt.Println(sort.IntsAreSorted(Scores))
	// true


	//how to remove a value from slices based on index 
	var index int = 2;
	Scores = append(Scores[:index], Scores[index+1:]...)
	fmt.Println(Scores)
	// [10 20 50 60 70 90]
}