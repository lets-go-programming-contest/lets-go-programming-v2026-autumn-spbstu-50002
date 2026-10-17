package service

import (
	"fmt"
	"os"
)

const (
	minInitialTemp = 15
	maxInitialTemp = 30
)

func Run() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения количества отделов:", err)

		return
	}

	for range departments {
		minTemp := minInitialTemp
		maxTemp := maxInitialTemp

		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения количества сотрудников:", err)

			return
		}

		for range employees {
			var operator string

			var tempEmp int
			if _, err := fmt.Scan(&operator, &tempEmp); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка чтения оператора и температуры:", err)

				return
			}

			switch operator {
			case ">=":
				if tempEmp > minTemp {
					minTemp = tempEmp
				}
			case "<=":
				if tempEmp < maxTemp {
					maxTemp = tempEmp
				}
			default:
				fmt.Fprintln(os.Stderr, "Ошибка: неверный оператор", operator)

				return
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
