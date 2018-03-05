package main

import "fmt"
import "os"

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
}

func (hero hero) moveBack (side int) {
	var moveX int
	moveX = hero.x

	if side == 0 {
		moveX = hero.x - (hero.attackRange -10)
	} else {
		moveX = hero.x + (hero.attackRange -10)
	}
	fmt.Fprintf(os.Stderr, "MOVE %d %d \n", hero.x, side)
	fmt.Fprintf(os.Stdout, "MOVE %d %d \n", moveX, 590)
}

func (hero hero) moveBackMinions (minions map[int]minion, side int) {
	var moveX int
	moveX = hero.x

	for _, minion := range minions {
		// le hero se trouve à gauche (0) sinon à droite (1)
		if side == 0 {
			moveX = minion.x - (hero.attackRange -10)
		} else {
			moveX = minion.x + (hero.attackRange -10)
		}
	}
	fmt.Fprintf(os.Stdout, "MOVE %d %d \n", moveX, 590)
}

func (hero *hero) attackMinions (minions map[int]minion) {
	fmt.Fprintf(os.Stdout, "ATTACK_NEAREST UNIT \n")
}

func (hero *hero) setHealth (health int) {
	if health == hero.health { hero.ishit = false } else { hero.ishit = true; hero.health = health }
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


func main() {
	var myTeam int
	fmt.Scan(&myTeam)
	
	// bushAndSpawnPointCount: usefrul from wood1, represents the number of bushes and the number of places where neutral units can spawn
	var bushAndSpawnPointCount int
	fmt.Scan(&bushAndSpawnPointCount)
	
	for i := 0; i < bushAndSpawnPointCount; i++ {
		// entityType: BUSH, from wood1 it can also be SPAWN
		var entityType string
		var x, y, radius int
		fmt.Scan(&entityType, &x, &y, &radius)
	}
	// itemCount: useful from wood2
	var itemCount int
	fmt.Scan(&itemCount)
	
	for i := 0; i < itemCount; i++ {
		// itemName: contains keywords such as BRONZE, SILVER and BLADE, BOOTS connected by "_" to help you sort easier
		// itemCost: BRONZE items have lowest cost, the most expensive items are LEGENDARY
		// damage: keyword BLADE is present if the most important item stat is damage
		// moveSpeed: keyword BOOTS is present if the most important item stat is moveSpeed
		// isPotion: 0 if it's not instantly consumed
		var itemName string
		var itemCost, damage, health, maxHealth, mana, maxMana, moveSpeed, manaRegeneration, isPotion int
		fmt.Scan(&itemName, &itemCost, &damage, &health, &maxHealth, &mana, &maxMana, &moveSpeed, &manaRegeneration, &isPotion)
	}

	playerMinions := make(map[int]minion)
	enemyMinions := make(map[int]minion)
	playerHeros := make(map[int]hero)
	enemyHeros := make(map[int]hero)

	for {
		var playerTower tower
		var enemyTower tower
		var enemyHero hero
		
		var gold int
		fmt.Scan(&gold)
		
		var enemyGold int
		fmt.Scan(&enemyGold)
		
		// roundType: a positive value will show the number of heroes that await a command
		var roundType int
		fmt.Scan(&roundType)
		if roundType < 0 {
			fmt.Println("DOCTOR_STRANGE")
		}
		
		var entityCount int
		fmt.Scan(&entityCount)
		
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
			
			switch (unitType) {
				case "TOWER" : {
					unit:= tower{ unitId, team, unitType, x, y, attackRange, health, maxHealth,attackDamage }
					if team == myTeam { playerTower = unit } else { enemyTower = unit }
				}
				
				case "HERO" : {
					if team == myTeam { 
						unit, exist := playerHeros[unitId] 
						if !exist {
							unit = hero{ unitId, team,unitType, x, y,attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration, heroType, isVisible, itemsOwned, true, false}
						} else {
							unit.setHealth(health)
							unit.setPos(x, y)
						}
						playerHeros[unitId] = unit
						//playerHero = unit
					} else {
						unit, exist := enemyHeros[unitId] 
						if !exist {
							unit = hero{ unitId, team,unitType, x, y,attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration, heroType, isVisible, itemsOwned, true, false}
						} else {
							unit.setHealth(health)
							unit.setPos(x, y)
						}
						enemyHeros[unitId] = unit
						enemyHero = unit
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
							//unit.setHealth(health)
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
							//unit.setHealth(health)
							unit.setPos(x, y)
						}
						enemyMinions[unitId] = unit
					}
				}
			}
			//fmt.Fprintf(os.Stderr, "unitid %d team %d unitType %d x %d y %d attackRange %d health %d maxhealth %d shield %d attackdamage %d movementspeed %d stunduration %d goldvalue %d countdown1 %d countdown2 %d countdown3 %d mana %d maxmana %d manaregeneration %d herotype %d isvisible %d itemsowned %d \n", unitId, team, unitType, x, y, attackRange, health, maxHealth, shield, attackDamage, movementSpeed, stunDuration, goldValue, countDown1, countDown2, countDown3, mana, maxMana, manaRegeneration, heroType, isVisible, itemsOwned)
		}

		for key, minion := range playerMinions {
			if minion.health < 200 {
				delete(playerMinions, key)
			}
		}

		for key, minion := range enemyMinions {
			if minion.health < 200 {
				delete(enemyMinions, key)
			}
		}
		fmt.Fprintf(os.Stderr, "LA VIE DES MINIONS !!!!! %d %d\n", len(playerMinions), len(enemyMinions));

		if minionsAreHit(playerMinions) {
			//fmt.Fprintf(os.Stderr, "ALERT BACK !!!!! %d \n", len(playerMinions), playerMinions);
			if len(playerMinions) < 2 {
				for _, hero := range playerHeros {
					hero.moveBack(myTeam)
				}
			} else if playersAreHit(playerHeros) {
				for _, hero := range playerHeros {
					hero.moveBack(myTeam)
				}
			} else {
				for _, hero := range playerHeros {
					hero.attackMinions(enemyMinions)
				}
			}
		} else {
			for _, hero := range playerHeros {
				hero.moveBackMinions(playerMinions, myTeam)
			}
		}

		fmt.Fprintf(os.Stderr, "%d %d %d", playerTower, enemyTower, enemyHero)
		// If roundType has a negative value then you need to output a Hero name, such as "DEADPOOL" or "VALKYRIE".
		// Else you need to output roundType number of any valid action, such as "WAIT" or "ATTACK unitId"
		// fmt.Println("ATTACK_NEAREST HERO")
	}
}