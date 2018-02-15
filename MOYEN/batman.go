package main

import "fmt"
import "strconv"
import "os"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H float64
	fmt.Scan(&W, &H)
	
	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)
	
	var X0, Y0 float64
	fmt.Scan(&X0, &Y0)
	
	var top float64
	var bot float64
	var left float64
	var right float64
	
	top = 0
	bot = H
	left = 0
	right = W
	
	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)        
	
		switch bombDir {
			case "U":
				bot = Y0 - 1
				left = X0
				right = X0
			case "UR":
				left = X0 + 1
				bot = Y0 - 1
			case "UL":    
				right = X0 -1
				bot = Y0 - 1
			case "L":
				right = X0 - 1
				top = Y0
				bot = Y0
			case "R":
				left = X0 + 1
				top = Y0
				bot = Y0
			case "DL":
				top = Y0 + 1
				right = X0 - 1  
			case "D":
				top = Y0 + 1
				left = X0
				right = X0
			case "DR":
				top = Y0 + 1
				left = X0 + 1
		}
		
		var posw float64
		var posy float64   
		
		
		if right == left {
			posw = left            
		} else {
			posw = math.Floor((right - left)/2) + left
		}
		
		if top == bot {
			posy = top 
		} else {
			posy = math.Floor((bot - top)/2) + top
		}
		
		X0 = posw
		Y0 = posy
		
		result := fmt.Sprintf("%s %s", strconv.FormatFloat(posw, 'f', -1, 64), strconv.FormatFloat(posy, 'f', -1, 64))
		fmt.Fprintf(os.Stderr, "Result: %s \n", result) 
		// fmt.Fprintln(os.Stderr, "Debug messages...")    
		// the location of the next window Batman should jump to.
		fmt.Println(result)
	}
}