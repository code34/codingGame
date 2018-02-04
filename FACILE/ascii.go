package main

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Split(bufio.ScanLines)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&L)
	
	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&H)
	
	var WORD string
	scanner.Scan()
	WORD = scanner.Text()
	var lines []string
	var indexes []int

	// initialise le slice avec l'alphabet ascii
	// chaque ligne represente les lignes composants les lettres
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	// Parcoure l'ensemble du slice text
	// à la recherche de la lettre
	var index int;
	for indexinput := 0; indexinput < len(WORD); indexinput++ {
			index = int(WORD[indexinput])            
			if((index > 96 ) && (index < 123)) { index -= 32 }
			if((index > 64) && (index <91)) { index -= 65} else {
				if!(index == 63) {index = 26}
			}
			indexes = append(indexes, index)
			fmt.Fprintf(os.Stderr, "Debug index: %d \n", indexes )
	}
  
	var beginchar int
	var endchar int
	var safeSubstring string

	fmt.Fprintf(os.Stderr, "Debug WORD: %s INDEX: %d LENCHARS %d NBCHARS %d \n", WORD, index, L)
	 
	for i := 0; i < H; i++ {
			safeSubstring = ""
			for indexinput := 0; indexinput < len(WORD); indexinput++ {
				index = indexes[indexinput]
				beginchar = index * L
				endchar = beginchar + L
				runes := []rune(lines[i])
				safeSubstring = safeSubstring + string(runes[beginchar:endchar])
			}
			fmt.Println(safeSubstring)
	};
}