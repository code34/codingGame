/*Chez Teads nous savons que pour maximiser l’impact d’une publicité,
il faut que le message qu’elle porte se propage largement et rapidement.
 
Le potentiel de viralité est une donnée dont vous disposez, il est représenté
par un réseau de personnes qui diffuseront un message auprès d’autres personnes.
On supposera qu’il n’y a jamais de relation cyclique dans ce réseau.
Par exemple, si une personne #1 est en relation avec une personne #2 et que #2
est en relation avec #3, alors il est impossible que #3 soit en relation directe avec #1.
 
Lorsqu’un individu diffuse un message il n’y a qu’une étape dont la durée est
indépendante du nombre de personnes avec qui il ou elle est en lien direct.
Nous considérerons que cette étape prend toujours 1 heure. */

package main

import "fmt"
import "os"

type node struct {
		pere []*node
		fils []*node
		value int
}

func main() {
	// n: the number of adjacency relations
	var n int
	fmt.Scan(&n)

	mymap := make(map[int]*node)

	for i := 0; i < n; i++ {
	// xi: the ID of a person which is adjacent to yi
	// yi: the ID of a person which is adjacent to xi
	var xi, yi int
	fmt.Scan(&xi, &yi)

		pere,_ := mymap[xi]
		if pere == nil {
			pere = &node{}
			pere.value = xi
			mymap[xi] = pere
		}
		
		fils,_ := mymap[yi]
		if fils == nil {
			fils = &node{}
			fils.value = yi
			mymap[yi] = fils
		}
		pere.fils = append(pere.fils, fils)
		fils.pere = append(fils.pere, pere)
	}

	var counter int
	mustcontinue := true
	for mustcontinue {
		var todelete []int
		counter++
		for i, _ := range mymap {
			element := mymap[i]
			flagme := false
			if len(element.pere) == 1 && len(element.fils) == 0 {
				flagme = true
			}
			if len(element.pere) == 0 && len(element.fils) == 1 {
				flagme = true
			}
			if flagme {
				todelete = append(todelete, i)
			}
		}

		for _, value := range todelete {
			element  := mymap[value]

			// s'il y a un pere et qu'il n'y a pas de fils
			if len(element.pere) == 1 && len(element.fils) == 0 {
				delete(mymap, value)
				for key, myfils := range element.pere[0].fils {
					if myfils == element {
						element.pere[0].fils = append (element.pere[0].fils[:key], element.pere[0].fils[key+1:]...)
					}
				}
			} 

			// s'il y a un fils et qu'il n'y a pas de pere
			if len(element.fils) == 1 && len(element.pere) == 0 {
				delete(mymap, value)
				//element.fils[0].pere = nil
				for key, myfils := range element.fils[0].pere {
					if myfils == element {
						element.fils[0].pere = append (element.fils[0].pere[:key], element.fils[0].pere[key+1:]...)
					}
				}
			}
		}
		if len(mymap) < 2 {mustcontinue = false}
		fmt.Fprintf(os.Stderr, "MYMAP: %d\n", len(mymap))
	}
	fmt.Println(counter)
}