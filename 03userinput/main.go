package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	// Comma, ok OR Comma, err
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", name)

}