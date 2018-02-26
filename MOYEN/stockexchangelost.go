/*Une entreprise spécialisée dans la finance réalise une étude sur les pires investissements en
bourse et souhaite s'équiper pour cela d'un programme. Ce programme devra être capable 
d'analyser une série chronologique de valeurs d’actions pour afficher la plus grande perte 
qu'il est possible de réaliser en achetant une action à un instant t0 et en la revendant à une 
date ultérieure t1. La perte sera exprimée par la différence de valeur entre t0 et t1. 
S'il n'y a pas de perte, la perte vaudra alors 0.*/

package main

import "fmt"

func main() {
	var n int
	var max, min, result, v int64
		
	fmt.Scanf("%d\n", &n)
	
	max = 0
	min = 0
	result = 0
	
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &v)
		if v < min { min = v; if (min - max) < result { result = min - max} }
		if v > max && i < n - 1 { max = v; min = v}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(result)// Write answer to stdout
}