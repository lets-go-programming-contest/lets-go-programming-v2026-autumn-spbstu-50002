package main

import (
	"errors"
	"fmt"
)

const (
	minCount int = 1
	maxCount int = 1000

	minDepartmentTemp = 15
	maxDepartmentTemp = 30
)

var (
	errInvalidNumber      = errors.New("invalid number")
	errInvalidOperator    = errors.New("invalid operator")
	errInvalidTemperature = errors.New("invalid temperature")
)

func readAndCheckNumber() (int, error) {
	var number int

	if _, err := fmt.Scan(&number); err != nil {
		return 0, fmt.Errorf("failed to scan number: %w", err)
	}

	if number < minCount || number > maxCount {
		return 0, fmt.Errorf("%w: must be between %d and %d", errInvalidNumber, minCount, maxCount)
	}

	return number, nil
}

func readAndCheckConstraint() (string, int, error) {
	var (
		operator string
		temp     int
	)

	if _, err := fmt.Scan(&operator); err != nil {
		return "", 0, fmt.Errorf("failed to scan operator: %w", err)
	}

	if operator != ">=" && operator != "<=" {
		return "", 0, fmt.Errorf("%w: %q", errInvalidOperator, operator)
	}

	if _, err := fmt.Scan(&temp); err != nil {
		return "", 0, fmt.Errorf("failed to scan temperature: %w", err)
	}

	if temp < minDepartmentTemp || temp > maxDepartmentTemp {
		return "", 0, fmt.Errorf(
			"%w: must be between %d and %d",
			errInvalidTemperature,
			minDepartmentTemp,
			maxDepartmentTemp,
		)
	}

	return operator, temp, nil
}

func processEmployeeConstraints(department *Department, numberOfEmployees int) error {
	for range numberOfEmployees {
		operator, temp, err := readAndCheckConstraint()
		if err != nil {
			return err
		}

		department.ApplyConstraint(operator, temp)

		if department.minTemp > department.maxTemp {
			fmt.Println(-1)
		} else {
			fmt.Println(department.minTemp)
		}
	}

	return nil
}

func main() {
	numberOfDepartments, err := readAndCheckNumber()
	if err != nil {
		fmt.Println(err)

		return
	}

	for range numberOfDepartments {
		numberOfEmployees, err := readAndCheckNumber()
		if err != nil {
			fmt.Println(err)

			return
		}

		department := NewDepartment(minDepartmentTemp, maxDepartmentTemp)
		if err := processEmployeeConstraints(department, numberOfEmployees); err != nil {
			fmt.Println(err)

			return
		}
	}
}
