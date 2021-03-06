/*Bender est un robot dépressif qui se soigne par la fête et l’alcool. Pour lui épargner 
une vie de débauche, ses créateurs ont reprogrammé son système de contrôle avec 
une intelligence plus rudimentaire. Malheureusement, il en a perdu tout son humour
et se retrouve rejeté par ses anciens amis. Bender est maintenant livré à lui-même et
 erre dans les rues de Futurama avec l'intention d’en finir dans une cabine à suicide.

Pour l'intercepter et le sauver d'une mort quasi-certaine, les autorités vous chargent 
d'une mission : écrire un programme qui permettra de prévoir le chemin parcouru par
Bender. Pour cela, on vous fournit la logique de la nouvelle intelligence avec laquelle
il a été programmé ainsi qu'une carte de la ville.*/

package main

import "fmt"
import "os"
import "bufio"
//import "strings"
//import "strconv"

type bender struct {
	xpos int
	ypos int
	casseur bool
	inverse bool
}

type teleport struct {
	xpos int
	ypos int
}

func (mybender *bender) convertmove (move string) (int,int) {
	xpos := 0
	ypos := 0
	switch move {
		case "SOUTH" : 
			xpos = mybender.xpos
			ypos = mybender.ypos + 1
		case "NORTH":
			xpos = mybender.xpos
			ypos = mybender.ypos - 1
		case "WEST":
			xpos = mybender.xpos - 1
			ypos = mybender.ypos
		case "EAST":
			xpos = mybender.xpos + 1
			ypos = mybender.ypos
	}
	return xpos, ypos
}

func (mybender *bender) changedir (town [][]string, move string) string {
	var dir []string
	if mybender.inverse {
		dir = []string{"WEST","NORTH", "EAST", "SOUTH"}
	} else {
		dir = []string{"SOUTH","EAST", "NORTH", "WEST"}
	}
	dir2 := []string{move}
	dir = append(dir2, dir...)

	for _, value := range dir {
		xpos, ypos := mybender.convertmove(value)
		if !mybender.checkobstacle(town, xpos, ypos) {
			mybender.xpos = xpos
			mybender.ypos = ypos
			return value
		}
	}
	return ""
}

func (mybender *bender) checkobstacle (town [][]string, x int, y int) bool {
	if !mybender.casseur {
		if town[y][x] == "X" || town[y][x] == "#" { return true }
	} else {
		if town[y][x] == "#" { return true }
	}
	return false;
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L, C int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&L, &C)
	var town [][]string
	var teleports []teleport
	var mybender bender
	var solution []string
	lastmove := "SOUTH"

	for i := 0; i < L; i++ {
		var street []string
		scanner.Scan()
		//fmt.Fprintf(os.Stderr, "%s \n", scanner.Text())
		for key, value := range scanner.Text() {
			char := fmt.Sprintf("%c", value)
			if char == "@" {
				mybender.xpos = key
				mybender.ypos = i
			}
			if char == "T" {
				myteleport := teleport{}
				myteleport.xpos = key
				myteleport.ypos = i
				teleports = append(teleports, myteleport)
			}
			street = append (street, char)
		}
		//fmt.Fprintf(os.Stderr, "%s %d \n", street, i)
		town = append(town, street)
	}

	tocontinue := true
	for tocontinue {
		var move string
		pos := town[mybender.ypos][mybender.xpos]
		switch pos {
			case "@" : move = "SOUTH"
			case "$" : move = "END"
			case "S" : move = "SOUTH"
			case "E" : move = "EAST"
			case "W" : move = "WEST"
			case "N" : move = "NORTH"
			case "I" : 
				move = lastmove
				mybender.inverse = !mybender.inverse
			case "X": 
				town[mybender.ypos][mybender.xpos] = " "
				move = lastmove
			case "B": 
				move = lastmove
				mybender.casseur = !mybender.casseur
			case "T": 
				for _, newpos := range teleports {
					if newpos.xpos != mybender.xpos || newpos.ypos != mybender.ypos {
						mybender.xpos = newpos.xpos
						mybender.ypos = newpos.ypos
						break
					}
				}
				move = lastmove
			case " ": 
				move = lastmove
		}

		if len(solution) > L * C {
			tocontinue = false
		} else if move != "END" {
			move = mybender.changedir(town, move)
			lastmove = move
		} else {
			move = ""
			tocontinue = false
		}
		//fmt.Fprintf(os.Stderr, "solution: %s %d %d %s\n", solution, mybender.xpos, mybender.ypos, town[mybender.ypos][mybender.xpos])
		solution = append(solution, move)
	}
	
	if len(solution) > L * C {
		fmt.Println("LOOP")
	} else {
		for _, themove := range solution {
			fmt.Println(themove)
		}
	}
}