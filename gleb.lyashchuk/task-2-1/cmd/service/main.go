package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		minTemp := 15
		maxTemp := 30

		for j := 0; j < k; j++ {
			var op string
			var temp int
			fmt.Scan(&op, &temp)

			if op == ">=" {
				if temp > minTemp {
					minTemp = temp
				}
			} else if op == "<=" {
				if temp < maxTemp {
					maxTemp = temp
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
