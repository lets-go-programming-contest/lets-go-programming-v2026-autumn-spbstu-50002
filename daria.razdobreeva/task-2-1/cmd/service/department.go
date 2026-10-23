package main

import (
	"errors"
	"fmt"
)

const (
	lowestTemp  = 15
	highestTemp = 30
	noSolution  = -1
)

var errUnknownSign = errors.New("unknown sign")

type Department struct {
	minTemp int
	maxTemp int
}

func NewDepartment() *Department {
	return &Department{minTemp: lowestTemp, maxTemp: highestTemp}
}

func (department *Department) Update(sign string, temp int) error {
	switch sign {
	case ">=":
		department.minTemp = max(department.minTemp, temp)
	case "<=":
		department.maxTemp = min(department.maxTemp, temp)
	default:
		return fmt.Errorf("%w: %q", errUnknownSign, sign)
	}

	return nil
}
