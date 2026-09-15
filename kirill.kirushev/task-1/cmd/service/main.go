package main

import "fmt"

func main() {
	var firstOperand int
	var secondOperand int
	var operation string

	if _, err := fmt.Scanln(&firstOperand); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scanln(&secondOperand); err != nil {
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
		result = firstOperand - secondOperand
	case "+":
		result = firstOperand + secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "/":
		if secondOperand == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = firstOperand / secondOperand
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Printf("%d\n", result)
}
