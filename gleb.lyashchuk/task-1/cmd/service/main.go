package main

import (
	"fmt"
	"strconv"
)

func main() {
	var operand1, operand2, operator string

	if _, err := fmt.Scan(&operand1); err != nil {
		fmt.Println("Error reading first operand:", err)
		return
	}

	num1, err := strconv.Atoi(operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scan(&operand2); err != nil {
		fmt.Println("Error reading second operand:", err)
		return
	}

	num2, err := strconv.Atoi(operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scan(&operator); err != nil {
		fmt.Println("Error reading operator:", err)
		return
	}

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(num1 / num2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}