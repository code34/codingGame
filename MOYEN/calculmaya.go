/*Suite à la découverte d'un nouveau site maya, des centaines de livres 
de mathématiques, de physique et d'astronomie ont été mis au jour.La 
fin du monde pourrait arriver plus vite que prévue, nous avons besoin de 
vous pour en être certain ! Ainsi, pour automatiser la vérification des calculs 
scientifiques maya, on vous demande de mettre au point un programme 
capable de réaliser les opérations arithmétiques basiques entre deux 
nombres maya.*/

package main

import "fmt"
import "os"
import "math"

func (table *table) convertStringToNum(chiffre []string) int {
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
		}
		if i == len(table.numeration) { return w}
	}
	return solution
}

func convertMayaToNum (mayanumber [][]string, matable table) int {
	expo := len(mayanumber) - 1
	var decimalnumber int
	for _, chiffre := range mayanumber {
		var result int
		if expo > 0 {
			result= int(math.Pow(20, float64(expo)))
			result = result * matable.convertStringToNum(chiffre)
		} else {
			result = matable.convertStringToNum(chiffre)
		}
		decimalnumber = decimalnumber + result
		expo--
	}
	return decimalnumber
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
	if len(base) == 0 { base = append(base, 0)}
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
	var chiffre2 []string
	var index int
	var mayanum1 [][]string
	var mayanum2 [][]string

	for i := 0; i < S1; i++ {
		var num1Line string
		fmt.Scan(&num1Line)
		if index < matable.H {
			chiffre1 = append(chiffre1, num1Line)
		} else {
			fmt.Fprintf(os.Stderr, "line:  \n")
			index = 0
			mayanum1 = append(mayanum1, chiffre1)
			chiffre1 = nil
			chiffre1 = append(chiffre1, num1Line)
		}
		index++
		fmt.Fprintf(os.Stderr, "line: %d %d %s \n", S1, matable.H, num1Line)
	}
	mayanum1 = append(mayanum1, chiffre1)
	
	var S2 int
	fmt.Scan(&S2)
	index = 0

	for i := 0; i < S2; i++ {
		var num2Line string
		fmt.Scan(&num2Line)
		if index < matable.H {
			chiffre2 = append(chiffre2, num2Line)
		} else {
			fmt.Fprintf(os.Stderr, "line:  \n")
			index = 0
			mayanum2 = append(mayanum2, chiffre2)
			chiffre2 = nil
			chiffre2 = append(chiffre2, num2Line)
		}
		index++
		fmt.Fprintf(os.Stderr, "line: %d %d %s \n", S2, matable.H, num2Line)
	}
	mayanum2 = append(mayanum2, chiffre2)
	
	var operation string
	fmt.Scan(&operation)

	num1:= convertMayaToNum(mayanum1, matable)
	num2:= convertMayaToNum(mayanum2, matable)
	fmt.Fprintf(os.Stderr, "CALCUL: %d %s %d\n", num1, operation, num2)

	var num int
	switch operation {
		case "+" : num = num1 + num2
		case "-" : num = num1 - num2
		case "/": num = num1 / num2
		case "*": num = num1 * num2
	}

	newbase := convertToBase20(num)
	for i := len(newbase)-1; i > -1; i--{
		result := matable.convertToMaya(newbase[i])
		for i:= range result {
			fmt.Println(result[i])
		}
	}
}