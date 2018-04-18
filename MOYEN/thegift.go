package main

import "fmt"
import "os"
import "sort"

func main() {
	var N int
	fmt.Scan(&N)
	
	var price int
	fmt.Scan(&price)
	var budgets []int
	var solution []int

	for i := 0; i < N; i++ {
		var budget int
		fmt.Scan(&budget)
		budgets = append(budgets, budget)
		fmt.Fprintf(os.Stderr, "budget: %d price: %d \n", budget, price)
	}
	sort.Ints(budgets)

	var index, max int
	for _, budget := range budgets {
		max = price / (len(budgets) - index)
		if budget > max { 
			solution = append(solution, max)
			price = price - max
		} else {
			solution = append(solution, budget)
			price = price - budget
		}
		index++
	}

	//fmt.Fprintf(os.Stderr, "price: %d map:%d budgets:%d \n", price, mymap, budgets)
	if price > 0 {
		fmt.Println("IMPOSSIBLE")
	} else {
		for _, value := range solution {
			fmt.Println(value)
		}
	}
}