package main

import "fmt"

func main(){
	fmt.Println("Welcome to learn array .")

	var fruitList [4]string 
	fruitList[0] = "Apple"
	fruitList[2] = "Banana"
	fruitList[3] = "Cherry"

	fmt.Println("Fruit list is: ", fruitList)

	fmt.Println("Fruit list length is: ", len(fruitList))

	var veggieList = [5]string{"Potato", "Tomato", "Cucumber", "Carrot", "Mushroom"}
	fmt.Println("Veggie list is: ", veggieList)
	fmt.Println("Veggie list length is: ", len(veggieList))
}