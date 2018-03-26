package main

import "fmt"
import "os"
import "sort"

func distance (x1, x2 *int) int {
	result := *x2 - *x1
	if result < 0 { result = result * -1 }
	return result
}

func main() {
	var N int
	fmt.Scan(&N)
	
	var habitations []int
	var xmin, xmax int
	xmin = 3000
	var median int
	var meters int
	var result int
	
	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
		habitations = append(habitations, Y)
		if X <= xmin  { xmin = X;}
		if X >= xmax  { xmax = X;}
	}
	
	if len(habitations) > 1 {
		sort.Ints(habitations)
		median = (N - 1) / 2
		result = habitations[median]
		meters = distance(&xmax, &xmin)
		for _, habitat := range habitations {
			meters += distance(&result, &habitat)
		}
	}
	fmt.Fprintf(os.Stderr, "median : %d %d %d \n", median, xmin, xmax)
	//fmt.Fprintf(os.Stderr, "result : %d meters %d \n", result, meters)
	//fmt.Fprintf(os.Stderr, "xmin: %d ymin: %d xmax : %d ymax : %d \n", xmin, xmax)
	fmt.Println(meters)
}