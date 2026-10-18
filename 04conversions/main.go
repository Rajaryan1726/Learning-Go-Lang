package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	welcome := "Welcome to the Go Conversions Program!"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	//  Conversion with trimming

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error converting input to float:", err)
	} else{
		fmt.Println("Number added:", num + 1)
	}
}