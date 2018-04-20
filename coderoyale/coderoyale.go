package main

import "fmt"
import "os"

type site struct{
	siteId, x, y, radius, ignore1, ignore2, structureType, owner, param1, param2 int
}

type unit struct {
	x, y, owner, unitType, health int
}

func main() {
	var numSites int
	fmt.Scan(&numSites)
	allsites := make(map[int]site)
	
	for i := 0; i < numSites; i++ {
		var siteId, x, y, radius int
		fmt.Scan(&siteId, &x, &y, &radius)
		value := allsites[siteId]
		value.x = x
		value.y = y
		value.radius = radius
		allsites[siteId] = value
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
			value := allsites[siteId]
			value.ignore1 = ignore1
			value.ignore2 = ignore2
			value.structureType = structureType
			value.owner = owner
			value.param1 = param1
			value.param2 = param2
			allsites[siteId] = value
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
		if touchedSite != -1 {
			fmt.Fprintf(os.Stdout,"BUILD %d BARRACKS-ARCHER \n", touchedSite)
		} else {
			fmt.Println("WAIT")
		}
		fmt.Println("TRAIN")
	}
}