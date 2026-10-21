package main

import (
	"fmt"
	"sort"
)

const maxDishes = 10000

func main() {
	var dishesCount int

	if _, err := fmt.Scan(&dishesCount); err != nil || dishesCount < 1 || dishesCount > maxDishes {
		fmt.Println("Invalid dishes number")

		return
	}

	dishes := make([]int, dishesCount)

	for i := range dishesCount {
		if _, err := fmt.Scan(&dishes[i]); err != nil || dishes[i] < -maxDishes || dishes[i] > maxDishes {
			fmt.Println("Invalid dish")

			return
		}
	}

	var priority int

	if _, err := fmt.Scan(&priority); err != nil || priority < 1 || priority > dishesCount {
		fmt.Println("Invalid priority number")

		return
	}

	sort.Ints(dishes)
	fmt.Println(dishes[dishesCount-priority])
}
