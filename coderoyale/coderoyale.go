package main

import "fmt"
import "os"

func main() {
	var numSites int
	fmt.Scan(&numSites)
	
	for i := 0; i < numSites; i++ {
		var siteId, x, y, radius int
		fmt.Scan(&siteId, &x, &y, &radius)
	}
	for {
		// touchedSite: -1 if none
		var gold, touchedSite int
		fmt.Scan(&gold, &touchedSite)
		
		for i := 0; i < numSites; i++ {
			// ignore1: used in future leagues
			// ignore2: used in future leagues
			// structureType: -1 = No structure, 2 = Barracks
			// owner: -1 = No structure, 0 = Friendly, 1 = Enemy
			var siteId, ignore1, ignore2, structureType, owner, param1, param2 int
			fmt.Scan(&siteId, &ignore1, &ignore2, &structureType, &owner, &param1, &param2)
			fmt.Fprintf(os.Stderr, "input:  siteid:%d ignore1:%d ignore2:%d structureType:%d owner:%d param1:%d param2:%d \n", siteId, ignore1, ignore2, structureType, owner, param1, param2)
		}
		var numUnits int
		fmt.Scan(&numUnits)
		
		for i := 0; i < numUnits; i++ {
			// unitType: -1 = QUEEN, 0 = KNIGHT, 1 = ARCHER
			var x, y, owner, unitType, health int
			fmt.Scan(&x, &y, &owner, &unitType, &health)
			fmt.Fprintf(os.Stderr, "units: x:%d y:%d owner:%d unitType:%d health:%d  \n", x, y, owner, unitType, health)
		}
		
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		// First line: A valid queen action
		// Second line: A set of training instructions
		fmt.Println("WAIT")
		fmt.Println("TRAIN")
	}
}