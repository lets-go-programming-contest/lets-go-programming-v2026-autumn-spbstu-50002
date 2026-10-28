package main

import (
	"container/heap"
	"fmt"
)

type DishesHeap []int

func (h *DishesHeap) Len() int {
	return len(*h)
}

func (h *DishesHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *DishesHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *DishesHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("DishesHeap needs an integer")
	}

	*h = append(*h, val)
}

func (h *DishesHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func inRange(n, bottom, top int) bool {
	return n >= bottom && n <= top
}

const (
	minDishes = 1
	maxDishes = 10000
)

func main() {
	var (
		numOfDishes int
		priority    int
		result      int
	)

	dishesHeap := &DishesHeap{}
	heap.Init(dishesHeap)

	if _, err := fmt.Scan(&numOfDishes); err != nil || !inRange(numOfDishes, minDishes, maxDishes) {
		fmt.Println("Invalid dishes number")

		return
	}

	for range numOfDishes {
		var dish int

		if _, err := fmt.Scan(&dish); err != nil || !inRange(dish, -maxDishes, maxDishes) {
			fmt.Println("Invalid dish")

			return
		}

		heap.Push(dishesHeap, dish)
	}

	if _, err := fmt.Scan(&priority); err != nil || !inRange(priority, minDishes, numOfDishes) {
		fmt.Println("Invalid priority number")

		return
	}

	for range priority {
		tempResult, ok := heap.Pop(dishesHeap).(int)
		if !ok {
			panic("How did this even happen!?")
		}

		result = tempResult
	}

	fmt.Println(result)
}
