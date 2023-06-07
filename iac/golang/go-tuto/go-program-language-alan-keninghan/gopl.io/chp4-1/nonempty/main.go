package main

import "fmt"

func main() {

	strs := []string{"one", "", "three", "", "five"}
	nonempty := make([]string, 0)

	for _, val := range strs {
		if val != "" {
			nonempty = append(nonempty, val)
		}

	}

	fmt.Println(nonempty)

}
