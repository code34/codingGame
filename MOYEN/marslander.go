package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func staticspeed(vSpeed int, landY int, Y int) {
	var speedlimit int = -30
	if Y - 100 < landY { speedlimit = 10 }
	if vSpeed < speedlimit  {
			fmt.Println("0 4")
	} else {
			fmt.Println("0 3")
	}
}

func main() {
	var surfaceN int
	var groundY, leftX, rightX int
	var landsitefound bool = false
	fmt.Scan(&surfaceN)
	
	for i := 0; i < surfaceN; i++ {
		var landX, landY int
		fmt.Scan(&landX, &landY)
		
		if !landsitefound {
			if groundY == landY {
				rightX = landX
				landsitefound = true
			} else {
				leftX = landX
				groundY = landY
			}
		}
	}

	var step = 1    
	for {
		var X, Y, hSpeed, vSpeed, fuel, rotate, power int
		fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)

		switch step {        
			case 1:{
				fmt.Fprintf(os.Stderr, "PHASE1 %d\n", power)
				if X < leftX {
					if hSpeed < 20 {
						fmt.Println("-20 4")
						fmt.Fprintf(os.Stderr, "PATHA %d\n", power)
					} else {
						if hSpeed > 20 {
							fmt.Println("30 4")
						} else {
							staticspeed(vSpeed, groundY, Y)
						}
					}   
				} else if X > rightX {
					if hSpeed > -20 {
						fmt.Println("20 4")
						fmt.Fprintf(os.Stderr, "PATHB %d\n", power)
					} else {
						if hSpeed < -20 {
							fmt.Println("-30 4")
						} else {
							staticspeed(vSpeed, groundY, Y)
						}
					}    
				} else if X > leftX && X < rightX {
					if hSpeed > -2 && hSpeed < 2 {
						fmt.Println("0 4")
						step = 2
						fmt.Fprintf(os.Stderr, "PHASE2B %d\n", power)                    
					} else {
						if hSpeed < -1 {  fmt.Println("-20 4") }
						if hSpeed > 1 { fmt.Println("20 4") }
					}
				}
			}
			
			case 2:{
				fmt.Fprintf(os.Stderr, "PHASE2 %d\n", power)
				if vSpeed < -30  {
					fmt.Println("0 4")
				} else {
					fmt.Println("0 3")
				}
			}
		}
	}
}