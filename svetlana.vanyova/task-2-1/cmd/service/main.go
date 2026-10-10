package main

import (
	"fmt"
)

type Department struct {
	min int
	max int
}

func NewDepartment() Department {
	return Department{min: 15, max: 30}
}

func (d *Department) AddTemperature(sign string, temp int) int {
	if sign == ">=" {
		if temp > d.min {
			d.min = temp
		}
	} else if sign == "<=" {
		if temp < d.max {
			d.max = temp
		}
	}
	if d.min > d.max {
		return -1
	}
	return d.min
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for range n {
		var staff int
		if _, err := fmt.Scan(&staff); err != nil {
			return
		}

		department := NewDepartment()

		for range staff {
			var (
				operation string
				value     int
			)

			if _, err := fmt.Scan(&operation, &value); err != nil {
				return
			}

			fmt.Println(department.AddTemperature(operation, value))
		}
	}
}
