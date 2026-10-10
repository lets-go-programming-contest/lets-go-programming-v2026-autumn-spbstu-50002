package calculator

import (
	"errors"
	"fmt"
)

// //////////

var operand1, operand2 int
var operator string

//////////////////////////

func Start() error {
	_, err1 := fmt.Scan(&operand1)
	if err1 != nil {
		return errors.New("Invalid first operand")
	}
	_, err2 := fmt.Scan(&operand2)
	if err2 != nil {
		return errors.New("Invalid second operand")
	}
	_, err3 := fmt.Scan(&operator)
	if err3 != nil {
		return errors.New("Invalid operation")
	}

	switch operator {
	case "+":
		fmt.Println(operand1 + operand2)
	case "-":
		fmt.Println(operand1 - operand2)
	case "*":
		fmt.Println(operand1 * operand2)
	case "/":
		if operand2 == 0 {
			return errors.New("Division by zero")
		}
		fmt.Println(operand1 / operand2)
	default:
		return errors.New("Invalid operation")
	}

	return nil
}

///////////////////////////////////////////////////////////////////////
