/*

Pointers to structs ::

Struct fields can be accessed through a struct pointer.

 */

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9  // Remark : Alternatively we could write (*p).X, but cumbersome, so instead the short hand is p.X
	fmt.Println(v)
}
