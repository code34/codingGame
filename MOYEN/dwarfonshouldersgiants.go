package main

import "fmt"
import "os"

type node struct {
	pere *node
	fils []*node
}

func recurse (node *node, counter int) int {
	counter++
	result := counter
	for _,value := range node.fils {
		temp:= recurse (value, counter)
		fmt.Fprintf(os.Stderr, "COUNTER: %d %d %d\n", counter, node, temp)
		if temp > result { result = temp }
	}
	return result
}

func main() {
	// n: the number of relationships of influence
	var n int
	fmt.Scan(&n)
	
	mymap := make(map[int]*node)
	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
		
		current, exist := mymap[x]
		if !exist {
			current = new(node)
			mymap[x] = current
		}
		fils, exist := mymap[y]
		if !exist {
			fils = new(node)
			fils.pere = current
			mymap[y] = fils
		}
		current.fils = append(current.fils, fils)
		//fmt.Fprintf(os.Stderr, "%d %d\n", x, y)
	}

	var max int
	for _, value := range mymap {
		result := recurse(value, 0)
		if result > max {max = result }
	}
	
	//fmt.Fprintf(os.Stderr, "TOP: %d \n", max)
	fmt.Println(max)
}