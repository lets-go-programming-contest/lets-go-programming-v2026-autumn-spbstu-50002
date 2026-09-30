package main

import "fmt"

func main() {
	var firstOperand int
	var secondOperand int
	var operationOperand string

	_, err := fmt.Scanln(&firstOperand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&secondOperand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operationOperand)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operationOperand {
	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	case "*":
		fmt.Println(firstOperand * secondOperand)
	case "/":
		if secondOperand == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(firstOperand / secondOperand)
		}
	default:
		fmt.Println("Invalid Operation")

	}

}
