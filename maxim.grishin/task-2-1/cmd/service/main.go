package main

import (
	"errors"
	"fmt"
)

type Bound struct {
	min int
	max int
}

var (
	errOutOfRange         = errors.New("out of range")
	errInvalidSign        = errors.New("invalid sign")
	errInvalidTemperature = errors.New("invalid temperature")
)

const (
	minN int = 1
	maxN int = 1000

	fixedMinTemp int = 15
	fixedMaxTemp int = 30
)

func readNumber() (int, error) {
	var number int

	if _, err := fmt.Scan(&number); err != nil {
		return 0, fmt.Errorf("failed to scan number: %w", err)
	}

	if number < minN || number > maxN {
		return 0, errOutOfRange
	}

	return number, nil
}

func readBound() (string, int, error) {
	var (
		sign        string
		temperature int
	)

	if _, err := fmt.Scan(&sign, &temperature); err != nil {
		return "", 0, fmt.Errorf("failed to scan bound: %w", err)
	}

	if sign != ">=" && sign != "<=" {
		return "", 0, errInvalidSign
	}

	if temperature < fixedMinTemp || temperature > fixedMaxTemp {
		return "", 0, errInvalidTemperature
	}

	return sign, temperature, nil
}

func processDepartment(bounds *Bound, numberOfEmployees int) {
	for range numberOfEmployees {
		sign, temperature, err := readBound()
		if err != nil {
			fmt.Println(-1)

			continue
		}

		if sign == ">=" {
			bounds.min = max(temperature, bounds.min)
		} else if sign == "<=" {
			bounds.max = min(temperature, bounds.max)
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
