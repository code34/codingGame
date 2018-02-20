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
	
	for {
		// hSpeed: the horizontal speed (in m/s), can be negative.
		// vSpeed: the vertical speed (in m/s), can be negative.
		// fuel: the quantity of remaining fuel in liters.
		// rotate: the rotation angle in degrees (-90 to 90).
		// power: the thrust power (0 to 4).
		var engineL, engine int
		var X, Y, hSpeed, vSpeed, fuel, rotate, power int
		fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)
		fmt.Fprintf(os.Stderr, "VARIABLES X: %d Y: %d hSpeed: %d vSpeed: %d fuel: %d rotate: %d power: %d\n", X, Y, hSpeed, vSpeed, fuel, rotate, power)
		
		if X < leftX {
			engineL = -20
		} else if X > rightX {
			engineL = 20
		} else {
			if hSpeed < -11 {
				engineL = -20
			} else if hSpeed > 11 {
				engineL = 20
			} else {
				engineL = 0
			}
		}
		
		if hSpeed < -30 {
			engineL = -20
		} else if hSpeed > 30 {
			engineL = 20
		} else {
			engineL = engineL   
		}
			
		if hSpeed > 50 {
			engineL = 50
			engine = 4
		} else if hSpeed < -50 {
			engineL = -50
			engine = 4
		}

		if vSpeed < -30 {
			engine = 4
		} else {
			engine = 3
		}

		fmt.Fprintf(os.Stderr, "VARIABLES vertical speed: %d %d %d\n", engineL, hSpeed, vSpeed)
		fmt.Println(engineL,engine)
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		
		// rotate power. rotate is the desired rotation angle. power is the desired thrust power.
		//fmt.Println("-20 3")
	}
}