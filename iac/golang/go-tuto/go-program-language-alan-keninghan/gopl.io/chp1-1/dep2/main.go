// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args

	for _, file := range files[1:] {

		countduplicate(file)

	}

}

func countduplicate(filename string) {
	// f, err := os.Open("data.txt")
	f, err := os.Open(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	} else {

		freader := bufio.NewScanner(f)
		repeats := make(map[string]int)
		for freader.Scan() {
			repeats[freader.Text()]++

		}

		for k, v := range repeats {

			fmt.Printf("%s,%d \n", k, v)
		}

	}
	f.Close()

}
