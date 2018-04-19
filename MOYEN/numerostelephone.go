/*En rejoignant l’équipe de développement du smartphone iDroid, on
vous a confié la responsabilité de développer le gestionnaire de contacts.
Évidement, ce qu’on a oublié de vous préciser c’est que sur l’iDroid les
contraintes techniques sont fortes : le système dispose de peu de
mémoire et le processeur est aussi véloce qu’un Cyrix des années 90...

Dans le cahier des charges, deux points retiennent votre attention :

1. Assistance intelligente à la saisie des numéros
Le ou les numéros correspondant aux premiers chiffres 
saisis devront être affichés à l’utilisateur, quasi instantanément.

2. Optimisation du stockage des numéros
Les premiers chiffres communs aux numéros ne devront pas être
dupliqués en mémoire.*/

package main

import "fmt"

type numero struct {
	value int
	fils []*numero
	exist bool
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
	compteur := 0
	var liste [10]numero
	
	for i := 0; i < N; i++ {
		var telephone string
		fmt.Scan(&telephone)
		var current *numero
		for key, value := range telephone {
			num := int(value) - 48
			if key == 0 {
				current = &liste[num]
				if !current.exist {
					current.value = num
					current.exist = true
					compteur++
				}
			} else {
				current = current.addChild(num, &compteur)
			}
		}
	}
	
	// The number of elements (referencing a number) stored in the structure.
	fmt.Println(compteur)
}