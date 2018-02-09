package main

import "fmt"
import "sort"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int
	fmt.Scan(&N)
	
	var chevaux[] int
	var result int
	var min int
			
	for i := 0; i < N; i++ {
		var Pi int
		fmt.Scan(&Pi)
		chevaux = append(chevaux, Pi)
	}
	
	min = 100000
	sort.Ints(chevaux)
	//fmt.Printf("%d \n", chevaux) 
	
	for i := 1; i < len(chevaux); i++ {
		result = chevaux[i] - chevaux[i-1]
		if(result < min) { min = result }
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(min)// Write answer to stdout
}