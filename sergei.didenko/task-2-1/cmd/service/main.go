package main

import "fmt"

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		minTemp := 15
		maxTemp := 30

		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			return
		}

		for range employees {
			var op string
			var tempEmp int

			if _, err := fmt.Scan(&op, &tempEmp); err != nil {
				return
			}

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
