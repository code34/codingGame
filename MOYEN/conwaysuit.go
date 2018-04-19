/*Votre mission est d'afficher une ligne de la suite de Conway.*/

package main

import "fmt"
//import "os"

func main() {
	var R int
	fmt.Scan(&R)
	
	var L int
	fmt.Scan(&L)
	
	var liste []int
	liste = append(liste, R)
	
	for i:= 0; i < L - 1; i++{
		current := -1
		counter := 0
		var newliste []int
		for _, value := range liste {
			if current == -1 {
				current = value
			}
			if current != value {
				newliste = append(newliste, counter) 
				newliste = append(newliste, current)
				counter = 1
				current = value
			} else {
				counter++
			}
		}
		if counter > 0 {
			newliste = append(newliste, counter) 
			newliste = append(newliste, current)
			counter = 0
		}
		liste = newliste
	}
	
	var result string
	for key, value := range liste {
		result += fmt.Sprintf("%d", value)
		if key != len(liste) -1 {
			result += fmt.Sprintf(" ")
		}
	}
	fmt.Println(result)// Write answer to stdout
}