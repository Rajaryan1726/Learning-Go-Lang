package main

import "fmt"

func main(){
	fmt.Println("Today's topic is maps in Go!")

	courses := make(map[string]string)

	courses["JS"] = "JavaScript"
	courses["PY"] = "Python"
	courses["GO"] = "Golang"
	fmt.Println(courses)

	for key, value := range courses {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
