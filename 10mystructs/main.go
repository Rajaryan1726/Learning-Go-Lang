package main

import "fmt"



func main(){
	fmt.Println("Today's topic is structs in Go!")

	raj := User{
		Name:"Raj Aryan",
		Email: "rajaryan08161@gmail.com",
		Age:22,
		Verified:true,
	}
	fmt.Printf("User: %+v\n", raj)

	fmt.Printf("User Name: %v\n", raj.Name)
}

type User struct{
	Name string
	Email string
	Age int
	Verified bool
}