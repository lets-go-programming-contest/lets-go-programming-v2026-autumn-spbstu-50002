package main

import (
	"fmt"
)

type Bound struct {
	min int
	max int
}

func main() {
	var (
		numberOfDepartments, numberOfEmployees int
	)
	_, err := fmt.Scan(&numberOfDepartments)
	if err != nil || numberOfDepartments < 1 || numberOfDepartments > 1000 {
		return
	}

	for i := 0; i < numberOfDepartments; i++ {
		_, err = fmt.Scan(&numberOfEmployees)
		if err != nil || numberOfEmployees < 1 || numberOfEmployees > 1000 {
			return
		}

		departmentBounds := Bound{min: 15, max: 30}
		for j := 0; j < numberOfEmployees; j++ {
			var (
				sign        string
				temperature int
			)

			_, err = fmt.Scan(&sign)
			if err != nil || (sign != ">=" && sign != "<=") {
				return
			}

			_, err = fmt.Scan(&temperature)
			if err != nil || temperature < 15 || temperature > 30 {
				return
			}

			if sign == ">=" {
				departmentBounds.min = max(temperature, departmentBounds.min)
			} else {
				departmentBounds.max = min(temperature, departmentBounds.max)
			}

			var result int

			if departmentBounds.min > departmentBounds.max {
				result = -1
			} else {
				result = departmentBounds.min
			}

			fmt.Println(result)
		}
	}
}
