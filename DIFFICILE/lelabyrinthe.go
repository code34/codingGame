package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type element struct {
	x int
	y int
}

type themap struct {
	mymap [200][100]string
}

func distance (source element, destination element) int {
	newx := destination.x - destination.y
	newy := source.x - source.y
	if newx < 0 {
		newx = newx * -1
	}
	if newy < 0 {
		newy = newy * -1
	}
	return (newx+newy)
}

func (carte *themap) findNextMove (position element) element {
	var result element
	cross := []element { {position.x - 1, position.y}, {position.x, position.y - 1}, {position.x + 1, position.y}, {position.x, position.y +1}}
	for _, neighbour := range cross {
		if carte.mymap[neighbour.x][neighbour.y] != "#" {
			result = neighbour
			fmt.Fprintf(os.Stderr, "KIRK: %d\n", result)
		}
	}
	return result
}

func main() {
	// R: number of rows.
	// C: number of columns.
	// A: number of rounds between the time the alarm countdown is activated and the time the alarm goes off.

	var R, C, A int
	var carte themap
	var path []element
	kirk := element{-1,-1}
	start := element{-1,-1}
	end := element{-1,-1}
	goal := end
	fmt.Scan(&R, &C, &A)

	for {
		// KR: row where Kirk is located.
		// KC: column where Kirk is located.
		fmt.Scan(&kirk.y, &kirk.x)

		for y := 0; y < R; y++ {
			// ROW: C of the characters in '#.TC?' (i.e. one line of the ASCII maze).
			var ROW string
			fmt.Scan(&ROW)
			for x, value  := range ROW {
				if start.x == -1 && string(value) == "T" {
					start.x = x
					start.y = y
				}
				if end.x == -1 && string(value) == "C" {
					end.x = x
					end.y = y
					goal.x = x
					goal.y = y
				}
				if y == kirk.y && x == kirk.x {
					fmt.Fprintf(os.Stderr, "K")
					carte.mymap[x][y] = "K"
				} else{
					fmt.Fprintf(os.Stderr, "%s", string(value))
					carte.mymap[x][y] = string(value)
				}
			}
			fmt.Fprintf(os.Stderr, "\n")
			if end == kirk {
				goal = start
			}
		}
		path = append(path, kirk)
		nextmove:= carte.findNextMove (kirk)
		fmt.Fprintf(os.Stderr, "%d %d %d %d %d\n", kirk, start, end, goal, path)
		fmt.Fprintf(os.Stderr, "NEXTMOVE: %d\n", nextmove)
		fmt.Println("RIGHT") // Kirk's next move (UP DOWN LEFT or RIGHT).
	}
}