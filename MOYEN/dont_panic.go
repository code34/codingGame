/*Vous devez aider Marvin et ses clones (ou est-ce l'inverse ?) à atteindre la sortie
pour s'échapper de la zone du générateur.*/

package main

import "fmt"

func main() {
	var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
	fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)

	elevators := make([]int, nbElevators+1)    
	for i := 0; i < nbElevators; i++ {
		var elevatorFloor, elevatorPos int
		fmt.Scan(&elevatorFloor, &elevatorPos)
		elevators[elevatorFloor] = elevatorPos
	}
	

	for {
		var cloneFloor, clonePos int
		var direction string
		fmt.Scan(&cloneFloor, &clonePos, &direction)

		var dirPos int        
		if exitFloor == cloneFloor {
			dirPos = exitPos 
		} else { 
			if ((len(elevators) > 0) && (cloneFloor > -1)) {
				dirPos = elevators[cloneFloor]
			}
		}

		var action = "WAIT"
		if cloneFloor > -1 {
			if clonePos > dirPos && direction == "RIGHT" { action = "BLOCK" }
			if clonePos < dirPos && direction == "LEFT" { action = "BLOCK" } 
		}
		
		fmt.Println(action) // action: WAIT or BLOCK
	}
}