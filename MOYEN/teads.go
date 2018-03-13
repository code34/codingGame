package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type node struct {
	current int   
	voisin []node
	value int
}

func (node *node) recurse (counter int) int {
	max := counter
	for _, thenode := range node.voisin {
		counter++
		result:= thenode.recurse(counter)
		if result > max {
			max = result
		}
	}
	fmt.Fprintf(os.Stderr, "max: %d\n", node.current)
	return max
}

func (funcnode *node) exist (current int) (*node, bool) {
	result := node{current, []node{}, 0}
	if funcnode.current == current { 
		return funcnode, true
	} else {
		for _,thenode := range funcnode.voisin {
			theresult, success := thenode.exist(current)
			if success { return theresult, success }
		}
	}
	return &result, false
}

func main() {
	// n: the number of adjacency relations
	
	var n int
	fmt.Scan(&n)
	
	firstnode := node {}
	
	for i := 0; i < n; i++ {
		// xi: the ID of a person which is adjacent to yi
		// yi: the ID of a person which is adjacent to xi
		var xi, yi int
		fmt.Scan(&xi, &yi)
		fmt.Fprintf(os.Stderr, "x:%d y:%d \n", xi, yi)
		
		currentnode,_ := firstnode.exist(xi)
		newnodeyi,_ := firstnode.exist(yi)
		if i == 0 { firstnode = *currentnode }
		currentnode.voisin = append(currentnode.voisin, *newnodeyi)
		//fmt.Fprintf(os.Stderr, "current: %d \n", firstnode)
	}
	
	//for i:= 0; i < len(liste) - 1; i++ {
		//currentnode := liste[i]
		result := firstnode.recurse(0)
		fmt.Fprintf(os.Stderr, "current: %d \n", result)
		//if result > 0 { currentnode.value = result }
	//}
	
	//fmt.Fprintf(os.Stderr, "liste:%d \n", liste)
	
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	// The minimal amount of steps required to completely propagate the advertisement
	fmt.Println("1")
}