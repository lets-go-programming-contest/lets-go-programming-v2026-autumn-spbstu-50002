package main

import (
	"errors"
	"fmt"
)

const (
	minCount = 1
	maxCount = 1000
)

var errCountOutOfRange = errors.New("count out of range")

func readCount() (int, error) {
	var count int

	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("read count: %w", err)
	}

	if count < minCount || count > maxCount {
		return 0, fmt.Errorf("%w: %d", errCountOutOfRange, count)
	}

	return count, nil
}

func checkDepartment() error {
	employees, err := readCount()
	if err != nil {
		return err
	}

	department := NewDepartment()

	for range employees {
		var (
			sign string
			temp int
		)

		if _, err = fmt.Scan(&sign, &temp); err != nil {
			return fmt.Errorf("read constraint: %w", err)
		}

		if err = department.Update(sign, temp); err != nil {
			return err
		}

		if department.minTemp > department.maxTemp {
			fmt.Println(noSolution)
		} else {
			fmt.Println(department.minTemp)
		}
	}

	return nil
}

func main() {
	departments, err := readCount()
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		if err = checkDepartment(); err != nil {
			fmt.Println(err)

			return
		}
	}
}
