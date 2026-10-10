package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var operation string

	if _, err := fmt.Scanln(&num1); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scanln(&num2); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scanln(&operation); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	var result int

	switch operation {
	case "-":
		result = num1 - num2
	case "+":
		result = num1 + num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
