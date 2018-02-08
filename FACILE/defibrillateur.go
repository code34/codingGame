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

	var LON string
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&LON)
	LON = strings.Replace(LON,",",".",-1)
	RLON,_ := strconv.ParseFloat(LON,64)
	RLON = RLON * (math.Pi / 180)
	
	var LAT string
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&LAT)
	LAT = strings.Replace(LAT,",",".",-1)
	RLAT,_ := strconv.ParseFloat(LAT,64)
	RLAT = RLAT * (math.Pi / 180)
	
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&N)

	var distancemin float64 = 10000000
	var name string
	
	for i := 0; i < N; i++ {
		scanner.Scan()
		defib := scanner.Text()
		temp := strings.Split(defib, ";")
		xdefib:= strings.Replace(temp[4],",",".",-1)
		ydefib:= strings.Replace(temp[5],",",".",-1)
		rxdefib,_ := strconv.ParseFloat(xdefib,64)
		rydefib,_ := strconv.ParseFloat(ydefib,64)           
		rxdefib = rxdefib * math.Pi / 180
		rydefib = rydefib * math.Pi / 180
		xfinal := (rxdefib - RLON) * math.Cos((rydefib + RLAT)/2)
		yfinal := rydefib - RLAT
		d := math.Sqrt(math.Pow(xfinal,2) + math.Pow(yfinal,2)) * 6371 
		if(d < distancemin) {
			distancemin = d
			name = temp[1]
		}
		//fmt.Printf("%f %f %f %f %f\n", RLON, RLAT, xfinal, yfinal, d)
	}
	
	//fmt.Printf("%s %s %d \n", LON, LAT, N)
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(name)// Write answer to stdout
}