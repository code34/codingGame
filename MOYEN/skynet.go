/*Votre virus a créé une backdoor sur le réseau Skynet vous permettant d'envoyer
 de nouvelles instructions au virus en temps réél.

Vous décidez de passer à l'attaque active en empêchant Skynet de communiquer
 sur son propre réseau interne.

Le réseau Skynet est divisé en sous-réseaux. Sur chaque sous-réseau un agent Skynet 
a pour tâche de transmettre de l'information en se déplaçant de noeud en noeud le 
long de liens et d'atteindre une des passerelles qui mène vers un autre sous-réseau.

Votre mission est de reprogrammer le virus pour qu'il coupe les liens dans le but 
d'empêcher l'agent Skynet de sortir de son sous-réseau et ainsi d'informer le hub 
central de la présence de notre virus.*/

package main

import "fmt"
import "os"
import "strconv"

type link struct {
	s,t int
}

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)
	fmt.Fprintf(os.Stderr, "nodes: %d links: %d gw: %d \n", N, L, E) 
	var liste []link

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var newlink link
		fmt.Scan(&newlink.s, &newlink.t)
		liste = append(liste, newlink)
		fmt.Fprintf(os.Stderr, "link: %d %d \n", newlink.s, newlink.t)
	}

	var listegw []int
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var gw int
		fmt.Scan(&gw)
		listegw = append(listegw, gw)        
	}
	
	var index int
	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		var result = ""
		
		fmt.Scan(&SI)
		//fmt.Fprintf(os.Stderr, "AGENT: %d \n", SI) 

		for _, gw := range listegw {
			for key, currentlink := range liste {
				if((currentlink.t == gw)||(currentlink.t == SI)) && ((currentlink.s == SI)||(currentlink.s == gw)) {
					result = strconv.Itoa(SI)+" "+strconv.Itoa(gw)
					index = key
				}
			}
		}
		
		if result == "" {
			for _, gw := range listegw {
				for key, currentlink := range liste {
					if (currentlink.t == gw)||(currentlink.s == gw) {
						result = strconv.Itoa(currentlink.s)+" "+strconv.Itoa(currentlink.t)
						index = key
					}
				}
			}
		}
		
		// Example: 0 1 are the indices of the nodes you wish to sever the link between
		liste = append(liste[:index], liste[index+1:]...)
		fmt.Println(result)
	}
}