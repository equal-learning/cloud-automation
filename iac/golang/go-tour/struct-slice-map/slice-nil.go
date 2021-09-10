/*

"Nil slices" ::

The zero like value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

 */

package main

import "fmt"

func main() {
	var s []int  // A nil slice

	num := [2]int{1,2}

	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}

	s=num[0:1]

	if s == nil {
		fmt.Println("nil!")
	}

}