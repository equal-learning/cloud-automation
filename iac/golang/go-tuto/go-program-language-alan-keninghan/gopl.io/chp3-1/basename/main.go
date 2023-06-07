package main

import (
	"fmt"
	"strings"
)

// Basename1 reads file names from stdin and prints the base name of each one.

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c

func main() {

	bsnm1 := "/usr/local/hello.txt"

	if strings.Contains(bsnm1, "/") {
		fmt.Println("complex one")
		parts := strings.Split(bsnm1, "/")

		fmt.Printf(parts[len(parts)-1])

	} else if strings.Contains(bsnm1, ".") {
		fmt.Println("Simple one")

	}

}
