package main

type Department struct {
	minTemp int
	maxTemp int
}

func NewDepartment(minTemp int, maxTemp int) *Department {
	return &Department{minTemp: minTemp, maxTemp: maxTemp}
}

func (department *Department) ApplyConstraint(operator string, temp int) {
	switch operator {
	case "<=":
		if temp < department.maxTemp {
			department.maxTemp = temp
		}

	case ">=":
		if temp > department.minTemp {
			department.minTemp = temp
		}
	}
}
