// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//buff := make([]byte, 5)
	reader := bufio.NewScanner(os.Stdin)
	repeats := make(map[string]int)

	for reader.Scan() {

		/*
			text := reader.Text()

			fmt.Println(text)
			wds := strings.Split(text, " ")

			for i := 0; i <= len(wds)-1; i++ {
				i, ok := mp[wds[i]]
				if !ok {
					mp[wds[i]] = 1
				} else {

					mp[wds[i]] = mp[wds[i]] + 1
				}

			}
		*/

		repeats[reader.Text()]++

	}

	for k, val := range repeats {
		fmt.Printf("%s,%d", k, val)
	}

}
