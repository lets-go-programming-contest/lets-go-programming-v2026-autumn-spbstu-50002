package department

import (
	"errors"
	"fmt"
)

const NoSolution = -1

var ErrUnknownSign = errors.New("unknown sign")

type Department struct {
	minTemp int
	maxTemp int
}

func NewDepartment(minTemp, maxTemp int) *Department {
	return &Department{minTemp: minTemp, maxTemp: maxTemp}
}

func (department *Department) Update(sign string, temp int) error {
	switch sign {
	case ">=":
		department.minTemp = max(department.minTemp, temp)
	case "<=":
		department.maxTemp = min(department.maxTemp, temp)
	default:
		return fmt.Errorf("%w: %q", ErrUnknownSign, sign)
	}

	return nil
}

func (department *Department) OptimalTemp() int {
	if department.minTemp > department.maxTemp {
		return NoSolution
	}

	return department.minTemp
}
