package main

import "fmt"
import "os"

type site struct{
	siteId, x, y, radius, gold, maxMineSize, structureType, owner, param1,param2 int
}

type unit struct {
	x, y, owner, unitType, health int
}

type queen struct {
	unit
	//x, y, owner, unitType, health int
}

func (myqueen *queen) distance(x int, y int) int {
	newx := myqueen.x - x
	newy := myqueen.y - y
	if newx < 0 { newx = -newx}
	if newy < 0 { newy = -newy}
	return (newx + newy)
}

func (mysite *site) distance(x int, y int) int {
	newx := mysite.x - x
	newy := mysite.y - y
	if newx < 0 { newx = -newx}
	if newy < 0 { newy = -newy}
	return (newx + newy)
}

func (myqueen *queen) moveTo(x int, y int) {
	fmt.Fprintf(os.Stdout, "MOVE %d %d\n",x, y)
}

func findSafeZone(towers []site) (x,y int) {
	num := 0
	result := 0
	for key, tower := range towers {
		count := 0
		for _, dest := range towers {
			if tower.distance(dest.x, dest.y) < 60 {
				count++
			}
		}
		if count > num { 
			num = count
			result = key
		}
	}
	return towers[result].x, towers[result].y
}

func main() {
	var queens [2]queen
	var numSites int
	var initstartx,initstarty int

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

	var retreat bool

	for {
		// touchedSite: -1 if none
		var gold, touchedSite int
		fmt.Scan(&gold, &touchedSite)
		
		for i := 0; i < numSites; i++ {
			// gold: used in future leagues
			// maxMineSize: used in future leagues
			// structureType: -1 = No structure, 2 = Barracks
			// owner: -1 = No structure, 0 = Friendly, 1 = Enemy
			var siteId, gold, maxMineSize, structureType, owner, param1, param2 int
			fmt.Scan(&siteId, &gold, &maxMineSize, &structureType, &owner, &param1, &param2)
			value := allsites[siteId]
			value.siteId = siteId
			value.gold = gold
			value.maxMineSize = maxMineSize
			value.structureType = structureType
			value.owner = owner
			value.param1 = param1
			value.param2 = param2
			allsites[siteId] = value
			//fmt.Fprintf(os.Stderr, "site:  siteid:%d gold:%d maxMineSize:%d structureType:%d owner:%d param1:%d param2:%d \n", siteId, gold, maxMineSize, structureType, owner, param1, param2)
		}
		var numUnits int
		var allunits []unit

		fmt.Scan(&numUnits)
		for i := 0; i < numUnits; i++ {
			// unitType: -1 = QUEEN, 0 = KNIGHT, 1 = ARCHER
			var x, y, owner, unitType, health int
			fmt.Scan(&x, &y, &owner, &unitType, &health)
			if initstartx == 0 && owner == 0 && unitType == -1 {
				initstartx = x
				initstarty = y
				fmt.Fprintf(os.Stderr, "start x:%d y:%d\n", initstartx, initstarty)
			}
			if unitType == -1 {
				var newqueen queen
				newqueen.x = x
				newqueen.y = y
				newqueen.owner = owner
				newqueen.unitType = unitType
				newqueen.health = health
				queens[owner] = newqueen
			} else {
				var newunit unit
				newunit.x = x
				newunit.y = y
				newunit.owner = owner
				newunit.unitType = unitType
				newunit.health = health
				allunits = append(allunits, newunit)
			}
			//fmt.Fprintf(os.Stderr, "units: unit.x:%d unit.y:%d owner:%d unitType:%d health:%d  \n", newunit.x, newunit.y, newunit.owner, newunit.unitType, newunit.health)
			//fmt.Fprintf(os.Stderr, "queens: %d \n", queens)
		}
		
		var countknight int
		for _,tempunits := range allunits {
			if tempunits.unitType == 0 && tempunits.owner == 0 {
				countknight ++
			}
		}

		min := 2000
		var target site
		haveknight := false
		havearcher := false
		havegiant := false
		var listtower []site
		var listmine []site
		var barcher,bknight,bgiant site
		var enemytowercount int

		// on se deplace vers le prochain site vide -- done
		// on se deplace pour eviter les enemis -- des qu'inférieur à 60m
		// ou pour se placer au milieu de la zone de defense
		// ou pour se placer en bordure d écran
		// ou pour upgrader un site existant

		for _, barrack := range allsites {
			if queens[0].distance(barrack.x, barrack.y) < min && barrack.owner == -1 {
				target = barrack
				min = queens[0].distance(barrack.x, barrack.y)
			}
			if barrack.owner == 1 && barrack.structureType == 1 {
				enemytowercount++
			}
			//fmt.Fprintf(os.Stderr, "params: owner %d param2 %d type %d\n", barrack.owner, barrack.param2, barrack.structureType)
			if barrack.owner == 0 && barrack.structureType == 2 {
				switch(barrack.param2) {
					case 0:
						haveknight = true
						bknight = barrack
					case 1:
						havearcher = true
						barcher = barrack
					case 2:
						havegiant = true
						bgiant = barrack
				}
			} else if barrack.owner == 0 && barrack.structureType == 0 {
					listmine = append(listmine, barrack)
			} else if barrack.owner == 0 && barrack.structureType == 1 {
					listtower = append(listtower, barrack)
			}
		}

		retreat = false
		min = 400
		for _, theunit := range allunits {
			if  theunit.owner == 1 && queens[0].distance(theunit.x, theunit.y) < min {
				retreat = true
			}
		}
		

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		// First line: A valid queen action
		// Second line: A set of training instructions
		fmt.Fprintf(os.Stderr,"RETREAT %b \n", retreat)
		if touchedSite != -1 && allsites[touchedSite].owner == -1 {
			if retreat {
				fmt.Fprintf(os.Stdout,"BUILD %d TOWER\n", touchedSite)
			} else if !haveknight {
				fmt.Fprintf(os.Stdout,"BUILD %d BARRACKS-KNIGHT\n", touchedSite)
			} else if len(listmine) < 2 && allsites[touchedSite].maxMineSize > 2 && allsites[touchedSite].gold > 10 {
				fmt.Fprintf(os.Stdout,"BUILD %d MINE\n", touchedSite)
			} else if len(listtower) < 2 {
				fmt.Fprintf(os.Stdout,"BUILD %d TOWER\n", touchedSite)
			} else if !havearcher {
				fmt.Fprintf(os.Stdout,"BUILD %d BARRACKS-ARCHER\n", touchedSite)
			} else if !havegiant {
				fmt.Fprintf(os.Stdout,"BUILD %d BARRACKS-GIANT\n", touchedSite)
			} else {
				if allsites[touchedSite].gold > 10 {
					fmt.Fprintf(os.Stdout,"BUILD %d MINE\n", touchedSite)
				} else {
					fmt.Fprintf(os.Stdout,"BUILD %d TOWER\n", touchedSite)
				}
			}
		} else {
			if !retreat && (allsites[touchedSite]).structureType == 0 && allsites[touchedSite].param1 < allsites[touchedSite].maxMineSize && allsites[touchedSite].owner == 0 && allsites[touchedSite].gold > 50 {
				//fmt.Fprintf(os.Stderr,"BUILD max: %d %d\n", allsites[touchedSite].maxMineSize, allsites[touchedSite].param2)
				fmt.Fprintf(os.Stdout,"BUILD %d MINE\n", touchedSite)
			} else if !retreat && (allsites[touchedSite]).structureType == 1 && allsites[touchedSite].param1 < allsites[touchedSite].param2 && allsites[touchedSite].owner == 0 {
				//fmt.Fprintf(os.Stderr,"BUILD max: %d %d\n", allsites[touchedSite].maxMineSize, allsites[touchedSite].param2)
				fmt.Fprintf(os.Stdout,"BUILD %d TOWER\n", touchedSite)
			} else if retreat && len(listtower) > 1 {
				newx, newy := findSafeZone(listtower)
				queens[0].moveTo(newx, newy)
			} else {
				queens[0].moveTo(target.x, target.y)
			}

			// on se deplace vers le prochain site vide
			// on se deplace pour eviter les enemis 
			// ou pour se placer au milieu de la zone de defense
			// ou pour se placer en bordure d écran
			// ou pour upgrader un site existant
		}

		//fmt.Fprintf(os.Stdout,"BUILD %d BARRACKS-KNIGHT \n", touchedSite)

		if !retreat {
			if enemytowercount > 2 && bgiant.x > 0 {
				fmt.Fprintf(os.Stdout, "TRAIN %d\n", bgiant.siteId)
			} else {
				fmt.Fprintf(os.Stdout, "TRAIN %d\n", bknight.siteId)
			}
		} else if retreat {
			if countknight > 0 {
				fmt.Fprintf(os.Stdout, "TRAIN %d\n", barcher.siteId)
			} else {
				fmt.Fprintf(os.Stdout, "TRAIN %d\n", bknight.siteId)
			}
		} else {
			fmt.Fprintf(os.Stdout, "TRAIN %d\n", bgiant.siteId)
		}
	}
}