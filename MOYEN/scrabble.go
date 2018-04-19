/*Lorsqu'on joue au Scrabble©, chaque joueur tire 7 lettres et doit 
trouver le mot qui permet d'obtenir le meilleur score en utilisant 
ces lettres. Un joueur n'a pas l'obligation de former un mot de 7 lettres,
le mot peut être plus court. La seule contrainte est que le mot ne peut
être constitué que des lettres que le joueur a tiré. Par exemple, avec les
lettres etaenhs, les mots possibles (en anglais) sont : ethane, hates,
sane, ant. Votre objectif est de trouver le mot qui marque le plus de
points en utilisant les lettres disponibles (1 à 7 lettres).*/

package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var words [][]string
	var solution string
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&N)
	
	for i := 0; i < N; i++ {
		scanner.Scan()
		//fmt.Fprintf(os.Stderr, "%s \n", scanner.Text())
		word := strings.Split(scanner.Text(), "")
		words = append(words, word)
	}
	
	scanner.Scan()
	letters := strings.Split(scanner.Text(), "")
	fmt.Fprintf(os.Stderr, "%s \n", letters)

	scoremax := 0
	for _,word := range words {
		score := 0
		lettercount := 0
		for _,letter := range letters {
			for _, letterofword := range word {
				if letter == letterofword {
					if letter == "e" || letter == "a" || letter == "i" || letter == "o" || letter == "n" || letter == "r" || letter == "t" || letter == "l" || letter == "s" || letter == "u" {
						score++
					} else if letter == "d" || letter == "g" {
						score = score + 2
					} else if letter == "b" || letter == "c" || letter == "m" || letter == "p" {
						score = score + 3
					} else if letter == "f" || letter == "h" || letter == "v" || letter == "w" || letter == "y" {
						score = score + 4
					} else if letter == "k" {
						score = score + 5
					} else if letter == "j" || letter == "x" {
						score = score + 8
					} else if letter == "q" || letter == "z" {
						score = score + 10
					}
					lettercount++
					break
				}
			}
		}
		if score > scoremax && len(word) <= lettercount {
			solution = strings.Join(word, "")
			scoremax = score
		}
		fmt.Fprintf(os.Stderr, "%d %d %s %d %d\n", len(letters), score, word, len(word), lettercount)
	}
	fmt.Println(solution)// Write answer to stdout
}