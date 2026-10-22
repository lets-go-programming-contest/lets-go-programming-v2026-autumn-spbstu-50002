package main

import (
	"fmt"
)

const (
	defMinTemp = 15
	defMaxTemp = 30
)

func addTemp(operation string, value int) int {
	var (
		max = defMaxTemp
		min = defMinTemp
	)
	switch operation {
	case "<=":
		if value < max {
			max = value
		}
	case ">=":
		if value > min {
			min = value
		}
	default:
		break
	}

	if min > max {
		return -1
	}

	return min
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	for range n {
		var numStaff int

		_, err := fmt.Scan(&numStaff)
		if err != nil {
			return
		}

		for range numStaff {
			var (
				operation string
				value     int
			)

			_, err := fmt.Scan(&operation, &value)
			if err != nil {
				return
			}

			fmt.Println(addTemp(operation, value))
		}
	}

}
