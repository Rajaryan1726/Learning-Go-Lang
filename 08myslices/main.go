package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// Declare a slice
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// Append elements
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// Slice the first three elements
	fruitList = fruitList[:3]
	fmt.Println(fruitList)

	// Create a slice of integers
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println(highScores)

	// Append more elements
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	// Check if sorted
	fmt.Println(sort.IntsAreSorted(highScores))

	// Sort in ascending order
	sort.Ints(highScores)
	fmt.Println(highScores)

	// Remove a value from a slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}