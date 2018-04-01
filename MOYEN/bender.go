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

func (mybender *bender) checkobstacle (town [][]string) {

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

		if move == "SOUTH" { mybender.ypos++}
		if move == "NORTH" { mybender.ypos--}
		if move == "EAST" {mybender.xpos++}
		if move == "WEST" {mybender.xpos--}
		if move == "TELEPORT" {
			for _, newpos := range teleports {
				if newpos.xpos != mybender.xpos && newpos.ypos != mybender.ypos {
					mybender.xpos = newpos.xpos
					mybender.ypos = newpos.ypos
					break
				}
			}
		}
		if move == "OBSTACLE" {
			mybender.checkobstacle(town)
		}
		if move == "INVERSE" {
			mybender.inverse = !mybender.inverse
		}
		if move == "END" {
			tocontinue = false
		} else {
			solution = append(solution, move)
			direction = move
		}
		fmt.Fprintf(os.Stderr, "solution: %s %d %d \n", solution, mybender.xpos, mybender.ypos)
	}


	fmt.Fprintf(os.Stderr, "position: %s %s %d %d \n", mybender.xpos, mybender.ypos)
	fmt.Fprintf(os.Stderr, "teleport: %d \n", teleports)
	fmt.Println(solution)// Write answer to stdout
}