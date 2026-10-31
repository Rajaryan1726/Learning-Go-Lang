package main

import "fmt"

func main(){
	fmt.Println("Today, I'm learning if-else statement in go lang.")

	loginCount := 5

	if loginCount < 10 {
		fmt.Println("Login count is less than 10")
	} else if loginCount == 10 {
		fmt.Println("Login count is equal to 10")
	} else {
		fmt.Println("Login count is greater than 10")
	}


	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// if err != nil {

	// }
}