/*Batman va rechercher les otages d'un batiment donné en sautant de fenêtre 
en fenêtre à l'aide de son grappin. Le but de Batman est d'arriver sur la fenêtre
 de la pièce où les otages se trouvent afin de désamorcer les bombes du Joker.
  Malheureusement il n'a qu'un nombre de sauts limités avant que les bombes n'explosent...*/

package main

import "fmt"

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
	
	var top int = 0
	var bot int = H - 1
	var left int = 0
	var right int = W - 1
	
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
		
		X0 = (right + left) / 2
		Y0 = (bot + top) / 2
		fmt.Printf("%d %d\n", X0, Y0)
	}
}