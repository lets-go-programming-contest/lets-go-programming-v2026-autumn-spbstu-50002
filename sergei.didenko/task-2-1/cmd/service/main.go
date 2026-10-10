package main

import "fmt"

func main() {
	var departments int
	fmt.Scan(&departments)

	for i := 0; i < departments; i++ {
		minTemp := 15
		maxTemp := 30

		var employees int
		fmt.Scan(&employees)

		for j := 0; j < employees; j++ {
			var op string
			var tempEmp int
			fmt.Scan(&op, &tempEmp)

			switch op {
			case ">=":
				if tempEmp > minTemp {
					minTemp = tempEmp
				}
			case "<=":
				if tempEmp < maxTemp {
					maxTemp = tempEmp
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}

		}
	}
}
