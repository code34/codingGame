
package main

import "fmt"
import "os"

func (table *table) convertToNum(chiffre []string) int {
	var solution int
	for w:= 0; w < 20; w++ {
		var i int
		for _, value := range table.numeration {
			index := w * table.L
			text:= value[index:index+table.L]
			if text != chiffre[i] {
				break
			}
			i++
			//fmt.Fprintf(os.Stderr, "NUM: %s \n", text)
		}
		if i == len(table.numeration) { return w}
	}
	return solution
}

func (table *table) convertToMaya(chiffre int) []string {
	var solution []string
	for _, value := range table.numeration {
		solution = append(solution, value[chiffre * table.L:(chiffre * table.L)+table.L])
	}
	return solution
}

type table struct{
	L, H int
	numeration []string
}

func main() {
	var matable table
	fmt.Scan(&matable.L, &matable.H)
	
	for i := 0; i < matable.H; i++ {
		var numeral string
		fmt.Scan(&numeral)
		matable.numeration = append(matable.numeration, numeral)
	}

	var S1 int
	fmt.Scan(&S1)
	var chiffre1 []string
	for i := 0; i < S1; i++ {
		var num1Line string
		fmt.Scan(&num1Line)
		chiffre1 = append(chiffre1, num1Line)
		//fmt.Fprintf(os.Stderr, "line: %s \n", num1Line)
	}
	
	var S2 int
	fmt.Scan(&S2)
	var chiffre2 []string
	for i := 0; i < S2; i++ {
		var num2Line string
		fmt.Scan(&num2Line)
		chiffre2 = append(chiffre2, num2Line)
		//fmt.Fprintf(os.Stderr, "line: %s \n", num2Line)
	}
	var operation string
	fmt.Scan(&operation)

	//fmt.Fprintf(os.Stderr, "chiffre1: %d \n", matable.convertToNum(chiffre1))
	//fmt.Fprintf(os.Stderr, "chiffre2: %d \n", matable.convertToNum(chiffre2))

	num1:= matable.convertToNum(chiffre1)
	num2:= matable.convertToNum(chiffre2)
	fmt.Fprintf(os.Stderr, "resultats: %d %d\n", num1, num2)

	var num int
	switch operation {
		case "+" : num = num1 + num2
		case "-" : num = num1 - num2
		case "/": num = num1 / num2
		case "*": num = num1 * num2
	}
	result := matable.convertToMaya(num)
	
	fmt.Fprintln(os.Stderr, "Debug messages...")
	for i:= range result {
		fmt.Println(result[i])
	}
}