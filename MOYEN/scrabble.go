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
	var solution [][]string
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

	for _,word := range words {
		index := 0
		for _,letter := range letters {
			for _, letterofword := range word {
				if letter == letterofword {
					index++
				}
			}
		}
		if index > 0 {
			solution = append(solution, word)
		}
		fmt.Fprintf(os.Stderr, "%d %d %s\n", len(letters), index, word)
	}
	fmt.Fprintf(os.Stderr, "%s \n", solution)
	fmt.Fprintf(os.Stderr, "%s \n", letters)
	fmt.Println("non")// Write answer to stdout
}