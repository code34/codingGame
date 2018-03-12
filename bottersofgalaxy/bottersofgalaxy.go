package main

import "fmt"
import "os"
import "math"

/**
 * Made with love by AntiSquid, Illedan and Wildum.
 * You can help children learn to code while you participate by donating to CoderDojo.
 **/
 
type tower struct {
	unitId int
	team int
	unitType string
	x int
	y int
	attackRange int
	health int
	maxHealth int
	attackDamage int
	//ishit int
}

type hero struct {
	unitId int
	team int
	unitType string
	x int
	y int
	attackRange int
	health int
	maxHealth int
	shield int
	attackDamage int
	movementSpeed int
	stunDuration int
	goldValue int
	countDown1 int
	countDown2 int
	countDown3 int
	mana int
	maxMana int
	manaRegeneration int
	heroType string
	isVisible int
	itemsOwned int
	ishit bool
	isstun bool
	iscritical int
}

func (hero *hero) moveBack (tower tower, reason string) {
	//var moveX int
	//moveX = hero.x
	//if side == 0 {
	//	moveX = hero.x - (hero.attackRange -10)
	//	if moveX < 10 { moveX = 10}
	//} else {
	//	moveX = hero.x + (hero.attackRange -10)
	//	if moveX > 1900 { moveX = 1900}
	//}
	fmt.Fprintf(os.Stderr, "HERO MOVE BACK TO TOWER %s \n", reason)
	fmt.Fprintf(os.Stdout, "MOVE %d %d \n", tower.x, tower.y)
}

func distance (x1, x2, y1, y2 int) int {
	var x, y,z float64
	x = float64((x2 - x1) * (x2 - x1))
	y = float64((y2 - y1) * (y2 - y1))
	z = math.Sqrt(x + y)
	//fmt.Fprintf(os.Stderr, "DISTANCE CHECK %d \n ", z);
	return int(z)
}

// Le hero se replie en arrière d'une distance d'attack
// du minion le plus avancé sur la ligne
// sinon renvoie -1 pour dire qu'il ne faut pas se déplacer
func (hero *hero) moveBackMinions (minions map[int]minion, side int) int {
	var moveX int
	moveX = hero.x
	var maxX int

	if side == 0 { maxX = hero.x } else { maxX = 1920}
	for _, minion := range minions {
		// le hero se trouve à gauche (0) sinon à droite (1)
		if side == 0 {
			if minion.x > maxX { maxX = minion.x}
		} else {
			if minion.x < maxX { maxX = minion.x}
		}
	}

	if side == 0 {
		if maxX - hero.x < 30 {
			moveX = maxX - (hero.attackRange -10)
			if moveX < 10 { moveX = 10 }
			return moveX
		} else {
			return -1
		}
	} else {
		if hero.x - maxX < 30 {
			moveX = maxX + (hero.attackRange -10)
			if moveX > 1910 { moveX = 1910}
			return moveX
		} else{
			return -1
		}
	}
	return -1
}

func countMinionsFrontOfPos (minions map[int]minion, side int, x int) int {
	var count int
	for _, minion := range minions {
		// le hero se trouve à gauche (0) sinon à droite (1)
		if side == 0 {
			if minion.x >= x { count++}
		} else {
			if minion.x <= x { count++}
		}
	}
	return count
}

func findWeaskest (heros map[int]hero) hero {
	var result hero
	max := 3000
	for _, myhero := range heros {
		if myhero.health < max {
			max = myhero.health
			result = myhero
		}
	}
	return result
}

func (hero *hero) countMinionsFront (minions map[int]minion, side int) int {
	var count int
	for _, minion := range minions {
		// le hero se trouve à gauche (0) sinon à droite (1)
		if side == 0 {
			if minion.x >= hero.x { count++}
		} else {
			if minion.x <= hero.x { count++}
		}
	}
	return count
}

func (hero *hero) attackLeastHealthMinions (minions map[int]minion) int {
	var min int
	var result int
	min = 3000
	result = -1
	for _, minion := range minions {
		if minion.health > 0 && minion.health < min && distance(hero.x, minion.x, hero.y, minion.y) < hero.attackRange {
			min = distance(hero.x, minion.x, hero.y, minion.y)
			result = minion.unitId
		}
	}
	return result
}

