package main

import "fmt"

func main() {
	var username string = "Raj Aryan"
	fmt.Println(username)
	fmt.Printf("Variable type is: %T \n", username)

	//  Implicit declearation of Variables
	var age = 21

	fmt.Println(age)
	fmt.Printf("Variable type is: %T \n", age)

	// Without using var keyword
	country := "India"
	fmt.Println(country)
	fmt.Printf("Variable type is: %T \n", country)
}

