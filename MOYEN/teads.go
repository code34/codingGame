package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func parseTree(tableau []node, index int, topvalue int, botvalue int) (int, int) {
	var max int
	topvalue += 1
	tableau[index].topvalue = topvalue
	if len(tableau[index].fils) > 0 {botvalue = 1}
	for _,element := range tableau[index].fils {
		_,temp := parseTree(tableau, element, topvalue, botvalue)
		//fmt.Fprintf(os.Stderr, "data:%d %d\n", element, temp)
		if temp > max { max = temp }
	} 
	tableau[index].botvalue = (botvalue + max)
	return topvalue, (botvalue + max)
}

func findTop(tableau []node, start int) int{
	var top int
	var run bool
	run = true
	
	top = tableau[start].pere
	for run {
		if tableau[top].pere > 0 {
			top = tableau[top].pere
		} else {
			run = false   
		}
	}
	return top
}

type node struct {
	pere int
	fils []int
	topvalue int
	botvalue int
}

func (this *node) addFils(element int) {
	this.fils = append(this.fils, element)
}

func main() {
	// n: the number of adjacency relations
	var n,max,top int
	fmt.Scan(&n)
	
	//var tableau []node
	//mymap := make(map[int]node)
	//var tableau [5]node
	var tableau = make([]node,2000,2000) 

	for i := 0; i < n; i++ {
		// xi: the ID of a person which is adjacent to yi
		// yi: the ID of a person which is adjacent to xi
		var xi, yi int
		fmt.Scan(&xi, &yi)
		if top == 0 { top = xi }
		tableau[xi].addFils(yi)
		tableau[yi].pere = xi
		fmt.Fprintf(os.Stderr, "data:%d %d\n", xi, yi)
	}

	top = findTop(tableau, top)
	parseTree(tableau, top, -1, -1)
	fmt.Fprintf(os.Stderr, "top:%d tree:%d\n", top, tableau)
	
	var oldtop int
	var oldbot int
	var finish bool
	var index int
	index = top
	finish = true

	for finish { 
		if tableau[index].botvalue == tableau[index].topvalue {
			max = tableau[index].botvalue
			if len (tableau[tableau[index].pere].fils) > 1 {
				oldtop = tableau[tableau[index].pere].topvalue
				oldbot = tableau[tableau[index].pere].botvalue
				if oldtop > oldbot { max = oldtop } else { max = oldbot } 
			}
			finish = false
			fmt.Fprintf(os.Stderr, "MAX:%d count:%d index: %d\n", max,index)
		}
		if index == len(tableau) { finish = false }
		index = tableau[index].fils[0]
	}
	
	//fmt.Fprintf(os.Stderr, "max:%d count:%d index: %d\n", max, count, index)
	
	// The minimal amount of steps required to completely propagate the advertisement
	fmt.Println(max)
}