// Калькулятор 5130904/50002 Диденког Сергей
package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var operation string

	_, error1 := fmt.Scan(&num1)
	if error1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, error2 := fmt.Scan(&num2)
	if error2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, error3 := fmt.Scan(&operation)
	if error3 != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid operation")
		return
	}
}
