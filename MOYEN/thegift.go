/*Les Oods veulent offrir un cadeau à l'un d'entre eux. Seulement, tous
ont un budget différent à investir dans ce cadeau, et bien sûr, leur
unique souhait est de parvenir à déterminer le partage qui soit le
plus équitable possible tout en restant dans les limites des budgets
de chacun. Les Oods réfléchissent à la méthode optimale depuis des
jours et n'ont toujours pas réussi à s'accorder sur une solution
satisfaisante. Le Docteur décide donc de leur donner un coup de
main en créant un programme. Son idée est de s'assurer que les
Oods ont assez d'argent pour acheter le cadeau et si oui, de déterminer
la somme dont chaque Ood devra s'acquitter dans la limite de son budget.*/

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