func (hero *hero) attackMinions (minions map[int]minion) int {
	for _, minion := range minions {
		if minion.health > 0 && minion.health <= hero.attackDamage && distance(hero.x, minion.x, hero.y, minion.y) < hero.attackRange {
			return minion.unitId
		}
	}
	return -1
}

func (hero *hero) setHealth (health int) {
	if health == hero.health { 
		hero.ishit = false
		if hero.iscritical > 0 {
			hero.iscritical--
		}
	} else { 
		hero.iscritical++
		hero.ishit = true
		hero.health = health
	}
}

func (hero *hero) setPos (x int, y int) {
	if x == hero.x && y == hero.y { hero.isstun = true } else { hero.isstun = false; hero.x = x; hero.y = y }
}

type minion struct {
	unitId int
	team int
	unitType string
	x int
	y int
	attackRange int
	health int
	maxHealth int
	shield int
	attackDamage int
	movementSpeed int
	stunDuration int
	goldValue int
	ishit bool
	isstun bool
}

func (minion *minion) setHealth (health int) {
	if health == minion.health { 
		minion.ishit = false 
	} else { 
		minion.ishit = true; 
		minion.health = health;
	}
}

func (minion *minion) setPos (x int, y int) {
	if x == minion.x && y == minion.y { minion.isstun = true } else { minion.isstun = false; minion.x = x; minion.y = y }
}

func playersAreHit (heros map[int]hero) bool {
	var result bool
	result = false
	for _, hero := range heros {
		if hero.ishit { result = true }
	}
	//fmt.Fprintf(os.Stderr, "PLAYER GO BACK %t \n", result)
	return result
}

func minionsAreHit (minions map[int]minion) bool {
	var result bool
	result = false
	for _, minion := range minions {
		if minion.ishit { result = true }
	}
	return result
}

func minionsLastHit (minions map[int]minion, damage int) int {
	var result int
	result = -1


	for _, minion := range minions {
		//fmt.Fprintf(os.Stderr, "LAST HIT %d %d \n", minion.health, damage);
		if minion.health < damage && minion.health < (minion.maxHealth * 40 / 100) {
			result = minion.unitId
			return result
		}
	}
	return result
}

type item struct {
	itemCost int
	damage int
	health int
	maxHealth int
	mana int
	maxMana int
	moveSpeed int
	manaRegeneration int
	isPotion int
}

type forest struct {
	x	int
	y	int
	radius	int
}

func buy (items map[string]item, itemname string, gold int, heroItem *[]item) bool {
	for key, item := range items {
		switch itemname {
			case "manapotion" : 
				if gold > item.itemCost && item.mana > 0 && item.isPotion == 1{
					fmt.Fprintf(os.Stdout, "BUY %s \n", key)
					*heroItem = append (*heroItem, item)
					return true
				}
			case "healthpotion" : 
				if gold > item.itemCost && item.health > 0 && item.isPotion == 1 {
					fmt.Fprintf(os.Stdout, "BUY %s \n", key)
					*heroItem = append (*heroItem, item)
					return true
				}

			case "weapon" : 
				if gold > item.itemCost && item.damage > 0 && item.isPotion == 0 {
					fmt.Fprintf(os.Stdout, "BUY %s \n", key)
					*heroItem = append (*heroItem, item)
					return true
				}

			case "boot" : 
				if gold > item.itemCost && item.moveSpeed > 0 && item.isPotion == 0 {
					fmt.Fprintf(os.Stdout, "BUY %s \n", key)
					*heroItem = append (*heroItem, item)
					return true

				}
		}
	}
	return false
}

