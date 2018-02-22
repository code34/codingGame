package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var n int
	var max, min, result, v int64
		
	fmt.Scanf("%d\n", &n)
	
	max = 0
	min = 0
	result = 0
	
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &v)
		if v < min { min = v; if (min - max) < result { result = min - max} }
		if v > max && i < n - 1 { max = v; min = v}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(result)// Write answer to stdout
}