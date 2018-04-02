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
	indexdir int
}

type teleport struct {
	xpos int
	ypos int
}

func (bender *bender) changedir () string {
	bender.indexdir++
	var dir []string
	if bender.inverse {
		//OUEST, NORD, EST, SUD
		dir = []string{"WEST", "NORTH", "EAST", "SOUTH", "LOOP"}
	} else {
		//SUD, EST, NORD et OUEST
		dir = []string{"EAST", "NORTH", "WEST", "LOOP"}
	}
	return dir[bender.indexdir]
}

func checkobstacle (town [][]string, x int, y int) bool {
	if town[y][x] == "X" || town[y][x] == "#" { return true } else { return false }
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
	var direction string

	for i := 0; i < L; i++ {
		var line []string
		scanner.Scan()
		fmt.Fprintf(os.Stderr, "%s \n", scanner.Text())
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
			line = append (line, char)
		}
		town = append(town, line)
	}

	mybender.indexdir = -1
	tocontinue := true
	for tocontinue {
		var move string
		pos := town[mybender.xpos][mybender.ypos]
		switch pos {
			case "@" : move = "SOUTH"
			case "$" : move = "END"
			case "#" : move = "OBSTACLE"
			case "X" : move = "OBSTACLE"
			case "S" : move = "SOUTH"
			case "E" : move = "EAST"
			case "W" : move = "WEST"
			case "N" : move = "NORTH"
			case "I" : move = "INVERSE"
			case "B": move = "CASSEUR"
			case "T": move = "TELEPORT"
			case " ": move = direction
		}

		if move == "SOUTH" { 
			if checkobstacle(town, mybender.xpos, mybender.ypos + 1) {
				move = mybender.changedir()
				fmt.Fprintf(os.Stderr, "move: %d \n", move)
			} else {
				mybender.ypos++
			}
		}
		if move == "NORTH" { 
			if checkobstacle(town, mybender.xpos, mybender.ypos - 1) {
				move = mybender.changedir()
			} else {
				mybender.ypos--
			}
		}
		if move == "EAST" {
			if checkobstacle(town, mybender.xpos + 1, mybender.ypos) {
				move = mybender.changedir()
			} else {
				mybender.xpos++
			}
		}
		if move == "WEST" {
			if checkobstacle(town, mybender.xpos - 1, mybender.ypos) {
				move = mybender.changedir()
			} else {
				mybender.xpos--
			}
		}
		if move == "TELEPORT" {
			for _, newpos := range teleports {
				if newpos.xpos != mybender.xpos && newpos.ypos != mybender.ypos {
					mybender.xpos = newpos.xpos
					mybender.ypos = newpos.ypos
					break
				}
			}
		}
		if move == "INVERSE" {
			mybender.inverse = !mybender.inverse
		}
		if move == "END" {
			move = ""
			tocontinue = false
		} else {
			solution = append(solution, move)
			direction = move
		}
		fmt.Fprintf(os.Stderr, "solution: %s %d %d \n", solution, mybender.xpos, mybender.ypos)
		fmt.Println(move)
	}
	//fmt.Println(solution)// Write answer to stdout
}