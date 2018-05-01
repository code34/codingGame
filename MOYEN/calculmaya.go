package main

import "fmt"
import "os"
import "math"

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

func convertToBase20(chiffre int) []int {
	value := float64(chiffre)
	var base []int
	for value > 0 {
		newvalue := math.Floor(value / 20)
		if newvalue > 0 {
			base = append(base, int(math.Mod(value, 20)))
		} else {
			base = append(base, int(value))
		}
		value = newvalue
	}
	return base
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
	var index int
	var mynum1 [][]string
	for i := 0; i < S1; i++ {
		var num1Line string
		fmt.Scan(&num1Line)
		if index < matable.H {
			chiffre1 = append(chiffre1, num1Line)
		} else {
			fmt.Fprintf(os.Stderr, "line:  \n")
			index = 0
			mynum1 = append(mynum1, chiffre1)
			chiffre1 = nil
			chiffre1 = append(chiffre1, num1Line)
		}
		index++
		fmt.Fprintf(os.Stderr, "line: %d %d %s \n", S1, matable.H, num1Line)
	}
	mynum1 = append(mynum1, chiffre1)
	fmt.Fprintf(os.Stderr, "num: %d \n", mynum1)
	
	var S2 int
	fmt.Scan(&S2)
	var chiffre2 []string
	for i := 0; i < S2; i++ {
		var num2Line string
		fmt.Scan(&num2Line)
		chiffre2 = append(chiffre2, num2Line)
		fmt.Fprintf(os.Stderr, "line2: %d %d %s \n", S2, matable.H, num2Line)
	}
	
	var operation string
	fmt.Scan(&operation)
	fmt.Fprintf(os.Stderr, "operation: %s \n", operation)

	expo := len(mynum1) - 1
	fmt.Fprintf(os.Stderr, "expo: %d\n", expo)
	var num1 int
	for _, numero := range mynum1 {
		var result int
		if expo > 0 {
			result= int(math.Pow(20, float64(expo)))
		}
		num1 = result + matable.convertToNum(numero)
		fmt.Fprintf(os.Stderr, "BERRRKKK: %d %d\n", num1, result)
		expo--
	}

	num2:= matable.convertToNum(chiffre2)
	fmt.Fprintf(os.Stderr, "resultats: %d %d\n", num1, num2)

	var num int
	switch operation {
		case "+" : num = num1 + num2
		case "-" : num = num1 - num2
		case "/": num = num1 / num2
		case "*": num = num1 * num2
	}

	newbase := convertToBase20(num)
	fmt.Fprintf(os.Stderr, "newbase: %d\n", newbase)

	for i := len(newbase)-1; i > -1; i--{
		result := matable.convertToMaya(newbase[i])
		for i:= range result {
			fmt.Println(result[i])
		}
	}
}