package main

import "fmt"

func main() {
	fmt.Println("Enter a numner: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input / 20
	fmt.Println(output)
}
