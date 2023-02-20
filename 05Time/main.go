package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study");

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("Monday 01-02-2006"))

	createdDate := time.Date(2001,time.August,1,12,00,00,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Monday"))
}