package main

import "fmt"

func main() {
	var first, second int
	var symbol string

	_, err := fmt.Scanln(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scanln(&second)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scanln(&symbol)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch symbol {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(first / second)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
