package main

import "fmt"
import "os"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func checkcard(card string) int {
	if len(card) == 3 { card = card[0:2] } else { card = card[0:1] }
	if card == "J" { card = "11"}
	if card == "Q" { card = "12"}
	if card == "K" { card = "13"}
	if card == "A" { card = "14"}
	intcard,_ := strconv.Atoi(card)
	return intcard
}

func battle(card1 int, card2 int) int {
	return card1
}

func main() {
	// n: the number of cards for player 1
	var n int
	fmt.Scan(&n)

	var handplayer1 []int
	var handplayer2 []int
	var defausse1 []int
	var defausse2 []int

	for i := 0; i < n; i++ {
		// cardp1: the n cards of player 1
		var cardp1 string
		fmt.Scan(&cardp1)
		handplayer1 = append(handplayer1, checkcard(cardp1))
	}
	// m: the number of cards for player 2
	var m int
	fmt.Scan(&m)
	
	for i := 0; i < m; i++ {
		// cardp2: the m cards of player 2
		var cardp2 string
		fmt.Scan(&cardp2)
		handplayer2 = append(handplayer2, checkcard(cardp2))
	}
	
	fmt.Fprintf(os.Stderr, "CARDS P1: %d \n", handplayer1)
	fmt.Fprintf(os.Stderr, "CARDS P2: %d \n", handplayer2)
	
	var count int
	var winner int
	var end bool = false
	
	for {
		if len(handplayer1) == 0 { winner = 2; break }
		if len(handplayer2) == 0 { winner = 1; break }
		if end { winner = 0; break}
		
		result:= handplayer1[0] - handplayer2[0]
		switch true {
			case 0 == result:
				fmt.Fprintf(os.Stderr, "BATAILLE: %d %d \n", handplayer1, handplayer2)
				if len(handplayer1) < 4 { end = true; break }
				if len(handplayer2) < 4 { end = true; break }
				defausse1 = append(defausse1, handplayer1[0:4]...)
				defausse2 = append(defausse2, handplayer2[0:4]...)
				handplayer1 = handplayer1[4:]
				handplayer2 = handplayer2[4:]
				//fmt.Fprintf(os.Stderr, "DEFAUSSE: %d %d %d %d \n", defausse1, defausse2, handplayer1, handplayer2)
			case result > 0:
				defausse1 = append(defausse1, handplayer1[0])
				defausse2 = append(defausse2, handplayer2[0])
				handplayer1 = append(handplayer1, defausse1...)
				handplayer1 = append(handplayer1, defausse2...)
				fmt.Fprintf(os.Stderr, "SET PLAYER1: %d %d %d %d \n", defausse1, defausse2, handplayer1, handplayer2)
				defausse1 = nil
				defausse2 = nil
				handplayer1 = handplayer1[1:]
				handplayer2 = handplayer2[1:]
				fmt.Fprintf(os.Stderr, "PLAYER1 WIN %d %d\n", handplayer1, handplayer2)
				count++
			case result < 0:
				defausse1 = append(defausse1, handplayer1[0])
				defausse2 = append(defausse2, handplayer2[0])
				handplayer2 = append(handplayer2, defausse1...)
				handplayer2 = append(handplayer2, defausse2...)
				fmt.Fprintf(os.Stderr, "SET PLAYER2: %d %d %d %d \n", defausse1, defausse2, handplayer1, handplayer2)
				defausse1 = nil
				defausse2 = nil
				handplayer1 = handplayer1[1:]
				handplayer2 = handplayer2[1:]
				fmt.Fprintf(os.Stderr, "PLAYER2 WIN %d \n", handplayer2)
				count++
		}
	}
	
	//fmt.Fprintf(os.Stderr, "CARD: %s \n", cardp2)
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	if winner > 0 { 
		fmt.Printf("%d %d\n", winner, count)
	} else {
		fmt.Println("PAT")// Write answer to stdout
	}
}