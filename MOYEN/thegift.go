package main

import "fmt"
import "os"
import "sort"

func main() {
	var N int
	fmt.Scan(&N)
	
	var price int
	fmt.Scan(&price)
	mymap := make(map[int]int)
	var budgets []int
	var solution []int

	for i := 0; i < N; i++ {
		var budget int
		fmt.Scan(&budget)
		_, ok := mymap[budget]
		if !ok {
			mymap[budget] = 1
			budgets = append(budgets, budget)
		} else {
			mymap[budget]++
		}
		fmt.Fprintf(os.Stderr, "budget: %d price: %d \n", budget, price)
	}
	sort.Ints(budgets)

	for _, budget := range budgets {
		value := mymap[budget]
		if value == 1 {
			price -= budget
			solution = append(solution, budget)
		} else {
			if value * budget > price {
				budget = (price / value)
			}
			for i := 0; i < value; i++ {
				price = price - budget
				if price == 1 && i + 1 == value { budget += 1}
				solution = append(solution, budget)
			}
		}
		//fmt.Fprintf(os.Stderr, "price: %d \n", solution)
	}

	//fmt.Fprintf(os.Stderr, "price: %d map:%d budgets:%d \n", price, mymap, budgets)
	if price > 1 {
		fmt.Println("IMPOSSIBLE")
	} else {
		for _, value := range solution {
			fmt.Println(value)
		}
	}
}