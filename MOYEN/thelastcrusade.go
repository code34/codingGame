/*Votre mission consiste à écrire un programme capable de prédire le chemin que va prendre Indy
 lors de sa chute dans un tunnel. Indy ne pourra rester prisonnier lors de cette première mission.*/

package main

import "fmt"
import "os"

func checkPiece(piece int, pos string) (int, int) {
	switch piece {
		case 0 : { return 0,0 }
		case 1 : { return 0,1 }
		case 2 : { if pos == "LEFT" { return 1,0 } else if pos == "RIGHT" { return -1,0 }}
		case 3 : { return 0,1 }
		case 4 : { if pos == "TOP" { return -1,0 } else if pos == "RIGHT" { return 0,1 }}
		case 5 : { if pos == "TOP" { return 1,0 } else if pos == "LEFT" { return 0,1 }}
		case 6 : { if pos == "LEFT" { return 1,0 } else if pos == "RIGHT" { return -1,0 } else if pos =="TOP" { return 0,0 }}
		case 7 : { if pos == "TOP" { return 0,1 } else if pos == "RIGHT" { return 0,1 }}
		case 8 : { if pos == "LEFT" { return 0,1 } else if pos == "RIGHT" { return 0,1 }}
		case 9 : { if pos == "LEFT" { return 0,1 } else if pos == "TOP" { return 0,1 }}
		case 10 : { if pos == "TOP" { return -1,0 } else if pos == "LEFT" { return 0,0 }}
		case 11 : { if pos == "TOP" { return 1,0 } else if pos == "RIGHT" { return 0,0 }}
		case 12: { if pos == "RIGHT" { return 0,1 }}
		case 13: { if pos == "LEFT" { return 0,1 }}
	}
	return 0,0
}

func main() {
	var W, H int
	fmt.Scan(&W, &H)
	
	labyrinth := [20][20]int{}
	
	for i := 0; i < H; i++ {
		for z := 0; z < W; z++ {
			fmt.Scan(&labyrinth[z][i])
		}
	}
	
	// EX: the coordinate along the X axis of the exit (not useful for this first mission, but must be read).
	var EX int
	fmt.Scan(&EX)
	
	for {
		var XI, YI int
		var POS string
		fmt.Scan(&XI, &YI, &POS)
		newx, newy := checkPiece(labyrinth[XI][YI],POS)
		fmt.Fprintf(os.Stderr, "piece x:%d y:%d pos:%s type:%d \n", XI, YI, POS, labyrinth[XI][YI])
		XI += newx
		YI += newy
		fmt.Println(XI, YI)
	}
}