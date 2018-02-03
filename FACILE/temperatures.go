package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// n: the number of temperatures to analyse
	var n int
	var min float64
	var minreel float64
	var result float64
	min = 5527
	
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&n)
	
	scanner.Scan()
	inputs := strings.Split(scanner.Text()," ")
	
	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t,_ := strconv.ParseFloat(inputs[i],32)
		_ = t
		
		if (math.Abs(t) < min) {
			min = math.Abs(t)
			minreel = t
			result = t
		} else {
			if (math.Abs(t) == min) {
				if!(t == minreel){
					min = math.Abs(t)
					result = math.Abs(t)
				}
			}
		}
	}
	if n == 0 { result = 0 }
	
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(result)// Write answer to stdout
}