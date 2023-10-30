package main

import "fmt"

func main() {
	var operation string
	var number1 int
	var number2 int

	fmt.Println("Calculator Go 1.0")
	fmt.Println("========")
	fmt.Println("Which operation do you want to perform? (Add, Subtract, Multiply, Divide)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter your first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter your second number")
	fmt.Scanf("%d", &number2)

	switch operation {
	case "Add":
		fmt.Println(number1 + number2)
	case "Subtract":
		fmt.Println(number1 - number2)
	case "Multiply":
		fmt.Println(number1 * number2)
	case "Divide":
		fmt.Println(float32(number1) / float32(number2))
	default:
		fmt.Printf("No operation found matching %s", operation)
	}
}
