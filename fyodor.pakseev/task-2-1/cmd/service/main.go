package main

import (
	"fmt"
)

const (
	defMinTemp = 15
	defMaxTemp = 30
)

func addTemp(operation string, value int, maxTemp, minTemp int) (int, int) {
	switch operation {
	case "<=":
		if value < maxTemp {
			maxTemp = value
		}
	case ">=":
		if value > minTemp {
			minTemp = value
		}
	}

	return minTemp, maxTemp
}

func main() {
	var departs int
	_, err := fmt.Scan(&departs)
	if err != nil {
		return
	}

	for range departs {
		var numStaff int

		_, err := fmt.Scan(&numStaff)
		if err != nil {
			return
		}

		minTemp, maxTemp := defMinTemp, defMaxTemp

		for range numStaff {
			var (
				operation string
				value     int
			)

			_, err := fmt.Scan(&operation, &value)
			if err != nil {
				return
			}
			minTemp, maxTemp = addTemp(operation, value, minTemp, maxTemp)
			if maxTemp > minTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
