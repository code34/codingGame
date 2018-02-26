/*Le jeu se joue sur une grille rectangulaire de taille variable représentant une micro-puce.
 Certaines cellules de la grille contiennent des nœuds d'alimentation. Les autres cellules sont vides.

Le but est d'indiquer, s'ils existent, le voisin horizontal et le voisin vertical de chaque
 nœud d'alimentation.*/

package main

import "fmt"
import "os"
import "bufio"
import "strconv"

type node struct {
	x,y int
}

func add(x int, y int) int {
	return x + y
}

func checkrightnode(slice []node, currentnode node) int {
	var result int;
	result = -1
	for _,val := range slice {
		if((val.x > currentnode.x) && (val.y == currentnode.y)) {
			result = val.x
			break
		}
	}
	return result
}

func checkbotnode(slice []node, currentnode node) int {
	var result int;
	result = -1
	for _,val := range slice {
		if((val.x == currentnode.x) && (val.y > currentnode.y)) { 
			result = val.y
			break;
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// width: the number of cells on the X axis
	var width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&width)
	
	// height: the number of cells on the Y axis
	var height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&height)

	var cellule[] node
	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		for key, val := range line {
			if(string(val) == "0") {
				var newnode node
				newnode.x = key
				newnode.y = i
				cellule = append(cellule, newnode)
			}
		}
	}
	
	for _, currentnode := range cellule {
		result := strconv.Itoa(currentnode.x)+" "+strconv.Itoa(currentnode.y)+" "
		rightnode := checkrightnode(cellule, currentnode)
		if rightnode > -1 {
			result = result + strconv.Itoa(rightnode)+" "+strconv.Itoa(currentnode.y)+" "
		} else {
			result = result + "-1 -1" + " "
		}
		botnode := checkbotnode(cellule, currentnode)
		if botnode > -1 {
			result = result + strconv.Itoa(currentnode.x)+" "+strconv.Itoa(botnode)
		} else {
			result = result + "-1 -1"
		}
		fmt.Println(result)
	}
}