// Fetch saves the contents of a URL into a local file.

package main

import (
	"fmt"
	"os"
)

// Fetch downloads the URL and returns the
// name and length of the local file.

func fetch(fname string) (name string, len int64) {

}

func main() {

	fmt.Println("Hello ..")

	if len(os.Args) < 1 {
		os.Exit(1)
	}

}
