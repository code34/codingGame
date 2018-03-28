package main

import "fmt"
//import "os"

type numero struct {
	value int
	fils []*numero
}

func (mynumero *numero) addChild (chiffre int, compteur *int) *numero {
	for _, fils := range mynumero.fils {
		if fils.value == chiffre {
			return fils
		}
	}
	fils := &numero{}
	fils.value = chiffre
	mynumero.fils = append(mynumero.fils, fils)
	*compteur++
	return fils
}

func main() {
	var N int
	fmt.Scan(&N)
	var compteur int
	var liste [10]numero
	
	for i := 0; i < N; i++ {
		var telephone string
		fmt.Scan(&telephone)
		var current *numero
		
		for key, value := range telephone {
			num := fmt.Sprintf("%d \n", int(value) - 48)
			if key == 0 {
				current = &liste[num]
			} else {
				current = current.addChild(num, &compteur)
			}
			//num := fmt.Sprintf(os.Stderr, "%d \n", int(value) - 48)
		}
	}
	
	// The number of elements (referencing a number) stored in the structure.
	fmt.Println("number")
}