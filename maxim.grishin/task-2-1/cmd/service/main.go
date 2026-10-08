package main

import (
	"fmt"
)

type Bound struct {
	min int
	max int
}

const (
	minN int = 1
	maxN int = 1000

	fixedMinTemp int = 15
	fixedMaxTemp int = 30
)

func readNumber() (int, error) {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		return 0, err
	}

	if n < minN || n > maxN {
		return 0, fmt.Errorf("out of range")
	}

	return n, nil
}

func readBound() (string, int, error) {
	var sign string
	var temp int

	if _, err := fmt.Scan(&sign, &temp); err != nil {
		return "", 0, err
	}

	if sign != ">=" && sign != "<=" {
		return "", 0, fmt.Errorf("invalid sign")
	}

	if temp < fixedMinTemp || temp > fixedMaxTemp {
		return "", 0, fmt.Errorf("invalid temperature")
	}

	return sign, temp, nil
}

func processDepartment(bounds *Bound, numberOfEmployees int) {
	for range numberOfEmployees {
		sign, temp, err := readBound()
		if err != nil {
			fmt.Println(-1)
			continue
		}

		if sign == ">=" {
			bounds.min = max(temp, bounds.min)
		} else if sign == "<=" {
			bounds.max = min(temp, bounds.max)
		}

		if bounds.min > bounds.max {
			fmt.Println(-1)
		} else {
			fmt.Println(bounds.min)
		}
	}
}

func main() {
	numberOfDepartments, err := readNumber()
	if err != nil {
		return
	}

	for range numberOfDepartments {
		numberOfEmployees, err := readNumber()
		if err != nil {
			return
		}

		departmentBounds := Bound{min: fixedMinTemp, max: fixedMaxTemp}
		processDepartment(&departmentBounds, numberOfEmployees)
	}
}
