package main

import "fmt"
import "os"

type element struct {
	x int
	y int
}

type themap struct {
	mymap [200][100]string
	teleport element
	controler element
}

func distance (source element, destination element) int {
	newx := destination.x - source.x
	newy := destination.y - source.y
	if newx < 0 {
		newx = newx * -1
	}
	if newy < 0 {
		newy = newy * -1
	}
	return (newx+newy)
}

func (carte *themap) checkNextPosition(position element) element {
	var result element
	cross := []element { {position.x - 1, position.y}, {position.x, position.y - 1}, {position.x + 1, position.y}, {position.x, position.y +1}}
	for _, neighbour := range cross {
		if carte.mymap[neighbour.x][neighbour.y] != "#" && carte.mymap[neighbour.x][neighbour.y] != "V"{
			result = neighbour
		}
	}
	return result
}

func (carte *themap) defineGoal(position element) element {
	var goal element
	if carte.controler.x == -1 && carte.controler.y == -1 {
		goal = carte.checkNextPosition(position)
	} else {
		goal = carte.controler
	}
	return goal
}

func (carte *themap) findNextMove (position element) element {
	var result element
	var goal element
	min := 20000
	cross := []element { {position.x - 1, position.y}, {position.x, position.y - 1}, {position.x + 1, position.y}, {position.x, position.y +1}}
	goal = carte.defineGoal(position)
	for _, neighbour := range cross {
		if carte.mymap[neighbour.x][neighbour.y] != "#" && distance(element{neighbour.x, neighbour.y}, goal) < min {
			min = distance(element{neighbour.x, neighbour.y}, goal)
			result = neighbour
		}
	}
	return result
}

func (carte *themap) findDirection (source element, target element) string {
	if source.x > target.x {
		return "LEFT"
	} else if source.x < target.x {
		return "RIGHT"
	} else if source.y > target.y {
		return "UP"
	} else {
		return "DOWN"
	}
}

func main() {
	// R: number of rows.
	// C: number of columns.
	// A: number of rounds between the time the alarm countdown is activated and the time the alarm goes off.

	var R, C, A int
	var carte themap
	var path []element
	reverse := false
	kirk := element{-1,-1}
	carte.teleport = element{-1,-1}
	carte.controler = element{-1,-1}
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
				if carte.teleport.x == -1 && string(value) == "T" {
					carte.teleport.x = x
					carte.teleport.y = y
				}
				if carte.controler.x == -1 && string(value) == "C" {
					carte.controler.x = x
					carte.controler.y = y
				}
				if y == kirk.y && x == kirk.x {
					fmt.Fprintf(os.Stderr, "K")
					carte.mymap[x][y] = "K"
				} else{
					fmt.Fprintf(os.Stderr, "%s", string(value))
					if carte.mymap[x][y] != "V" {
						carte.mymap[x][y] = string(value)
					}
				}
			}
			fmt.Fprintf(os.Stderr, "\n")
		}
		nextmove:= carte.findNextMove (kirk)
		if kirk == carte.controler || reverse {
			reverse = true
			nextmove = path[len(path)-1]
			path = path[:len(path)-1]
		} else {
			path = append(path, kirk)
		}
		direction:= carte.findDirection(kirk, nextmove)
		carte.mymap[kirk.x][kirk.y] = "V"
		
		fmt.Fprintf(os.Stderr, "%d %d %d\n", carte.teleport, carte.controler, path)
		fmt.Fprintf(os.Stderr, "NEXTMOVE: %d KIRK: %d\n", nextmove, kirk)
		fmt.Println(direction) // Kirk's next move (UP DOWN LEFT or RIGHT).
	}
}