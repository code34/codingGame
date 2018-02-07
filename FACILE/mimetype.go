package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// N: Number of elements which make up the association table.
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&N)
	
	// Q: Number Q of file names to be analyzed.
	var Q int
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&Q)

	hashmap := make(map[string]string)
	
	// EXT: file extension
	// MT: MIME type.
	for i := 0; i < N; i++ {
		var EXT, MT string
		scanner.Scan()
		fmt.Sscan(scanner.Text(),&EXT, &MT)
		hashmap[strings.ToLower(EXT)] = MT
	}
	
	var fileext string
	
	//fmt.Printf("Mime: %d Fichiers: %d Extension: %d %s\n", N, Q, len(extension))
	
	// FNAME: file name
	for i := 0; i < Q; i++ {
		scanner.Scan()
		FNAME := scanner.Text()
		//fmt.Printf("fichier: %s \n",FNAME)
		index := strings.LastIndex(FNAME, ".")
		if index >= 0 { fileext = (FNAME)[index+1:len(FNAME)] } else { fileext = "-999" }
		//fmt.Printf("extension: %s \n",fileext)
		
		result,ok := hashmap[strings.ToLower(fileext)]
		
		if!(ok) {
		   fmt.Println("UNKNOWN")
			//fmt.Fprintln(os.Stderr, "UNKNOWN")
		} else {
			fmt.Println(result)
			//fmt.Fprintln(os.Stderr,fileext)
		}
	}

	//fmt.Printf("%d %s\n", N, extension)
	//fmt.Printf("%s \n", FNAME)    
	//fmt.Fprintln(os.Stderr, N)
	//fmt.Fprintln(os.Stderr, Q)
	// For each of the Q filenames, display on a line the corresponding MIME type. If there is no corresponding type, then display UNKNOWN.
}