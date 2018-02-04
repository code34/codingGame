package main

import "fmt"
import "os"
import "bufio"

// Affiche un texte passé en paramètre
// en utilisant des lettres de hauteur H, largeur L 
// selon une font ASCII passé en paramètre

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
	
	// Transforme le texte sous forme de slices d'index A=0, B=1, etc.
	// Passe tous les caractères en upper
	// Transforme les charactères non [a-z][A-Z] en ?
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
	
	// Pour chacune des lignes des fonts
	// Pour chacun des charactères
	// Affiche l'ensemble des charactères le composant
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