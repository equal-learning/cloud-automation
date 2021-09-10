/*

STRUCTS ::

A struct is a collection of fields.

 */

package main

import "fmt"

type Vertex struct {  // Two steps process: first declare then use
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}