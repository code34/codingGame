package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/



func main() {
	// nbFloors: number of floors
	// width: width of the area
	// nbRounds: maximum number of rounds
	// exitFloor: floor on which the exit is found
	// exitPos: position of the exit on its floor
	// nbTotalClones: number of generated clones
	// nbAdditionalElevators: ignore (always zero)
	// nbElevators: number of elevators
	var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
	fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)

	var floordone int = -1
	var elevators map[int]int
	elevators = make(map[int]int)

	
	for i := 0; i < nbElevators; i++ {
		// elevatorFloor: floor on which this elevator is found
		// elevatorPos: position of the elevator on its floor
		var elevatorFloor, elevatorPos int
		fmt.Scan(&elevatorFloor, &elevatorPos)
		elevators[elevatorFloor] = elevatorPos
	}
	

	for {
		// cloneFloor: floor of the leading clone
		// clonePos: position of the leading clone on its floor
		// direction: direction of the leading clone: LEFT or RIGHT
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
		if floordone < cloneFloor {
			if clonePos > dirPos && direction == "RIGHT" { action = "BLOCK" }
			if clonePos < dirPos && direction == "LEFT" { action = "BLOCK" } 
			floordone += 1
		}
		
		fmt.Println(action) // action: WAIT or BLOCK
	}
}