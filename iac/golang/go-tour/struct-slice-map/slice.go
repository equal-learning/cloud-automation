/*

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, FLEXIBLE VIEW into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]
This selects a HALF-OPEN range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4]

More on slices : https://go.dev/blog/slices-intro

*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // Remark :: Slice has no fixed size
	fmt.Println(s)
}
