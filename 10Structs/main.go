package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in GoLang")
	// no class, inheritance, super etc in golang

	Atharv := User{"Atharv","Bobade",21}
	fmt.Println(Atharv)
	// {Atharv Bobade 21}
	fmt.Printf("Atharv's details are: %+v\n",Atharv)
	// Atharv's details are: {Name:Atharv Surname:Bobade Age:21} 
	fmt.Printf("Name is %v and surname is %v",Atharv.Name,Atharv.Surname)
	// Name is Atharv and surname is Bobade
}

type User struct {
	Name    string
	Surname string
	Age     int
}
