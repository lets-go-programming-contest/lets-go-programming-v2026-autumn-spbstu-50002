package main

import (
	"log"
	"task-1/internal/calculator"
)

func main() {
	log.SetFlags(0)
	err := calculator.Start()
	if err != nil {
		log.Println(err)
	}
}
