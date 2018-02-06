package main

import "fmt"
import "os"
import "bufio"
//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    scanner.Scan()
    message := scanner.Text()

	var newmessage string
	var bits string
	var change bool;
	var unary int32;

	for _,letter := range message {
		bits = fmt.Sprintf("%7b", letter);
		for _,bit := range bits {
			if!(bit == unary) { 
				unary = bit
				change = true
			}
			if(change) {
				if(bit == 49) {
					newmessage = newmessage + " 0 "
				} else {
					newmessage = newmessage + " 00 "
				}
				change = false;
			}
			newmessage = newmessage + "0"
		}
	}
	newmessage = newmessage[1:len(newmessage)]

	//fmt.Fprintln(os.Stderr, MESSAGE)
	fmt.Println(newmessage)// Write answer to stdout
}