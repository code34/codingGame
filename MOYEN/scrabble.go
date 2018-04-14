package main

import "fmt"
import "os"
import "bufio"
import "strings"
//import "strconv"

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

	max := 0
	for _,word := range words {
		index := 0
		for _,letter := range letters {
			for _, letterofword := range word {
				if letter == letterofword {
					if letter == "e" || letter == "a" || letter == "i" || letter == "o" || letter == "n" || letter == "r" || letter == "t" || letter == "l" || letter == "s" || letter == "u" {
						index++
					} else if letter == "d" || letter == "g" {
						index = index + 2
					} else if letter == "b" || letter == "c" || letter == "m" || letter == "p" {
						index = index + 3
					} else if letter == "f" || letter == "h" || letter == "v" || letter == "w" || letter == "y" {
						index = index + 4
					} else if letter == "k" {
						index = index + 5
					} else if letter == "j" || letter == "x" {
						index = index + 8
					} else if letter == "q" || letter == "z" {
						index = index + 10
					}
				}
			}
		}
		if index > max {
			solution = strings.Join(word, "")
			max = index
		}
		fmt.Fprintf(os.Stderr, "%d %d %s\n", len(letters), index, word)
	}
	//fmt.Fprintf(os.Stderr, "%s \n", solution)
	fmt.Println(solution)// Write answer to stdout
}