func main() {
	var myTeam,enemyTeam int
	fmt.Scan(&myTeam)
	if myTeam == 0 { enemyTeam = 1 } else { enemyTeam = 0}
	
	// bushAndSpawnPointCount: usefrul from wood1, represents the number of bushes and the number of places where neutral units can spawn
	var bushAndSpawnPointCount int
	fmt.Scan(&bushAndSpawnPointCount)
	
	forests := make(map[string]forest)
	for i := 0; i < bushAndSpawnPointCount; i++ {
		// entityType: BUSH, from wood1 it can also be SPAWN
		var entityType string
		var x, y, radius int
		fmt.Scan(&entityType, &x, &y, &radius)
		forests[entityType] = forest{ x, y, radius}
	}
	// itemCount: useful from wood2
	var itemCount int
	fmt.Scan(&itemCount)

	items := make(map[string]item)
	for i := 0; i < itemCount; i++ {
		// itemName: contains keywords such as BRONZE, SILVER and BLADE, BOOTS connected by "_" to help you sort easier
		// itemCost: BRONZE items have lowest cost, the most expensive items are LEGENDARY
		// damage: keyword BLADE is present if the most important item stat is damage
		// moveSpeed: keyword BOOTS is present if the most important item stat is moveSpeed
		// isPotion: 0 if it's not instantly consumed
		var itemName string
		var itemCost, damage, health, maxHealth, mana, maxMana, moveSpeed, manaRegeneration, isPotion int
		fmt.Scan(&itemName, &itemCost, &damage, &health, &maxHealth, &mana, &maxMana, &moveSpeed, &manaRegeneration, &isPotion)
		items[itemName] = item {itemCost, damage, health, maxHealth, mana, maxMana, moveSpeed, manaRegeneration, isPotion}
	}
	playerMinions := make(map[int]minion)
	enemyMinions := make(map[int]minion)
	playerHeros := make(map[int]hero)
	enemyHeros := make(map[int]hero)
	var sortedplayerHeros []int
	var heroItem []item
	countheros := 0

	for {
		var playerTower tower
		var enemyTower tower
		
		var gold int
		fmt.Scan(&gold)
		
		var enemyGold int
		fmt.Scan(&enemyGold)
		
		// roundType: a positive value will show the number of heroes that await a command
		var roundType int
		fmt.Scan(&roundType)
		if roundType < 0 {
			if countheros == 0 { 
				fmt.Println("DOCTOR_STRANGE")
				countheros++
			} else {
				fmt.Println("HULK")
			}
		}
		
		var entityCount int
		fmt.Scan(&entityCount)
		var aliveunit []int

		for i := 0; i < entityCount; i++ {
			// unitType: UNIT, HERO, TOWER, can also be GROOT from wood1
			// shield: useful in bronze
			// stunDuration: useful in bronze
			// countDown1: all countDown and mana variables are useful starting in bronze
			// heroType: DEADPOOL, VALKYRIE, DOCTOR_STRANGE, HULK, IRONMAN
			// isVisible: 0 if it isn't
			// itemsOwned: useful from wood1

			var unitId, team int
			var unitType string
			var x, y, attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration int
			var heroType string
			var isVisible, itemsOwned int
			fmt.Scan(&unitId, &team, &unitType, &x, &y, &attackRange, &health, &maxHealth, &shield, &attackDamage, &movementSpeed, &stunDuration, &goldValue, &countDown1, &countDown2, &countDown3, &mana, &maxMana, &manaRegeneration, &heroType, &isVisible, &itemsOwned)

			aliveunit= append(aliveunit, unitId)
			switch (unitType) {
				case "TOWER" : {
					unit:= tower{ unitId, team, unitType, x, y, attackRange, health, maxHealth,attackDamage }
					if team == myTeam { playerTower = unit } else { enemyTower = unit }
				}
				
				case "HERO" : {
					if team == myTeam { 
						fmt.Fprintf(os.Stderr, "TYPE %s %d \n", heroType, unitId)
						unit, exist := playerHeros[unitId] 
						if !exist {
							unit = hero{ unitId, team,unitType, x, y,attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration, heroType, isVisible, itemsOwned, true, false, 0}
							sortedplayerHeros = append(sortedplayerHeros, unitId)
						} else {
							unit.setHealth(health)
							unit.setPos(x, y)
							unit.countDown1 = countDown1
							unit.countDown2 = countDown2
							unit.countDown3 = countDown3
						}
						playerHeros[unitId] = unit
						//playerHero = unit
					} else {
						unit, exist := enemyHeros[unitId] 
						if !exist {
							unit = hero{ unitId, team,unitType, x, y,attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration, heroType, isVisible, itemsOwned, true, false, 0}
						} else {
							unit.setHealth(health)
							unit.setPos(x, y)
							unit.countDown1 = countDown1
							unit.countDown2 = countDown2
							unit.countDown3 = countDown3
						}
						enemyHeros[unitId] = unit
					}
				}
				
				case "UNIT" : {
					if team == myTeam { 
						unit, exist := playerMinions[unitId] 
						if !exist {
							unit = minion{ unitId, team, unitType, x, y, attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration,goldValue, false, false}
							playerMinions[unitId] = unit
						} else {
							unit.setHealth(health)
							unit.setPos(x, y)
						}
						playerMinions[unitId] = unit
					} else {
						unit, exist := enemyMinions[unitId]
						if !exist {
							unit = minion{ unitId, team, unitType, x, y, attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration,goldValue, false, false}
							enemyMinions[unitId] = unit
						} else {
							unit.setHealth(health)
							unit.setPos(x, y)
						}
						enemyMinions[unitId] = unit
					}
				}
			}
			//fmt.Fprintf(os.Stderr, "unitid %d team %d unitType %d x %d y %d attackRange %d health %d maxhealth %d shield %d attackdamage %d movementspeed %d stunduration %d goldvalue %d countdown1 %d countdown2 %d countdown3 %d mana %d maxmana %d manaregeneration %d herotype %d isvisible %d itemsowned %d \n", unitId, team, unitType, x, y, attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration, heroType, isVisible, itemsOwned)
		}


		var minionToLastHit int
		minionToLastHit = -1
		var enemyminionToLastHit int
		var action int
		var movebackpos int
		var numberofminionsfront int
		var numberofenemyminionsfront int
		var enemytowerdistance int

		for key, _ := range playerMinions {
			todelete := true
			for _, alive := range aliveunit {
				if alive == key { 
					todelete = false
				}
			}
			if todelete {
				delete(playerMinions, key)
			}
		}

		for key, _ := range enemyMinions {
			todelete := true
			for _, alive := range aliveunit {
				if alive == key { 
					todelete = false
				}
			}
			if todelete {
				delete(enemyMinions, key)
			}
		}

		for key, _ := range playerHeros {
			todelete := true
			for _, alive := range aliveunit {
				if alive == key { 
					todelete = false
				}
			}
			if todelete {
				delete(playerHeros , key)
			}
		}

		for key, _ := range enemyHeros {
			todelete := true
			for _, alive := range aliveunit {
				if alive == key { 
					todelete = false
				}
			}
			if todelete {
				delete(enemyHeros, key)
			}
		}

		weakesthero := findWeaskest(enemyHeros)
		numberofminionsfrontweakest := weakesthero.countMinionsFront(enemyMinions, enemyTeam)

		//for _, value := range sortedplayerHeros {
		for i:= len(sortedplayerHeros) - 1 ; i > -1 ; i-- {
			value := sortedplayerHeros[i]
			thehero := playerHeros[value]
			var nearestEnemyhero hero
			distancemax := 3000
			for _, enemyHero := range enemyHeros {
				if distance(enemyHero.x, thehero.x, enemyHero.y, thehero.y) < distancemax {
					distancemax = distance(enemyHero.x, thehero.x, enemyHero.y, thehero.y)
					nearestEnemyhero = enemyHero
				}
			}

			//if len(playerMinions) > len(enemyMinions) {
			minionToLastHit = minionsLastHit(playerMinions, thehero.attackDamage)
			//}
			enemyminionToLastHit = minionsLastHit(enemyMinions, thehero.attackDamage)
			if enemyminionToLastHit > -1 {
				if countMinionsFrontOfPos(playerMinions, myTeam, enemyMinions[enemyminionToLastHit].x) < 3 {enemyminionToLastHit = -1}
			}
			movebackpos = thehero.moveBackMinions(playerMinions, myTeam)
			numberofminionsfront = thehero.countMinionsFront(playerMinions, myTeam)
			numberofenemyminionsfront = thehero.countMinionsFront(enemyMinions, myTeam)
			enemytowerdistance = distance(thehero.x, enemyTower.x, thehero.y, enemyTower.y)

			if enemytowerdistance > enemyTower.attackRange && numberofminionsfrontweakest < 2{
				if thehero.countDown1 == 0 && thehero.mana > 49 {
					switch thehero.heroType {
						case "HULK" :
							fmt.Fprintf(os.Stdout, "CHARGE %d \n", weakesthero.unitId)
						case "DOCTOR_STRANGE":
							fmt.Fprintf(os.Stdout, "AOEHEAL %d %d \n", thehero.x, thehero.y)
					}
				} else if thehero.countDown2 == 0 && thehero.mana > 40{
					switch thehero.heroType {
						case "HULK" :
							fmt.Fprintf(os.Stdout, "EXPLOSIVESHIELD\n")
						case "DOCTOR_STRANGE":
							fmt.Fprintf(os.Stdout, "SHIELD %d \n", thehero.unitId)
					}
				} else if thehero.countDown3 == 0 && thehero.mana > 40 {
					switch thehero.heroType {
						case "HULK" :
							fmt.Fprintf(os.Stdout, "BASH %d \n", weakesthero.unitId)
						case "DOCTOR_STRANGE":
							fmt.Fprintf(os.Stdout, "PULL %d \n", weakesthero.unitId)
					}
				} else {
					fmt.Fprintf(os.Stderr, "ATTACK BANZAI %d \n", weakesthero.unitId)
					fmt.Fprintf(os.Stdout, "ATTACK %d \n", weakesthero.unitId)
				}
			} else if enemytowerdistance < (enemyTower.attackRange + 100) || numberofminionsfront < 1 {
				// On se replie si le hero est trop proche de la tour enemie
				fmt.Fprintf(os.Stderr, "MINIONS FRONT %d \n", numberofminionsfront)
				thehero.moveBack(playerTower, "NO MINIONS FRONT")
			} else if !minionsAreHit(playerMinions) && movebackpos > -1 {
				//Si les minions alliés ne sont pas agros et que le hero doit se replier
				fmt.Fprintf(os.Stderr, "MOVE BACK MINIONS %d \n", movebackpos)
				fmt.Fprintf(os.Stdout, "MOVE %d %d \n", movebackpos, thehero.y)
			//} else if thehero.iscritical > 5 && thehero.health < 400 {
				// On se replie si notre hero est touché
				//fmt.Fprintf(os.Stderr, "HERO IS HIT \n")
			//	thehero.moveBack(playerTower, "HERO IS HIT")
			} else if minionToLastHit > -1 {
				// on last hit le minion allié
				fmt.Fprintf(os.Stderr, "MINION TO LAST HIT %d \n", minionToLastHit)
				fmt.Fprintf(os.Stdout, "ATTACK %d \n", minionToLastHit)
				delete(enemyMinions, minionToLastHit)
			 } else if enemyminionToLastHit > -1 {
			 	// on last hit le minion enemie
		 		fmt.Fprintf(os.Stderr, "ENEMY MINION TO LAST HIT %d \n", enemyminionToLastHit)
		 		fmt.Fprintf(os.Stdout, "ATTACK %d \n", enemyminionToLastHit)
		 		delete(enemyMinions, minionToLastHit)
			} else if numberofenemyminionsfront < 2 && enemytowerdistance < (enemyTower.attackRange + 200) {
				// on attack la tour enemie si plus de minions
				fmt.Fprintf(os.Stdout, "ATTACK %d \n", enemyTower.unitId)
			} else if len(enemyMinions) < 3 && distance(thehero.x, thehero.y, nearestEnemyhero.x, nearestEnemyhero.y) < 200 && enemytowerdistance > (enemyTower.attackRange + 100) {
			 	// on kill le hero enemie
			 	fmt.Fprintf(os.Stderr, "ATTACK NEAREST HERO %d \n", nearestEnemyhero.unitId)
				fmt.Fprintf(os.Stdout, "ATTACK %d \n", nearestEnemyhero.unitId)
			} else {
				action = thehero.attackLeastHealthMinions(enemyMinions)
				if action > -1 {
					fmt.Fprintf(os.Stderr, "ATTACK LESS HEALTH MINIONS %d \n", action)
					fmt.Fprintf(os.Stdout, "ATTACK %d \n", action)
					unit := enemyMinions[action]
					unit.setHealth(unit.health - thehero.attackDamage)
				} else {
					// On tue les minions les plus proches
					if gold > 400 && len(heroItem) < 3 {
						fmt.Fprintf(os.Stderr, "HERO COUNT ITEMS %d %d\n", len(heroItem), heroItem)
						if thehero.health < 400 { 
							buy(items, "healthpotion", gold, &heroItem)
						} else {
							buy(items, "weapon", gold, &heroItem)
						}
					} else {
						fmt.Fprintf(os.Stderr, "ATTACK NEAREST GARBAGE MINIONS  \n")
						fmt.Fprintf(os.Stdout, "ATTACK_NEAREST UNIT \n")
					}
				}
			}
		}

		//fmt.Fprintf(os.Stderr, "%d %d \n", playerTower, enemyTower)
		// If roundType has a negative value then you need to output a Hero name, such as "DEADPOOL" or "VALKYRIE".
		// Else you need to output roundType number of any valid action, such as "WAIT" or "ATTACK unitId"
		// fmt.Println("ATTACK_NEAREST HERO")
	}
}