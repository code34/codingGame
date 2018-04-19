package main

import "fmt"
import "os"

func main() {
	var L, H int
	fmt.Scan(&L, &H)
	
	for i := 0; i < H; i++ {
		var numeral string
		fmt.Scan(&numeral)
		fmt.Fprintf(os.Stderr, "line: %s \n", numeral)
	}
	var S1 int
	fmt.Scan(&S1)
	
	for i := 0; i < S1; i++ {
		var num1Line string
		fmt.Scan(&num1Line)
		fmt.Fprintf(os.Stderr, "line: %s \n", num1Line)
	}
	var S2 int
	fmt.Scan(&S2)
	
	for i := 0; i < S2; i++ {
		var num2Line string
		fmt.Scan(&num2Line)
		fmt.Fprintf(os.Stderr, "line: %s \n", num2Line)
	}
	var operation string
	fmt.Scan(&operation)

	for i := 0; i < H; i++ {
		for w:= 0; w < 20;  L++ {

		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("result")// Write answer to stdout
}