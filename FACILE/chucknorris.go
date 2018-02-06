package main

import "fmt"
import "os"
import "bufio"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	message := scanner.Text()

	var newmessage string
	var bits string
	var prev int32;

	for _,letter := range message {
		bits = fmt.Sprintf("%.7b", letter);
		for _,bit := range bits {
			if(bit == prev) { 
				newmessage = newmessage + "0"
			} else {
				if(bit == 49) {
					newmessage = newmessage + " 0 0"
				} else {
					newmessage = newmessage + " 00 0"
				}
				prev = bit
			}
		}
	}
	newmessage = newmessage[1:len(newmessage)]

	//fmt.Fprintln(os.Stderr, MESSAGE)
	fmt.Println(newmessage)// Write answer to stdout
}