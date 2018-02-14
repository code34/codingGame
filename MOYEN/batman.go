package main

import "fmt"
import "strconv"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)
	
	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)
	
	var X0, Y0 int
	fmt.Scan(&X0, &Y0)
	
	matrix := [][]bool{}
	
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			matrix[h][w] = true
		}
	}
	
	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)
		
		switch bombDir {
			case "U":
				for h := 0; h < Y0; h++ {
					for w := 0; w < W; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
			case "UR":
				for h := 0; h < Y0; h++ {
					for w := X0; w < W; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
			case "UL":    
				for h := 0; h < Y0; h++ {
					for w := 0; w < X0; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
			case "L":
				for h := 0; h < H; h++ {
					for w := 0; w < X0; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
			case "R":
				for h := 0; h < H; h++ {
					for w := X0; w < W; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
			case "DL":
				for h := Y0+1; h < H; h++ {
					for w := 0; w < X0; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
			case "D":
				for h := Y0+1; h < H; h++ {
					for w := 0; w < W; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
			case "DR":
				for h := Y0+1; h < H; h++ {
					for w := X0; w < W; w++ {
						if matrix[h][w] == true {
							matrix[h][w] = false
						}
					}
				}
		}
		
		result := ""
		for h := 0; h < H; h++ {
			for w := 0; w < W; w++ {
				if matrix[h][w] == true { 
					result = fmt.Sprintf("%s %s", strconv.Itoa(h), strconv.Itoa(w))
					break
				}
			}
		}
		
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		
		// the location of the next window Batman should jump to.
		fmt.Println(result)
	}
}