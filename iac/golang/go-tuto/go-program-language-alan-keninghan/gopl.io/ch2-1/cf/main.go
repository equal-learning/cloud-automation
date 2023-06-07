package main

import (
	"fmt"
	"os"
	"strconv"

	"tempconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("No value supplied")
		os.Exit(1)
	} else {
		v, err := strconv.ParseFloat(os.Args[1], 64)

		if err != nil {
			fmt.Printf("Could not parse String to Fload")
			os.Exit(1)
		} else {

			f := tempconv.Fahrenheit(v)
			fmt.Printf("%s", f)

		}

		fmt.Sprintf("%s", f)
		fmt.Sprintf("%s", c)

	}

}
