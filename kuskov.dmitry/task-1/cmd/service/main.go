package main

import "fmt"

func main() {
	var first, second int
	var op string
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&op)
	if (err != nil) || (op != "+" && op != "-" && op != "*" && op != "/") {
		fmt.Println("Invalid operation")
		return
	}

	var result int
	switch op {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = first / second
	}

	fmt.Println(result)
}
