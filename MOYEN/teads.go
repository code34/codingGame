package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
 
type node struct {
	ptrnode *node
	references []reference
}

type reference struct {
	value int
	ptrnode *node
}

func (node *node) addReference (value int, ptrnode *node) {
	newref:= reference{value, ptrnode}
	node.references = append(node.references, newref)
}

func (node *node) recurse (counter int, source *node) int {
	max := 0
	if len(node.references) == 1 && counter > 1 { return counter }
	counter++
	for _,reference := range node.references {
			if source != reference.ptrnode {
				var result int
				tmpsource := node
				if reference.value == 0 {
					result = reference.ptrnode.recurse(counter, tmpsource)
					reference.value = result
				} else {
					result = reference.value
				}
				if result > max { max = result }
			}
	}
	//fmt.Fprintf(os.Stderr, "RESULTS: %d\n", max)
	return max
}

func main() {
	// n: the number of adjacency relations
	var n int
	mymap := make(map[int]*node)
	
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		// xi: the ID of a person which is adjacent to yi
		// yi: the ID of a person which is adjacent to xi
		var xi, yi int
		fmt.Scan(&xi, &yi)
		
		newnode,_ := mymap[xi]
		if newnode == nil {
			tmpnode := node{}
			newnode = &tmpnode
		}
		
		targetnode,_ := mymap[yi]
		if targetnode == nil {
			tmpnode2 := node{}
			targetnode = &tmpnode2
		}
		
		newnode.addReference(0, targetnode)
		targetnode.addReference(0, newnode)
		newnode.ptrnode = newnode
		targetnode.ptrnode = targetnode
		mymap[xi] = newnode
		mymap[yi] = targetnode
		//fmt.Fprintf(os.Stderr, "input: %d %d\n", xi, yi)
	}

	resultmin := 3000
	for i:= 0; i < len(mymap);i++{
		element := mymap[i]
		max := element.recurse(0, element)
		if max < resultmin { resultmin = max }
	}

	fmt.Println(resultmin)
}