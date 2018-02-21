package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// surfaceN: the number of points used to draw the surface of Mars.
	var surfaceN int
	var groundY, leftX, rightX int
	var landsitefound bool = false
	fmt.Scan(&surfaceN)
	
	for i := 0; i < surfaceN; i++ {
		// landX: X coordinate of a surface point. (0 to 6999)
		// landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
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
	fmt.Fprintf(os.Stderr, "LANDSITE %d %d %d\n", leftX, rightX, groundY)

	var center int
	center = (leftX + rightX) / 2
	var step = 1    
	for {
		// hSpeed: the horizontal speed (in m/s), can be negative.
		// vSpeed: the vertical speed (in m/s), can be negative.
		// fuel: the quantity of remaining fuel in liters.
		// rotate: the rotation angle in degrees (-90 to 90).
		// power: the thrust power (0 to 4).
		var engineL, engine int
		var X, Y, hSpeed, vSpeed, fuel, rotate, power int
		fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)

		switch step {
			case 1:{
				if X < center {
					if hSpeed < 20 {
						fmt.Println("-20 4")
						fmt.Fprintf(os.Stderr, "PATHA %d\n", power)
					} else {
						if hSpeed > 20 {
							fmt.Println("20 4")
						} else {
							fmt.Println("0 4")
							step = 2
							fmt.Fprintf(os.Stderr, "PHASE2A %d\n", power)
						}
					}   
				} else if X > center {
					// si on se décale à gauche +20ms
					if hSpeed > -20 {
						fmt.Println("20 4")
						fmt.Fprintf(os.Stderr, "PATHB %d\n", power)
					} else {
						// si on se décale à gauche à plus de 20ms
						if hSpeed < -20 {
							fmt.Println("-20 4")
						} else {
							// on se maintient à la meme altitude
							fmt.Println("0 4")
							step = 2
							fmt.Fprintf(os.Stderr, "PHASE2B %d\n", power)
						}
					}    
				} else if X > leftX && X < rightX {
					if hSpeed > -20 && hSpeed < 20 {
						fmt.Println("0 4")
						step = 2
						fmt.Fprintf(os.Stderr, "PHASE2B %d\n", power)                    
					}
				}
			}
			
			case 2:{
				if X > leftX && X < rightX {
					if hSpeed < 20 {
						   fmt.Println("-20 3")
					} else if hSpeed > 20 {
						fmt.Println("20 3")
					} else {
						fmt.Println("0 4")
						step = 3   
					}
				} else {
					if vSpeed < -30 && Y > (groundY + 100) {
						fmt.Println("0 4")
					} else {
						fmt.Println("0 3")
					}
				}
			}
			
			case 3:{
				if vSpeed > -39 {
					fmt.Println("0 4")
				} else {
					fmt.Println("0 3")
				}
			}
		}


		
		fmt.Fprintf(os.Stderr, "VARIABLES X: %d Y: %d hSpeed: %d vSpeed: %d fuel: %d rotate: %d power: %d\n", X, Y, hSpeed, vSpeed, fuel, rotate, power)
		fmt.Fprintf(os.Stderr, "VARIABLES vertical speed: %d %d %d %d\n", engine, engineL, hSpeed, vSpeed)
		//fmt.Println(engineL,engine)
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		
		// rotate power. rotate is the desired rotation angle. power is the desired thrust power.
	}
}