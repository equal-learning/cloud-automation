// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("No arguments !")

	} else {

		for _, arg := range args {

			fmt.Println(arg)

		}
	}

}
