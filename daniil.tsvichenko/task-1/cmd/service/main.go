package main

import (
	"fmt"
	"task-1/internal/calculator"
)

func main() {
	err := calculator.Start()
	if err != nil {
		fmt.Println(err)
	}
}
